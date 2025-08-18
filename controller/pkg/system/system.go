package system

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32              = windows.NewLazySystemDLL("user32.dll")
	monitorFromWindow   = user32.NewProc("MonitorFromWindow")
	getForegroundWindow = user32.NewProc("GetForegroundWindow")
	getWindowRect       = user32.NewProc("GetWindowRect")
	setWindowPos        = user32.NewProc("SetWindowPos")

	enumDisplayMonitors = user32.NewProc("EnumDisplayMonitors")
	getMonitorInfo      = user32.NewProc("GetMonitorInfoW")
)

type System struct {
}

func New() (*System, error) {
	var s System
	return &s, nil
}

func (z *System) GetCurrentWindowHandle() (uintptr, error) {
	hwnd, _, err := getForegroundWindow.Call()
	if hwnd == 0 {
		return 0, err
	}
	return hwnd, nil
}

func (z *System) getMonitorHandles() ([]syscall.Handle, error) {
	var hMonitors []syscall.Handle
	var monitorEnumCallback = syscall.NewCallback(
		func(hMonitor syscall.Handle, hdcMonitor syscall.Handle, lprcMonitor *RectRaw, dwData uintptr) uintptr {
			monitorsPtr := (*[]syscall.Handle)(unsafe.Pointer(dwData)) // nolint:unsafeptr
			*monitorsPtr = append(*monitorsPtr, hMonitor)
			return 1 // 继续枚举
		},
	)
	ret, _, err := enumDisplayMonitors.Call(
		0,
		0,
		monitorEnumCallback,
		uintptr(unsafe.Pointer(&hMonitors)),
	)
	if ret == 0 { // 成功时 err 不是 nil 是成功信息
		return nil, err
	}
	return hMonitors, nil
}

func (z *System) getMonitorHandle(hwnd uintptr) (syscall.Handle, error) {
	hMonitor, _, err := monitorFromWindow.Call(hwnd, 0x00000002)
	if hMonitor == 0 {
		return 0, err
	}
	return syscall.Handle(hMonitor), nil
}

func (z *System) getMonitorByHandle(hMonitor syscall.Handle) (Monitor, error) {
	var monitor MonitorRaw
	monitor.CbSize = uint32(unsafe.Sizeof(monitor))
	ret, _, err := getMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(&monitor)),
	)
	if ret == 0 {
		return Monitor{}, err
	}
	var m Monitor
	m.DeviceName = syscall.UTF16ToString(monitor.SzDevice[:])
	m.IsPrimary = monitor.DwFlags&0x00000001 != 0

	m.Screen.X = monitor.RcMonitor.Left
	m.Screen.Y = monitor.RcMonitor.Top
	m.Screen.Width = monitor.RcMonitor.Right - monitor.RcMonitor.Left
	m.Screen.Height = monitor.RcMonitor.Bottom - monitor.RcMonitor.Top

	m.Screen.Bottom = monitor.RcMonitor.Bottom
	m.Screen.Top = monitor.RcMonitor.Top
	m.Screen.Left = monitor.RcMonitor.Left
	m.Screen.Right = monitor.RcMonitor.Right

	m.Desktop.X = monitor.RcWork.Left
	m.Desktop.Y = monitor.RcWork.Top
	m.Desktop.Width = monitor.RcWork.Right - monitor.RcWork.Left
	m.Desktop.Height = monitor.RcWork.Bottom - monitor.RcWork.Top

	m.Desktop.Top = monitor.RcWork.Top
	m.Desktop.Bottom = monitor.RcWork.Bottom
	m.Desktop.Left = monitor.RcWork.Left
	m.Desktop.Right = monitor.RcWork.Right

	m.Handle = hMonitor

	monitorNumber, err := z.getMonitorNumberByHandle(hMonitor)
	if err != nil {
		return Monitor{}, err
	}
	m.Number = monitorNumber

	return m, nil
}

func (z *System) getMonitorNumberByHandle(hMonitor syscall.Handle) (int, error) {
	hMonitors, err := z.getMonitorHandles()
	if err != nil {
		return 0, err
	}
	for i, h := range hMonitors {
		if h == hMonitor {
			return i + 1, nil
		}
	}
	return 0, errors.New("monitor not found")
}

func (z *System) GetMonitors() ([]Monitor, error) {
	var monitors []Monitor
	hMonitors, err := z.getMonitorHandles()
	if err != nil {
		return nil, err
	}
	for i, hMonitor := range hMonitors {
		m, err := z.getMonitorByHandle(hMonitor)
		if err != nil {
			continue // 忽略获取信息失败的显示器
		}
		m.Number = i + 1
		monitors = append(monitors, m)
	}
	return monitors, nil
}

func (z *System) GetPrimaryMonitorNumber() (int, error) {
	monitors, err := z.GetMonitors()
	if err != nil {
		return 0, err
	}
	for _, m := range monitors {
		if m.IsPrimary {
			return m.Number, nil
		}
	}
	return 0, errors.New("no primary")
}

func (z *System) GetWindow(hwnd uintptr) (Window, error) {
	var w RectRaw
	ret, _, err := getWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&w)))
	if ret == 0 {
		return Window{}, err
	}
	hMonitor, err := z.getMonitorHandle(hwnd)
	if err != nil {
		return Window{}, err
	}
	monitor, err := z.getMonitorByHandle(hMonitor)
	if err != nil {
		return Window{}, err
	}
	var window Window
	window.Handle = hwnd
	window.X = w.Left - monitor.Desktop.Left
	window.Y = w.Top - monitor.Desktop.Top
	window.Width = w.Right - w.Left
	window.Height = w.Bottom - w.Top
	window.Left = w.Left - monitor.Desktop.Left
	window.Top = w.Top - monitor.Desktop.Top
	window.Right = w.Right - monitor.Desktop.Left
	window.Bottom = w.Bottom - monitor.Desktop.Top
	// window.MonitorHandle = hMonitor
	window.MonitorNumber = monitor.Number
	return window, nil
}

func (z *System) SetWindow(hwnd uintptr, x, y, width, height int32, monitorNumber int) error {
	monitors, err := z.GetMonitors()
	if err != nil {
		return err
	}
	monitor := monitors[monitorNumber-1]
	vX := monitor.Desktop.X + x
	vY := monitor.Desktop.Y + y
	// 设置窗口位置和大小
	// 跨不同缩放比的显示器切换时，第一次设置，窗口的大小总是不合预期；但发现第二次设置，因为不跨显示器，所以总是符合预期。
	// 尝试了各种高级方法，未果。
	// 最终选择设置两次。
	for i := 0; i < 2; i++ {
		ret, _, _ := setWindowPos.Call(
			hwnd,
			0, // HWND_TOP (将窗口置于Z序的顶部)
			uintptr(vX),
			uintptr(vY),
			uintptr(width),
			uintptr(height),
			uintptr(0x0004|0x0010), // SWP_NOZORDER (保持当前Z序，忽略第一个参数)
		)
		if ret == 0 {
			return errors.New("set window failed")
		}
	}
	// 校验
	var w RectRaw
	ret, _, err := getWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&w)))
	if ret == 0 {
		return err
	}
	nX := w.Left
	nY := w.Top
	nWidth := w.Right - w.Left
	nHeight := w.Bottom - w.Top
	if vX != nX || vY != nY || width != nWidth || height != nHeight {
		return errors.New("position result unexpected")
	}
	return nil
}
func (z *System) GetMonitor(hwnd uintptr) (Monitor, error) {
	hMonitor, err := z.getMonitorHandle(hwnd)
	if err != nil {
		return Monitor{}, err
	}
	m, err := z.getMonitorByHandle(hMonitor)
	if err != nil {
		return Monitor{}, err
	}
	return m, nil
}

func (z *System) GetMonitorByNumber(number int) (Monitor, error) {
	monitors, err := z.GetMonitors()
	if err != nil {
		return Monitor{}, err
	}
	for _, m := range monitors {
		if m.Number == number {
			return m, nil
		}
	}
	return Monitor{}, errors.New("monitor not found")
}
