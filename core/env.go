package core

import "os"

func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}
