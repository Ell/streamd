package helix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type Client struct {
	clientId    string
	accessToken string
	baseUrl     string
}

// NewClient creates and returns a new twitch helix client
func NewClient(clientId, accessToken string) *Client {
	client := &Client{clientId, accessToken, "https://api.twitch.tv/helix"}

	return client
}

// NewTestClient creates and returns a new twitch helix client pointing at a different api host
func NewTestClient(clientId, accessToken, baseUrl string) *Client {
	client := &Client{clientId, accessToken, baseUrl}

	return client
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
		log.Printf("Could not create GET request %s", err)
		return respData, err
	}

	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Client-Id", client.clientId)

	res, err := c.Do(req)
	if err != nil {
		log.Printf("Could not make helix request %s", err)
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
		log.Printf("Could not read helix response body %s", err)
		return respData, err
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Printf("Unable to unmarshal helix data %s", err)
		return respData, err
	}

	return respData, nil
}

func makeHelixPostRequest[T, U any](client *Client, url string, data T) (APIResponse[U], error) {
	var respData APIResponse[U]

	c := http.Client{Timeout: 5 * time.Second}

	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not parse request body %s", err)
		return respData, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		log.Printf("Could not create POST request %s", err)
		return respData, err
	}

	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Client-Id", client.clientId)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		log.Printf("Could not make helix request %s", err)
		return respData, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		errorBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var errorData struct {
			Error   string `json:"error"`
			Status  int    `json:"status"`
			Message string `json:"message"`
		}

		err = json.Unmarshal(errorBody, &errorData)
		if err != nil {
			log.Fatal(err)
		}

		return respData, fmt.Errorf("helix request failed with status code %+v\n", errorData)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Unable to close helix body reader %s", err)
		}
	}(res.Body)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read helix response body %s", err)
		return respData, err
	}

	if len(respBody) <= 0 {
		return respData, nil
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Printf("Unable to unmarshal helix data %s", err)
		return respData, err
	}

	return respData, nil
}

func makeHelixDeleteRequest[T, U interface{}](client *Client, url string, data T) (APIResponse[U], error) {
	var respData APIResponse[U]

	c := http.Client{Timeout: 5 * time.Second}

	v, err := query.Values(data)
	if err != nil {
		log.Printf("Error creating query string for helix request %s", err)
	}

	encodedUrl := url + "?" + v.Encode()

	req, err := http.NewRequest("DELETE", encodedUrl, nil)
	if err != nil {
		fmt.Printf("Could not create DELETE request %s", err)
		return respData, err
	}

	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Client-Id", client.clientId)

	res, err := c.Do(req)
	if err != nil {
		log.Printf("Could not make helix request %s", err)
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
		log.Printf("Could not read helix response body %s", err)
		return respData, err
	}

	if len(respBody) <= 0 {
		return respData, nil
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Printf("Unable to unmarshal helix data %s", err)
		return respData, err
	}

	return respData, nil
}

func makeHelixPutRequest[T, U interface{}](client *Client, url string, data T) (APIResponse[U], error) {
	var respData APIResponse[U]

	c := http.Client{Timeout: 5 * time.Second}

	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not parse request body %s", err)
		return respData, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		log.Printf("Could not create POST request %s", err)
		return respData, err
	}

	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Client-Id", client.clientId)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		log.Printf("Could not make helix request %s", err)
		return respData, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		errorBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var errorData struct {
			Error   string `json:"error"`
			Status  int    `json:"status"`
			Message string `json:"message"`
		}

		err = json.Unmarshal(errorBody, &errorData)
		if err != nil {
			log.Fatal(err)
		}

		return respData, fmt.Errorf("helix request failed with status code %+v\n", errorData)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Unable to close helix body reader %s", err)
		}
	}(res.Body)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read helix response body %s", err)
		return respData, err
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Printf("Unable to unmarshal helix data %s", err)
		return respData, err
	}

	return respData, nil
}

func makeHelixPatchRequest[T, U interface{}](client *Client, url string, data T) (APIResponse[U], error) {
	var respData APIResponse[U]

	c := http.Client{Timeout: 5 * time.Second}

	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not parse request body %s", err)
		return respData, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewReader(body))
	if err != nil {
		log.Printf("Could not create POST request %s", err)
		return respData, err
	}

	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Client-Id", client.clientId)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		log.Printf("Could not make helix request %s", err)
		return respData, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		errorBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var errorData struct {
			Error   string `json:"error"`
			Status  int    `json:"status"`
			Message string `json:"message"`
		}

		err = json.Unmarshal(errorBody, &errorData)
		if err != nil {
			log.Fatal(err)
		}

		return respData, fmt.Errorf("helix request failed with status code %+v\n", errorData)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Unable to close helix body reader %s", err)
		}
	}(res.Body)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read helix response body %s", err)
		return respData, err
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		log.Printf("Unable to unmarshal helix data %s", err)
		return respData, err
	}

	return respData, nil
}
