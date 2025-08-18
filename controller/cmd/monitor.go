package main

import (
	"fmt"
	"screen-layout-controller/pkg/core"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor in system",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := core.New()
		if err != nil {
			errExit(err)
		}
		fmt.Println()
		fmt.Println("[Current Window Infomation]")
		err = c.PrintMonitorInfos()
		if err != nil {
			errExit(err)
		}
	},
}

func init() {
	monitorCmd.Flags().SortFlags = false
}
