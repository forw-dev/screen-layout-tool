package system

import (
	"syscall"
)

type RectRaw struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type MonitorRaw struct {
	CbSize    uint32 // 结构体大小
	RcMonitor RectRaw
	RcWork    RectRaw
	DwFlags   uint32
	SzDevice  [32]uint16
}

type XYWH struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type TBLR struct {
	Top    int32
	Bottom int32
	Left   int32
	Right  int32
}

type Geometry struct {
	XYWH
	TBLR
}

type Window struct {
	Handle uintptr
	Geometry
	// MonitorHandle syscall.Handle
	MonitorNumber int
}

type Monitor struct {
	DeviceName string
	IsPrimary  bool
	Screen     Geometry
	Desktop    Geometry
	Handle     syscall.Handle
	Number     int
}
