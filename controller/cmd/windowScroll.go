package main

import (
	"screen-layout-controller/pkg/core"
	"strconv"

	"github.com/spf13/cobra"
)

type Scroll struct {
	layout string
	cross  bool
	step   string
}

var scroll Scroll

var scrollCmd = &cobra.Command{
	Use:   "scroll",
	Short: "scroll in layout",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := core.New()
		if err != nil {
			errExit(err)
		}
		if scroll.step == "none" {
			err = c.ToClosest(scroll.layout)
			if err != nil {
				errExit(err)
			}
			return
		}
		if scroll.step != "none" {
			step, err := strconv.Atoi(scroll.step)
			if err != nil {
				errExit(err)
			}
			err = c.Scroll(step, scroll.cross, scroll.layout)
			if err != nil {
				errExit(err)
			}
			return
		}
	},
}

func init() {
	scrollCmd.Flags().StringVarP(&scroll.layout, "layout", "l", "Default.json", "layout file")
	scrollCmd.Flags().BoolVarP(&scroll.cross, "cross", "c", false, "cross monitors")
	scrollCmd.Flags().StringVarP(&scroll.step, "step", "s", "none", "scroll step, none is closest")
	scrollCmd.Flags().SortFlags = false
}
