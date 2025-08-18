package core

import (
	"screen-layout-controller/pkg/system"
)

type Core struct {
	hwnd   uintptr
	System system.System
}

func New() (*Core, error) {
	var c Core
	hwnd, err := c.System.GetCurrentWindowHandle()
	if err != nil {
		return nil, err
	}
	c.hwnd = hwnd
	return &c, nil
}
