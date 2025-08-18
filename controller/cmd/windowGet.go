package main

import (
	"fmt"
	"screen-layout-controller/pkg/core"

	"github.com/spf13/cobra"
)

var windowGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get infomation",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := core.New()
		if err != nil {
			errExit(err)
		}
		fmt.Println()
		fmt.Println("[Current Window Infomation]")
		err = c.PrintPositionInfo()
		if err != nil {
			errExit(err)
		}
	},
}

func init() {
	windowGetCmd.Flags().SortFlags = false
}
