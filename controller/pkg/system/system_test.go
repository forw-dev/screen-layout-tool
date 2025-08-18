package system

import (
	"fmt"
	"testing"
	// "time"
	// "golang.org/x/sys/windows"
)

func Test(t *testing.T) {
	s, err := New()
	if err != nil {
		fmt.Println(err)
	}

	monitors, err := s.GetMonitors()
	if err != nil {
		fmt.Println(err)
	}
	print(monitors)

	// primaryMonitorNumber, err := s.GetPrimaryMonitorNumber()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("primaryMonitorNumber:", primaryMonitorNumber)

	// hwnd, err := s.GetCurrentWindowHandle()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("hwnd:", hwnd)

	// w, err := s.GetWindow(hwnd)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// print(w)

	// err = s.SetWindow(hwnd, 0, 0, 500, 500, 1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
