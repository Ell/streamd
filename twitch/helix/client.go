package helix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	clientId    string
	accessToken string
	baseUrl     string
}

// NewClient creates and returns a new twitch helix client
func NewClient(clientId, accessToken string) (*Client, error) {
	client := &Client{clientId, accessToken, "https://api.twitch.tv/helix"}

	return client, nil
}

// NewTestClient creates and returns a new twitch helix client pointing at a different api host
func NewTestClient(clientId, accessToken, baseUrl string) (*Client, error) {
	client := &Client{clientId, accessToken, baseUrl}

	return client, nil
}

func makeHelixGetRequest[T, U interface{}](client *Client, url string, data T) (APIResponse[U], error) {
	var respData APIResponse[U]

	c := http.Client{Timeout: 5 * time.Second}

	v, err := query.Values(data)
	if err != nil {
		fmt.Printf("Error creating query string for helix request %s", err)
	}

	url = url + "?" + v.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Could not create GET request %s", err)
		return respData, err
	}

	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Client-Id", client.clientId)

	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("Could not make helix request %s", err)
		return respData, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Unable to close helix body reader %s", err)
		}
	}(res.Body)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read helix response body %s", err)
		return respData, err
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		fmt.Printf("Unable to unmarshal helix data %s", err)
		return respData, err
	}

	return respData, nil
}

func makeHelixPostRequest[T, U interface{}](client *Client, url string, data T) (APIResponse[U], error) {
	var respData APIResponse[U]

	c := http.Client{Timeout: 5 * time.Second}

	body, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Could not parse request body %s", err)
		return respData, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Printf("Could not create POST request %s", err)
		return respData, err
	}

	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Client-Id", client.clientId)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("Could not make helix request %s", err)
		return respData, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Unable to close helix body reader %s", err)
		}
	}(res.Body)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read helix response body %s", err)
		return respData, err
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		fmt.Printf("Unable to unmarshal helix data %s", err)
		return respData, err
	}

	return respData, nil
}
