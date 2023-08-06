package core

import "os"

const (
	ENV_FILEPATH = "BOUQUETIN_FILEPATH"
	ENV_USER_ID  = "BOUQUETIN_ID"
)

func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}
