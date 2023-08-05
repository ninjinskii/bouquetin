package core

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, bodyType string, body interface{}) (*http.Response, error)
}

type DefaultHTTPClient struct {
	Client *http.Client
}

func (client *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return client.Client.Get(url)
}

// func (client *DefaultHTTPClient) Post(url string, bodyType string, body interface{}) (*http.Response, error) {
// 	// Ici, vous pouvez ajouter votre propre logique pour les requêtes POST.
// 	// Par exemple, sérialiser le corps en JSON, XML, etc., et effectuer la requête.
// 	return nil, fmt.Errorf("POST method not implemented")
// }

// Example d'utilisation
func main() {
	// Créez un client HTTP personnalisé en utilisant MyHTTPClient.
	client := &DefaultHTTPClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// Utilisez le client pour effectuer une requête GET.
	response, err := client.Get("https://api.example.com/data")
	if err != nil {
		fmt.Println("Erreur lors de la requête GET:", err)
		return
	}
	defer response.Body.Close()
}
