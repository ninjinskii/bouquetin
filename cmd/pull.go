package cmd

import (
	"fmt"
	"ninjinski/bouquetin/core"
	"strconv"

	"github.com/spf13/cobra"
)

const URL_PULL = "/pull"

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Download the file",
	Long:  "Download the file based on your environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		pull(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().String("history", "", "Choose a previous version of the file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func pull(cmd *cobra.Command, args []string) {
	historyPosition, _ := cmd.Flags().GetString("history")
	position, err := strconv.Atoi(historyPosition)
	filepath := core.GetEnvironmentVariable(core.ENV_FILEPATH)
	userId := core.GetEnvironmentVariable(core.ENV_USER_ID)
	httpClient := core.NewHttpClient()
	headers := core.BqtHttpHeaders{
		UserId:            userId,
		PreferredFilename: filepath, // TODO: get only filename instead of filepath
	}

	if err != nil {
		position = 0
	}

	response = core.HttpClient.Get(httpClient, URL_PULL, headers)

	// get user id & filename -
	// make a request with user-id & preferred-filename header -
	// download file as http content by sending request -
	// read file from http content -
	// create/replace the .old file
	// replace the current file with the new one

	fmt.Println(position)
}
