package core

import (
	"fmt"
	"screen-layout-controller/pkg/system"
)

func (z *Core) getMonitors() ([]system.Monitor, error) {
	ms, err := z.System.GetMonitors()
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (z *Core) getPosition(hwnd uintptr) (Position, error) {
	win, err := z.System.GetWindow(hwnd)
	if err != nil {
		return Position{}, err
	}
	mon, err := z.System.GetMonitor(hwnd)
	if err != nil {
		return Position{}, err
	}
	var pos Position
	x := float64(win.X) / float64(mon.Desktop.Width) * 100
	y := float64(win.Y) / float64(mon.Desktop.Height) * 100
	w := float64(win.Width) / float64(mon.Desktop.Width) * 100
	h := float64(win.Height) / float64(mon.Desktop.Height) * 100
	t := float64(win.Top) / float64(mon.Desktop.Height) * 100
	b := float64(win.Bottom) / float64(mon.Desktop.Height) * 100
	l := float64(win.Left) / float64(mon.Desktop.Width) * 100
	r := float64(win.Right) / float64(mon.Desktop.Width) * 100
	pos.X = &x
	pos.Y = &y
	pos.Width = &w
	pos.Height = &h
	pos.Top = &t
	pos.Bottom = &b
	pos.Left = &l
	pos.Right = &r
	pos.MonitorNumber = &mon.Number
	pos.WindowHandle = hwnd
	return pos, nil
}

func (z *Core) PrintPositionInfo() error {
	p, err := z.getPosition(z.hwnd)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Monitor\t", *p.MonitorNumber)
	fmt.Println("Handle\t", p.WindowHandle)
	fmt.Println("Top\t", *p.Top)
	fmt.Println("Bottom\t", *p.Bottom)
	fmt.Println("Left\t", *p.Left)
	fmt.Println("Right\t", *p.Right)
	fmt.Println("X\t", *p.X)
	fmt.Println("Y\t", *p.Y)
	fmt.Println("Width\t", *p.Width)
	fmt.Println("Height\t", *p.Height)
	fmt.Println()
	return nil
}

func (z *Core) printMonitorInfo(monitor system.Monitor) {
	fmt.Println()
	fmt.Println("Number\t\t", monitor.Number)
	fmt.Println("Monitor handle\t", monitor.Handle)
	fmt.Println("DeviceName\t", monitor.DeviceName)
	fmt.Println("IsPrimary\t", monitor.IsPrimary)
	fmt.Println("Desktop top\t", monitor.Desktop.Top)
	fmt.Println("Desktop left\t", monitor.Desktop.Left)
	fmt.Println("Desktop right\t", monitor.Desktop.Right)
	fmt.Println("Desktop bottom\t", monitor.Desktop.Bottom)
	fmt.Println("Screen top\t", monitor.Screen.Top)
	fmt.Println("Screen left\t", monitor.Screen.Left)
	fmt.Println("Screen right\t", monitor.Screen.Right)
	fmt.Println("Screen bottom\t", monitor.Screen.Bottom)
	fmt.Println()
}

func (z *Core) PrintMonitorInfos() error {
	monitors, err := z.getMonitors()
	if err != nil {
		return err
	}
	for _, m := range monitors {
		z.printMonitorInfo(m)
	}
	return nil
}
