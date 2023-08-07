package cmd

import (
	"encoding/json"
	"fmt"
	"ninjinski/bouquetin/core"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const URL_STATUS = URL_BASE + "status"

type StatusDto struct {
	FileHash         string
	UserCreationDate int64
	LastUploadDate   int64
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays your file status",
	Long:  "Show if your file is synced with the server. If not, your local file might be older OR newer, consider pushing or pulling right away",
	Run: func(cmd *cobra.Command, args []string) {
		status(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func status(cmd *cobra.Command, args []string) {
	fileHandler := &core.GoFileHandler{}
	filepath := core.GetEnvironmentVariable(core.ENV_FILEPATH)

	headers := getSimpleHttpHeaders(fileHandler, filepath)

	httpClient := &core.GoHttpClient{}
	response, statusCode := core.HttpClient.Get(httpClient, URL_STATUS, headers)

	dto := StatusDto{}
	json.Unmarshal([]byte(response), &dto)

	if statusCode == 200 {
		formattedUserCreation := formatDate(dto.UserCreationDate)
		formattedLastUpload := formatDate(dto.LastUploadDate)
		localFileHash := fileHandler.Digest(filepath)
		isSync := localFileHash == dto.FileHash

		color.Blue("User creation date: " + formattedUserCreation)
		color.Blue("Last upload: " + formattedLastUpload)

		if isSync {
			color.Green("STATUS: SUCCESS")
			color.Green("File is synced with server. Everything is fine")
		} else {
			color.Yellow("STATUS: WARNING")
			color.Yellow("Not synced. This might indicate that your local file is older OR newer than the distant one")
			color.Yellow("Consider pulling OR pushing right now")
		}

		return
	}

	color.Red("STATUS: ERROR ->")
	color.Red(string(statusCode))
	fmt.Println(response)
}

func formatDate(date int64) string {
	unix := time.UnixMilli(date)
	return unix.Format("January 2, 2006")
}

func getSimpleHttpHeaders(fileHandler core.FileHandler, filepath string) core.BqtHttpHeaders {
	filename := core.FileHandler.GetFilename(fileHandler, filepath)

	userId := core.GetEnvironmentVariable(core.ENV_USER_ID)
	headers := core.BqtHttpHeaders{
		UserId:            userId,
		PreferredFilename: filename,
	}

	return headers
}
