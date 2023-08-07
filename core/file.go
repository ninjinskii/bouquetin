package core

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	files "path/filepath"
)

type FileHandler interface {
	WriteFile(filepath string, content string)
	CopyFile(sourcePath string, destinationPath string)
	GetFilename(filepath string) string
	Digest(filepath string) string
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

func (GoFileHandler) GetFilename(filepath string) string {
	return files.Base(filepath)
}

func (GoFileHandler) Digest(filepath string) string {
	fileBuffer, error := os.ReadFile(filepath)

	if error != nil {
		panic(error)
	}

	hash := sha1.Sum(fileBuffer)
	hashArray := hash[:]
	return hex.EncodeToString(hashArray)
}
