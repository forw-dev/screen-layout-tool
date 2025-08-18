package main

import (
	"fmt"
	"screen-layout-controller/pkg/core"
	"strconv"

	"github.com/spf13/cobra"
)

type Resize struct {
	width  string
	height string
	locate int
}

var resize Resize

var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "resize window",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := core.New()
		if err != nil {
			errExit(err)
		}
		if resize.locate != 0 && resize.locate != 3 {
			fmt.Println("locate must be 0, 3")
			cmd.Help()
			return
		}
		if resize.width == "none" && resize.height == "none" {
			fmt.Println("at least one width/height")
			cmd.Help()
			return
		}
		if resize.width != "none" {
			width, err := strconv.ParseFloat(resize.width, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Resize, core.Width, resize.locate, width)
			if err != nil {
				errExit(err)
			}
		}
		if resize.height != "none" {
			height, err := strconv.ParseFloat(resize.height, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Resize, core.Height, resize.locate, height)
			if err != nil {
				errExit(err)
			}
		}
	},
}

func init() {
	resizeCmd.Flags().StringVarP(&resize.width, "width", "w", "none", "offset: -100-100 others: 0-100")
	resizeCmd.Flags().StringVarP(&resize.height, "height", "e", "none", "offset: -100-100 others: 0-100")
	resizeCmd.Flags().IntVarP(&resize.locate, "locate", "o", 0, "0: offset, 3: fix")
	resizeCmd.Flags().SortFlags = false
}
