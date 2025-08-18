package main

import (
	"screen-layout-controller/pkg/core"
	"strconv"

	"github.com/spf13/cobra"
)

type Jump struct {
	layout string
	index  string
}

var jump Jump

var jumpCmd = &cobra.Command{
	Use:   "jump",
	Short: "jump in layout",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := core.New()
		if err != nil {
			errExit(err)
		}
		if jump.index == "none" {
			err = c.ToClosest(jump.layout)
			if err != nil {
				errExit(err)
			}
			return
		}
		if jump.index != "none" {
			index, err := strconv.Atoi(jump.index)
			if err != nil {
				errExit(err)
			}
			err = c.ToIndex(index, jump.layout)
			if err != nil {
				errExit(err)
			}
			return
		}
	},
}

func init() {
	jumpCmd.Flags().StringVarP(&jump.layout, "layout", "l", "layout.json", "layout file")
	jumpCmd.Flags().StringVarP(&jump.index, "index", "i", "none", "index to jump, 0-n, none is closest")
	jumpCmd.Flags().SortFlags = false
}
