package cmd

import (
	"fmt"
	"ninjinski/bouquetin/core"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const URL_PUSH = URL_BASE + "push"

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Upload the file",
	Long:  "Upload the file based on your environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		push(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func push(cmd *cobra.Command, args []string) {
	filepath := core.GetEnvironmentVariable(core.ENV_FILEPATH)
	userId := core.GetEnvironmentVariable(core.ENV_USER_ID)
	httpClient := &core.GoHttpClient{}
	headers := core.BqtHttpHeaders{
		PreferredFilename: filepath,
		UserId:            userId,
	}

	response, statusCode := core.HttpClient.Multipart(httpClient, URL_PUSH, filepath, headers)

	if statusCode == 200 {
		isNewUser := response != ""

		if isNewUser {
			promptUserToSaveId(response)
		}

		color.Green("PUSH: SUCCESS")
		return
	}

	color.Red("PUSH: ERROR ->")
	color.Red(string(statusCode))
	fmt.Println(response)
}

func promptUserToSaveId(userId string) {
	fmt.Println("#################################")
	color.Red("YOU NEED TO SAVE YOUR ID BY DOING THE FOLLOWING:")
	fmt.Println("Open a terminal, then:")
	fmt.Println("Windows: set BOUQUETIN_ID=" + userId)
	fmt.Println("Linux/Android: export BOUQUETIN_ID=" + userId)
	fmt.Println("#################################")
}
