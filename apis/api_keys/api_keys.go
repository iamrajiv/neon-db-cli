package api_keys

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// ListAPIKeys retrieves the API keys for your Neon account.
// The response does not include API key tokens. A token is only provided when creating an API key.
// API keys can also be managed in the Neon Console.
// For more information, see Manage API keys: https://neon.tech/docs/manage/api-keys/.
func ListAPIKeys() ([]byte, error) {
	url := "https://console.neon.tech/api/v2/api_keys"
	apiKey := os.Getenv("API_KEY")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// CreateAPIKey creates an API key.
// The `keyName` is a user-specified name for the key.
// This method returns an `id` and `key`. The `key` is a randomly generated, 64-bit token required to access the Neon API.
// API keys can also be managed in the Neon Console.
// See [Manage API keys](https://neon.tech/docs/manage/api-keys/).
func CreateAPIKey(keyName string) ([]byte, error) {
	url := "https://console.neon.tech/api/v2/api_keys"
	apiKey := os.Getenv("API_KEY")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+apiKey)

	req.Body = ioutil.NopCloser(strings.NewReader(fmt.Sprintf(`{"key_name":"%s"}`, keyName)))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Revokes the specified API key.
// An API key that is no longer needed can be revoked.
// This action cannot be reversed.
// You can obtain `key_id` values by listing the API keys for your Neon account.
// API keys can also be managed in the Neon Console.
// See [Manage API keys](https://neon.tech/docs/manage/api-keys/).
func RevokeAPIKey(keyID string) ([]byte, error) {
	url := fmt.Sprintf("https://console.neon.tech/api/v2/api_keys/%s", keyID)
	apiKey := os.Getenv("API_KEY")
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
