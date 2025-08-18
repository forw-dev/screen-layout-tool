package core

import (
	"fmt"
	"testing"
	"time"
	// "time"
	// "golang.org/x/sys/windows"
)

func Test(t *testing.T) {
	w, err := New()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 1)

	// mons, err := w.GetMonitors()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// print(mons)

	// pos, err := w.GetPosition(w.hwnd)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// print(pos)

	// w.ChangeMonitor(1)

	// move
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Left, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Top, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Right, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Bottom, 10, Offset)

	// time.Sleep(time.Second * 1)
	// w.Change(Move, Left, 10, Coordinate)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Top, 10, Coordinate)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Right, 90, Coordinate)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Bottom, 90, Coordinate)

	// time.Sleep(time.Second * 1)
	// w.Change(Move, Left, 5, Margin)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Top, 5, Margin)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Right, 5, Margin)
	// time.Sleep(time.Second * 1)
	// w.Change(Move, Bottom, 5, Margin)

	// pull
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Left, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Top, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Right, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Bottom, 10, Offset)

	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Left, 10, Coordinate)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Top, 10, Coordinate)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Right, 90, Coordinate)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Bottom, 90, Coordinate)

	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Left, 5, Margin)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Top, 5, Margin)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Right, 5, Margin)
	// time.Sleep(time.Second * 1)
	// w.Change(Pull, Bottom, 5, Margin)

	// resize
	// time.Sleep(time.Second * 1)
	// w.Change(Resize, Width, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Resize, Height, 10, Offset)
	// time.Sleep(time.Second * 1)
	// w.Change(Resize, Width, 60, Fix)
	// time.Sleep(time.Second * 1)
	// w.Change(Resize, Height, 60, Fix)

	// fmt.Println(w.ToIndex(10, "layout.json"))

	// for i := 0; i <= 12; i++ {
	// 	time.Sleep(time.Second * 1)
	// 	fmt.Println("to:", i)
	// 	fmt.Println(w.ToIndex(i, "layout.json"))
	// }

	// time.Sleep(time.Second * 5)
	// fmt.Println(w.ToClosest("layout.json"))

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(w.Scroll(1, true, "layout.json"))
	}
}
