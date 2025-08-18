package main

import (
	"fmt"
	"screen-layout-controller/pkg/core"
	"strconv"

	"github.com/spf13/cobra"
)

type Pull struct {
	top    string
	bottom string
	left   string
	right  string
	locate int
}

var pull Pull

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull window",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := core.New()
		if err != nil {
			errExit(err)
		}
		if pull.locate < 0 || pull.locate > 2 {
			fmt.Println("locate must be 0, 1, 2")
			cmd.Help()
			return
		}
		if pull.top == "none" && pull.bottom == "none" && pull.left == "none" && pull.right == "none" {
			fmt.Println("at least one top/bottom/left/right")
			cmd.Help()
			return
		}
		if pull.top != "none" {
			top, err := strconv.ParseFloat(pull.top, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Pull, core.Top, pull.locate, top)
			if err != nil {
				errExit(err)
			}
		}
		if pull.bottom != "none" {
			bottom, err := strconv.ParseFloat(pull.bottom, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Pull, core.Bottom, pull.locate, bottom)
			if err != nil {
				errExit(err)
			}
		}
		if pull.left != "none" {
			left, err := strconv.ParseFloat(pull.left, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Pull, core.Left, pull.locate, left)
			if err != nil {
				errExit(err)
			}
		}
		if pull.right != "none" {
			right, err := strconv.ParseFloat(pull.right, 64)
			if err != nil {
				errExit(err)
			}
			err = c.ChangeGeometry(core.Pull, core.Right, pull.locate, right)
			if err != nil {
				errExit(err)
			}
		}
	},
}

func init() {
	pullCmd.Flags().StringVarP(&pull.top, "top", "t", "none", "offset: -100-100 others: 0-100")
	pullCmd.Flags().StringVarP(&pull.bottom, "bottom", "b", "none", "offset: -100-100 others: 0-100")
	pullCmd.Flags().StringVarP(&pull.left, "left", "l", "none", "offset: -100-100 others: 0-100")
	pullCmd.Flags().StringVarP(&pull.right, "right", "r", "none", "offset: -100-100 others: 0-100")
	pullCmd.Flags().IntVarP(&pull.locate, "locate", "o", 0, "0: offset, 1: coordinate, 2: margin to desktop side")
	pullCmd.Flags().SortFlags = false
}
