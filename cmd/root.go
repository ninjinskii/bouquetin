package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// TODO: adapt to production code
const URL_BASE = "https://bouquetin.njk.localhost/"

var rootCmd = &cobra.Command{
	Use:   "bqt",
	Short: "File transfer manager",
	Long:  "Upload / download files from a server based on your environment variables",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
