package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "screen-layout-tool-controller",
	Short: "This is a window controller.",
}

func main() {
	// appDataDir
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("get user cache dir failed: ", err)
		return
	}
	appDataDir := filepath.Join(cacheDir, "screen-layout-tool-controller")
	err = os.MkdirAll(appDataDir, 0755)
	if err != nil {
		fmt.Println("create app data dir failed: ", err)
		return
	}
	// logFile
	logFilePath := filepath.Join(appDataDir, "controller.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open log file failed: ", err)
		log.SetOutput(os.Stderr)
	}
	defer logFile.Close()
	// log.SetOutput(logFile)
	log.SetOutput(io.MultiWriter(logFile, os.Stderr))
	log.SetFlags(log.Ldate | log.Ltime)

	rootCmd.AddCommand(windowCmd)
	rootCmd.AddCommand(monitorCmd)
	cobra.EnableCommandSorting = false

	err = rootCmd.Execute()
	if err != nil {
		errExit(err)
	}
}

func errExit(err error) {
	log.Println(err)
	if strings.Contains(err.Error(), "position result unexpected") {
		os.Exit(11)
	}
	if strings.Contains(err.Error(), "layout file parse err") {
		os.Exit(12)
	}
	if strings.Contains(err.Error(), "open layout.json: The system cannot find") {
		os.Exit(13)
	}
	if strings.Contains(err.Error(), "toIndex: index not support") {
		os.Exit(14)
	}
	if strings.Contains(err.Error(), "duplicate positions in layout") {
		os.Exit(15)
	}
	if strings.Contains(err.Error(), "monitor not found") {
		os.Exit(16)
	}
	if strings.Contains(err.Error(), "load layout file failed") {
		os.Exit(17)
	}
	os.Exit(1)
}
