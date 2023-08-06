package core

import (
	"io"
	"os"
)

type FileHandler interface {
	WriteFile(filepath string, content string)
	CopyFile(sourcePath string, destinationPath string)
}

type GoFileHandler struct {
}

func (GoFileHandler) WriteFile(filepath string, content string) {
	file, error := os.Create(filepath)

	if error != nil {
		panic(error)
	}

	defer file.Close()

	_, error = file.WriteString(content)

	if error != nil {
		panic(error)
	}
}

func (GoFileHandler) CopyFile(sourcePath string, destinationPath string) {
	sourceFile, error := os.Open(sourcePath)

	if error != nil {
		panic(error)
	}

	defer sourceFile.Close()

	destinationFile, error := os.Create(destinationPath)

	if error != nil {
		panic(error)
	}

	defer destinationFile.Close()

	_, error = io.Copy(destinationFile, sourceFile)

	if error != nil {
		panic(error)
	}

	error = destinationFile.Sync()

	if error != nil {
		panic(error)
	}
}
