package core

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	HTTP_TIMEOUT_SECONDS      = 10
	HEADER_USER_ID            = "User-Id"
	HEADER_PREFERRED_FILENAME = "Preferred-Filename"
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
}

// func (client *DefaultHTTPClient) Post(url string, bodyType string, body interface{}) (*http.Response, error) {
// 	// Ici, vous pouvez ajouter votre propre logique pour les requêtes POST.
// 	// Par exemple, sérialiser le corps en JSON, XML, etc., et effectuer la requête.
// 	return nil, fmt.Errorf("POST method not implemented")
// }

// // Example d'utilisation
// func main() {
// 	// Créez un client HTTP personnalisé en utilisant MyHTTPClient.
// 	client := &DefaultHTTPClient{
// 		Client: &http.Client{
// 			Timeout: 10 * time.Second,
// 		},
// 	}

// 	// Utilisez le client pour effectuer une requête GET.
// 	response, err := client.Get("https://api.example.com/data")
// 	if err != nil {
// 		fmt.Println("Erreur lors de la requête GET:", err)
// 		return
// 	}
// 	defer response.Body.Close()
// }
