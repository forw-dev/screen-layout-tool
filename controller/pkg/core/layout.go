package core

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"strconv"
)

func (z *Core) loadLayoutFile(layoutFilePath string) ([]Position, error) {
	j, err := os.ReadFile(layoutFilePath)
	if err != nil {
		return nil, errors.New("load layout file failed: " + err.Error())
	}
	var layout Layout
	err = json.Unmarshal(j, &layout)
	if err != nil {
		return nil, errors.New("layout file parse err")
	}
	if len(layout.Positions) == 0 {
		return nil, errors.New("layout is empty")
	}
	primaryMonitor, err := z.System.GetPrimaryMonitorNumber()
	if err != nil {
		return nil, err
	}
	var ps []Position
	for i, p := range layout.Positions {
		// 处理主显示器
		if p.MonitorNumber == nil || *p.MonitorNumber == 0 {
			p.MonitorNumber = &primaryMonitor
		}
		// 参数转换
		//// 输入 TBLRWH，且 TBLR 为边缘定位，转换为 XYWH
		if p.X == nil {
			if p.Right != nil && p.Width != nil {
				x := 100 - *p.Right - *p.Width
				p.X = &x
			} else if p.Left != nil {
				p.X = p.Left
			} else {
				return nil, errors.New("layout position err: " + strconv.Itoa(i))
			}
		}
		if p.Y == nil {
			if p.Bottom != nil && p.Height != nil {
				y := 100 - *p.Bottom - *p.Height
				p.Y = &y
			} else if p.Top != nil {
				p.Y = p.Top
			} else {
				return nil, errors.New("layout position err: " + strconv.Itoa(i))
			}
		}
		if p.Width == nil {
			if p.Left != nil && p.Right != nil {
				w := 100 - *p.Right - *p.Left
				p.Width = &w
			} else if p.X != nil && p.Right != nil {
				w := 100 - *p.Right - *p.X
				p.Width = &w
			} else {
				return nil, errors.New("layout position err: " + strconv.Itoa(i))
			}
		}
		if p.Height == nil {
			if p.Top != nil && p.Bottom != nil {
				h := 100 - *p.Bottom - *p.Top
				p.Height = &h
			} else if p.Y != nil && p.Bottom != nil {
				h := 100 - *p.Bottom - *p.Y
				p.Height = &h
			} else {
				return nil, errors.New("layout position err: " + strconv.Itoa(i))
			}
		}
		/*
			//// 输入 TBLRWH，且 TBLR 为原点定位，转换为 XYWH
			if p.X == nil {
				if p.Right != nil && p.Width != nil {
					x := *p.Right - *p.Width
					p.X = &x
				} else if p.Left != nil {
					p.X = p.Left
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
			if p.Y == nil {
				if p.Bottom != nil && p.Height != nil {
					y := *p.Bottom - *p.Height
					p.Y = &y
				} else if p.Top != nil {
					p.Y = p.Top
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
			if p.Width == nil {
				if p.Left != nil && p.Right != nil {
					w := *p.Right - *p.Left
					p.Width = &w
				} else if p.X != nil && p.Right != nil {
					w := *p.Right - *p.X
					p.Width = &w
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
			if p.Height == nil {
				if p.Top != nil && p.Bottom != nil {
					h := *p.Bottom - *p.Top
					p.Height = &h
				} else if p.Y != nil && p.Bottom != nil {
					h := *p.Bottom - *p.Y
					p.Height = &h
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
		*/

		/*
			//// 输入 TBLRWH，且 TBLR 为原点定位，补齐 TBLR。最终结论是没必要补齐。
			if p.Top == nil {
				if p.Height != nil && p.Bottom != nil {
					t := *p.Bottom - *p.Height
					p.Top = &t
				} else if p.Y != nil {
					t := *p.Y
					p.Top = &t
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
			if p.Bottom == nil {
				if p.Height != nil && p.Top != nil {
					b := *p.Top + *p.Height
					p.Bottom = &b
				} else if p.Y != nil && p.Height != nil {
					b := *p.Y + *p.Height
					p.Bottom = &b
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
			if p.Left == nil {
				if p.Width != nil && p.Right != nil {
					l := *p.Right - *p.Width
					p.Left = &l
				} else if p.X != nil {
					l := *p.X
					p.Left = &l
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
			if p.Right == nil {
				if p.Left != nil && p.Width != nil {
					r := *p.Left + *p.Width
					p.Right = &r
				} else if p.X != nil && p.Width != nil {
					r := *p.Width + *p.X
					p.Right = &r
				} else {
					return nil, errors.New("layout position err: " + strconv.Itoa(i))
				}
			}
		*/

		ps = append(ps, p)
	}
	// 必须先补齐几何数据、升级显示器数据后再检查是否有重复
	if z.hasDuplicates(ps) {
		return nil, errors.New("duplicate positions in layout")
	}
	return ps, nil
}

func (z *Core) hasDuplicates(positions []Position) bool {
	seen := map[string]bool{}
	for _, p := range positions {
		m := strconv.Itoa(*p.MonitorNumber)
		x := strconv.FormatFloat(*p.X, 'f', -1, 64)
		y := strconv.FormatFloat(*p.Y, 'f', -1, 64)
		w := strconv.FormatFloat(*p.Width, 'f', -1, 64)
		h := strconv.FormatFloat(*p.Height, 'f', -1, 64)
		key := m + x + y + w + h
		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

func (z *Core) toIndex(hwnd uintptr, index int, positions []Position) error {
	if index < 0 || index > len(positions)-1 {
		return errors.New("toIndex: index not support")
	}
	p := positions[index]
	targetMonitor, err := z.System.GetMonitorByNumber(*p.MonitorNumber)
	if err != nil {
		return err
	}
	tx := int32(*p.X * float64(targetMonitor.Desktop.Width) / 100)
	ty := int32(*p.Y * float64(targetMonitor.Desktop.Height) / 100)
	tw := int32(*p.Width * float64(targetMonitor.Desktop.Width) / 100)
	th := int32(*p.Height * float64(targetMonitor.Desktop.Height) / 100)
	err = z.System.SetWindow(hwnd, tx, ty, tw, th, *p.MonitorNumber)
	if err != nil {
		return err
	}
	return nil
}
func (z *Core) toClosest(hwnd uintptr, positions []Position) (isRight bool, index int, err error) {
	pos, err := z.getPosition(hwnd)
	if err != nil {
		return false, 0, err
	}
	minDistance := math.MaxFloat64
	minIndex := 0
	isRight = false
	for i, p := range positions {
		if *p.MonitorNumber != *pos.MonitorNumber {
			continue
		}
		var dx, dy, dw, dh float64
		if p.X != nil {
			dx = math.Abs(*p.X - *pos.X)
		}
		if p.Y != nil {
			dy = math.Abs(*p.Y - *pos.Y)
		}
		if p.Width != nil {
			dw = math.Abs(*p.Width - *pos.Width)
		}
		if p.Height != nil {
			dh = math.Abs(*p.Height - *pos.Height)
		}
		distance := dx + dy + dw + dh
		// fmt.Println("distance:", distance)
		if distance < 0.5 {
			return true, i, nil
		}
		if distance < minDistance {
			minDistance = distance
			minIndex = i
		}
	}
	index = minIndex
	err = z.toIndex(z.hwnd, index, positions)
	if err != nil {
		return false, index, err
	}
	return false, index, nil
}
func (z *Core) scrollIn(hwnd uintptr, step int, crossMonitor bool, positions []Position) error {
	mon, err := z.System.GetMonitor(hwnd)
	if err != nil {
		return err
	}
	var ps []Position
	if crossMonitor {
		ps = positions
	} else {
		for _, p := range positions {
			if *p.MonitorNumber == mon.Number {
				ps = append(ps, p)
			}
		}
	}
	isRight, index, err := z.toClosest(z.hwnd, ps)
	if err != nil {
		return err
	}
	if isRight {
		targetIndex := index + step
		if targetIndex >= len(ps) {
			targetIndex = 0
		}
		if targetIndex < 0 {
			targetIndex = len(ps) - 1
		}
		err = z.toIndex(z.hwnd, targetIndex, ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *Core) ToIndex(index int, layoutFilePath string) error {
	positions, err := z.loadLayoutFile(layoutFilePath)
	if err != nil {
		return err
	}
	err = z.toIndex(z.hwnd, index, positions)
	if err != nil {
		return err
	}
	return nil
}
func (z *Core) ToClosest(layoutFilePath string) error {
	positions, err := z.loadLayoutFile(layoutFilePath)
	if err != nil {
		return err
	}
	_, _, err = z.toClosest(z.hwnd, positions)
	if err != nil {
		return err
	}
	return nil
}
func (z *Core) Scroll(step int, crossMonitor bool, layoutFilePath string) error {
	positions, err := z.loadLayoutFile(layoutFilePath)
	if err != nil {
		return err
	}
	err = z.scrollIn(z.hwnd, step, crossMonitor, positions)
	if err != nil {
		return err
	}
	return nil
}

// func print(i interface{}) {
// 	j, err := json.MarshalIndent(i, "", "  ") // 使用两个空格作为缩进
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println(string(j))
// }
