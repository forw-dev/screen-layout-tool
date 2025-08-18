package core

import (
	"errors"
)

const (
	Move   = 1 // 移动
	Pull   = 2 // 拉取
	Resize = 3 // 调整大小

	Top    = 1 // 上
	Bottom = 2 // 下
	Left   = 3 // 左
	Right  = 4 // 右
	Width  = 5 // 宽度
	Height = 6 // 高度

	Offset     = 0 // 自身定位 -100~100
	Coordinate = 1 // 坐标定位 0~100
	Margin     = 2 // 边缘定位 0~100
	Fix        = 3 // 固定值 0~100
)

func (z *Core) setGeometry(hwnd uintptr, action, field, locate int, value float64) error {
	// 验证和处理输入
	if locate == Coordinate && (value < 0 || value > 100) {
		return errors.New("value not support")
	}
	if locate == Margin && (value < 0 || value > 100) {
		return errors.New("value not support")
	}
	if locate == Offset && (value < -100 || value > 100) {
		return errors.New("value not support")
	}
	// 默认值
	win, err := z.System.GetWindow(hwnd)
	if err != nil {
		return err
	}
	tx := win.X
	ty := win.Y
	tw := win.Width
	th := win.Height
	// 变换
	var ai int32
	mon, err := z.System.GetMonitor(hwnd)
	if err != nil {
		return err
	}
	if field == Left || field == Right || field == Width {
		ai = int32(value * float64(mon.Desktop.Width) / 100)
	}
	if field == Top || field == Bottom || field == Height {
		ai = int32(value * float64(mon.Desktop.Height) / 100)
	}
	// move
	// move left
	if action == Move && field == Left && locate == Offset {
		tx = win.X - ai
	}
	if action == Move && field == Left && locate == Coordinate {
		tx = ai
	}
	if action == Move && field == Left && locate == Margin {
		tx = ai
	}
	// move top
	if action == Move && field == Top && locate == Offset {
		ty = win.Y - ai
	}
	if action == Move && field == Top && locate == Coordinate {
		ty = ai
	}
	if action == Move && field == Top && locate == Margin {
		ty = ai
	}
	// move right
	if action == Move && field == Right && locate == Offset {
		tx = win.X + ai
	}
	if action == Move && field == Right && locate == Coordinate {
		tx = ai - win.Width
	}
	if action == Move && field == Right && locate == Margin {
		tx = mon.Desktop.Width - ai - win.Width
	}
	// move bottom
	if action == Move && field == Bottom && locate == Offset {
		ty = win.Y + ai
	}
	if action == Move && field == Bottom && locate == Coordinate {
		ty = ai - win.Height
	}
	if action == Move && field == Bottom && locate == Margin {
		ty = mon.Desktop.Height - ai - win.Height
	}
	// pull
	// 对边不动
	// pull left
	if action == Pull && field == Left && locate == Offset {
		tx = win.X - ai
		tw = win.Width + ai
	}
	if action == Pull && field == Left && locate == Coordinate {
		tx = ai
		tw = win.Right - ai
	}
	if action == Pull && field == Left && locate == Margin {
		tx = ai
		tw = win.Width + win.Left - ai
	}
	// pull top
	if action == Pull && field == Top && locate == Offset {
		ty = win.Y - ai
		th = win.Height + ai
	}
	if action == Pull && field == Top && locate == Coordinate {
		ty = ai
		th = win.Bottom - ai
	}
	if action == Pull && field == Top && locate == Margin {
		ty = ai
		th = win.Height + win.Top - ai
	}
	// pull rihgt
	if action == Pull && field == Right && locate == Offset {
		tw = win.Width + ai
	}
	if action == Pull && field == Right && locate == Coordinate {
		tw = ai - win.Left
	}
	if action == Pull && field == Right && locate == Margin {
		tw = mon.Desktop.Width - ai - win.Left
	}
	// pull bottom
	if action == Pull && field == Bottom && locate == Offset {
		th = win.Height + ai
	}
	if action == Pull && field == Bottom && locate == Coordinate {
		th = ai - win.Top
	}
	if action == Pull && field == Bottom && locate == Margin {
		th = mon.Desktop.Height - ai - win.Top
	}
	// resize
	// 基于中点
	if action == Resize && field == Width && locate == Offset {
		tw = win.Width + ai
		tx = win.X - ai/2
	}
	if action == Resize && field == Width && locate == Fix {
		tw = ai
		tx = win.X - (tw-win.Width)/2
	}
	if action == Resize && field == Height && locate == Offset {
		th = win.Height + ai
		ty = win.Y - ai/2
	}
	if action == Resize && field == Height && locate == Fix {
		th = ai
		ty = win.Y - (th-win.Height)/2
	}
	// apply
	err = z.System.SetWindow(hwnd, tx, ty, tw, th, mon.Number)
	if err != nil {
		return err
	}
	return nil
}

func (z *Core) ChangeGeometry(action, field, locate int, value float64) error {
	return z.setGeometry(z.hwnd, action, field, locate, value)
}

func (z *Core) setMonitor(hwnd uintptr, m int) error {
	if m < 0 || m > 100 {
		return errors.New("monitor not support")
	}
	if m == 0 {
		primaryMonitorNumber, err := z.System.GetPrimaryMonitorNumber()
		if err != nil {
			return err
		}
		m = primaryMonitorNumber
	}
	win, err := z.System.GetWindow(hwnd)
	if err != nil {
		return err
	}
	mon, err := z.System.GetMonitor(hwnd)
	if err != nil {
		return err
	}
	targetMonitor, err := z.System.GetMonitorByNumber(m)
	if err != nil {
		return err
	}
	tx := win.X * targetMonitor.Desktop.Width / mon.Desktop.Width
	ty := win.Y * targetMonitor.Desktop.Height / mon.Desktop.Height
	tw := win.Width * targetMonitor.Desktop.Width / mon.Desktop.Width
	th := win.Height * targetMonitor.Desktop.Height / mon.Desktop.Height
	err = z.System.SetWindow(hwnd, tx, ty, tw, th, m)
	if err != nil {
		return err
	}
	return nil
}

func (z *Core) SetMonitor(m int) error {
	return z.setMonitor(z.hwnd, m)
}
