package core

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	HTTP_TIMEOUT_SECONDS      = 10
	HEADER_USER_ID            = "User-Id"
	HEADER_PREFERRED_FILENAME = "Preferred-Filename"
	HEADER_POSITION           = "Position"
	HEADER_CONTENT_TYPE       = "Content-Type"
)

type HttpClient interface {
	Get(url string, headers BqtHttpHeaders) (string, int32)
	Multipart(url string, path string, headers BqtHttpHeaders) (string, int32)
}

type GoHttpClient struct {
}

type BqtHttpHeaders struct {
	UserId            string
	PreferredFilename string
	Position          int
}

func (GoHttpClient) Get(url string, headers BqtHttpHeaders) (string, int32) {
	request, _ := http.NewRequest("GET", url, nil)
	setHttpHeaders(request, headers)
	fmt.Println(request.Header)
	client := &http.Client{
		Timeout: HTTP_TIMEOUT_SECONDS * time.Second,
	}

	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	return string(body), int32(response.StatusCode)
}

func (GoHttpClient) Multipart(url string, path string, headers BqtHttpHeaders) (string, int32) {
	file, _ := os.Open(path)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("a", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	request, _ := http.NewRequest("POST", url, body)

	setHttpHeaders(request, headers)
	request.Header.Add(HEADER_CONTENT_TYPE, writer.FormDataContentType())

	client := &http.Client{
		Timeout: HTTP_TIMEOUT_SECONDS * time.Second,
	}

	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	return string(responseBody), int32(response.StatusCode)
}

func setHttpHeaders(request *http.Request, headers BqtHttpHeaders) {
	request.Header.Set(HEADER_USER_ID, headers.UserId)
	request.Header.Set(HEADER_PREFERRED_FILENAME, headers.PreferredFilename)
	request.Header.Set(HEADER_POSITION, strconv.Itoa(headers.Position))
}
