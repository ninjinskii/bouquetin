package cmd

import (
	"fmt"
	"ninjinski/bouquetin/core"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const URL_PULL = URL_BASE + "pull"

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
}

func pull(cmd *cobra.Command, args []string) {
	historyPosition, _ := cmd.Flags().GetString("history")
	position, err := strconv.Atoi(historyPosition)
	fileHandler := &core.GoFileHandler{}
	filepath := core.GetEnvironmentVariable(core.ENV_FILEPATH)

	headers := getHttpHeaders(fileHandler, filepath, position)

	if err != nil {
		position = 0
	}

	httpClient := &core.GoHttpClient{}
	response, statusCode := core.HttpClient.Get(httpClient, URL_PULL, headers)

	if statusCode == 200 {
		writeBackupFile(fileHandler, filepath)
		fileHandler.WriteFile(filepath, response)
		color.Green("PULL: SUCCESS")
		return
	}

	color.Red("PULL: ERROR ->")
	color.Red(string(statusCode))
	fmt.Println(response)
}

func getHttpHeaders(fileHandler core.FileHandler, filepath string, position int) core.BqtHttpHeaders {
	filename := core.FileHandler.GetFilename(fileHandler, filepath)

	userId := core.GetEnvironmentVariable(core.ENV_USER_ID)
	headers := core.BqtHttpHeaders{
		UserId:            userId,
		PreferredFilename: filename,
		Position:          position,
	}

	return headers
}

func writeBackupFile(fileHandler core.FileHandler, filepath string) {
	backupPath := filepath + ".old"
	fileHandler.CopyFile(filepath, backupPath)
}
