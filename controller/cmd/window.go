package main

import (
	"github.com/spf13/cobra"
)

type Window struct {
}

var windowCmd = &cobra.Command{
	Use:   "window",
	Short: "window on system",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	windowCmd.AddCommand(windowGetCmd)
	windowCmd.AddCommand(jumpCmd)
	windowCmd.AddCommand(moveCmd)
	windowCmd.AddCommand(pullCmd)
	windowCmd.AddCommand(resizeCmd)
	windowCmd.AddCommand(scrollCmd)
	windowCmd.Flags().SortFlags = false
}
