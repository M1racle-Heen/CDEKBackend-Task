package cdekCostCalculator

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Client represents the client credentials and API URL needed for authentication.
type Client struct {
	Account string `json:"client_id"`
	Secret  string `json:"client_secret"`
	Token   string `json:"access_token"`
	URL     string `json:"api_url"`
}

// NewClient creates a new instance of the Client struct and returns a pointer to it and error while getting Auth Token.
func NewClient(account string, securePassword string, apiURL string) (*Client, error) {
	c := &Client{
		Account: account,
		Secret:  securePassword,
		URL:     apiURL,
	}

	// get Auth token
	err := c.getAuthToken()
	if err != nil {
		log.Fatalf("Got error while trying to get Auth Token %v.\n", err)
		return nil, err
	}

	return c, nil
}

// getAuthToken sends a POST request to the API URL to get an authentication token using the client_id and client_secret credentials.
func (c *Client) getAuthToken() error {
	URL := c.URL + "/oauth/token"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.Account)
	data.Set("client_secret", c.Secret)

	req, err := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Got error: %v\nWhile trying to make POST request to given url: %s.\n", err, URL)
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Got error: %v\nWhile trying to send POST request May be problem with: %v.\n", err, resp.Header)
		return err
	}

	//closes resp.Body while exited from getAuthToken method.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Got error: %v\nWhile trying to close client response.\n", err)
		}
	}(resp.Body)

	//Checks statusCode of response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get Auth token: returned non-OK status code %d.\n", resp.StatusCode)
		return nil
	}

	var tokenData struct {
		AccessToken string `json:"access_token"`
	}

	//Decodes Token for auth
	err = json.NewDecoder(resp.Body).Decode(&tokenData)

	if err != nil {
		log.Fatalf("Got error: %v\nWhile trying to decode Token for auth.\n", err)
		return err
	}

	c.Token = tokenData.AccessToken
	return nil
}
