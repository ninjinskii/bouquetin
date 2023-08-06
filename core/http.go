package core

import (
	"io"
	"net/http"
	"time"
)

const (
	HEADER_USER_ID            = "User-Id"
	HEADER_PREFERRED_FILENAME = "Preferred-Filename"
)

type HttpClient interface {
	Get(url string, headers BqtHttpHeaders) string
}

type DefaultHttpClient struct {
	Client *http.Client
}

type BqtHttpHeaders struct {
	UserId            string
	PreferredFilename string
}

func NewHttpClient() HttpClient {
	return &DefaultHttpClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (client *DefaultHttpClient) Get(url string, headers BqtHttpHeaders) string {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set(HEADER_USER_ID, headers.UserId)
	request.Header.Set(HEADER_PREFERRED_FILENAME, headers.PreferredFilename)
	response, error := client.Client.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	return string(body)
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
