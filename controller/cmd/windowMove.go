package main

import (
	"fmt"
	"screen-layout-controller/pkg/core"
	"strconv"

	"github.com/spf13/cobra"
)

type Move struct {
	monitor string
	top     string
	bottom  string
	left    string
	right   string
	locate  int
}

var move Move

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "move window",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := core.New()
		if err != nil {
			errExit(err)
		}
		if move.locate < 0 || move.locate > 2 {
			fmt.Println("locate must be 0, 1, 2")
			cmd.Help()
			return
		}
		if move.monitor == "none" && move.top == "none" && move.bottom == "none" && move.left == "none" && move.right == "none" {
			fmt.Println("at least one monitor/top/bottom/left/right")
			cmd.Help()
			return
		}
		if move.monitor != "none" {
			monitor, err := strconv.Atoi(move.monitor)
			if err != nil {
				errExit(err)
			}
			err = c.SetMonitor(monitor)
			if err != nil {
				errExit(err)
			}
		}
		if move.top != "none" {
			top, err := strconv.ParseFloat(move.top, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Move, core.Top, move.locate, top)
			if err != nil {
				errExit(err)
			}
		}
		if move.bottom != "none" {
			bottom, err := strconv.ParseFloat(move.bottom, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Move, core.Bottom, move.locate, bottom)
			if err != nil {
				errExit(err)
			}
		}
		if move.left != "none" {
			left, err := strconv.ParseFloat(move.left, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Move, core.Left, move.locate, left)
			if err != nil {
				errExit(err)
			}
		}
		if move.right != "none" {
			right, err := strconv.ParseFloat(move.right, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Move, core.Right, move.locate, right)
			if err != nil {
				errExit(err)
			}
		}
	},
}

func init() {
	moveCmd.Flags().StringVarP(&move.monitor, "monitor", "m", "none", "monitor number, 0-n, 0 is primary")
	moveCmd.Flags().StringVarP(&move.top, "top", "t", "none", "offset: -100-100 others: 0-100")
	moveCmd.Flags().StringVarP(&move.bottom, "bottom", "b", "none", "offset: -100-100 others: 0-100")
	moveCmd.Flags().StringVarP(&move.left, "left", "l", "none", "offset: -100-100 others: 0-100")
	moveCmd.Flags().StringVarP(&move.right, "right", "r", "none", "offset: -100-100 others: 0-100")
	moveCmd.Flags().IntVarP(&move.locate, "locate", "o", 0, "0: offset, 1: coordinate, 2: margin to desktop side")
	moveCmd.Flags().SortFlags = false
}
