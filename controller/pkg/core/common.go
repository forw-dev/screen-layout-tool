package core

type XYWH struct {
	X      *float64 `json:"x"`
	Y      *float64 `json:"y"`
	Width  *float64 `json:"width"`
	Height *float64 `json:"height"`
}

type TBLR struct {
	Top    *float64 `json:"top"`
	Bottom *float64 `json:"bottom"`
	Left   *float64 `json:"left"`
	Right  *float64 `json:"right"`
}

type Geometry struct {
	XYWH
	TBLR
}

type Position struct {
	MonitorNumber *int `json:"monitor"`
	WindowHandle  uintptr
	Geometry
}

type Layout struct {
	Author      string     `json:"author"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Positions   []Position `json:"positions"`
}
