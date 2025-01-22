package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	cloudflareAPIURL = "https://api.cloudflare.com/client/v4/accounts/ba3c5007d3b4ab6bf0afce10debfd04a/calls/apps/522a4982829211e2b4930b5788ad526d"
	cloudflareToken  = "XgT5WXqSQCGNlCmj-Fw4YNMqapONysnul900UWns"
)

func CreateCloudflareRoom(roomName string) (string, error) {
	log.Println("CreateCloudflareRoom: Creating room with name:", roomName)

	requestBody, _ := json.Marshal(map[string]string{
		"name": roomName,
	})

	log.Println("CreateCloudflareRoom: Request body:", string(requestBody))

	req, err := http.NewRequest("POST", cloudflareAPIURL+"/rooms", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("CreateCloudflareRoom: Failed to create request", err)
		return "", err
	}

	// Sử dụng API Token
	req.Header.Set("Authorization", "Bearer "+cloudflareToken)
	req.Header.Set("Content-Type", "application/json")

	log.Println("CreateCloudflareRoom: Sending request to Cloudflare API")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("CreateCloudflareRoom: Failed to send request", err)
		return "", err
	}
	defer resp.Body.Close()

	log.Println("CreateCloudflareRoom: Cloudflare API response status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		log.Println("CreateCloudflareRoom: Cloudflare API error response:", errorResponse)
		return "", errors.New("failed to create room")
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("CreateCloudflareRoom: Cloudflare API success response:", result)

	return result["roomID"].(string), nil
}

func JoinCloudflareRoom(roomID, userName string) (string, error) {
	log.Println("JoinCloudflareRoom: Joining room with ID:", roomID, "and user:", userName)

	requestBody, _ := json.Marshal(map[string]string{
		"userName": userName,
	})

	log.Println("JoinCloudflareRoom: Request body:", string(requestBody))

	req, err := http.NewRequest("POST", cloudflareAPIURL+"/rooms/"+roomID+"/join", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("JoinCloudflareRoom: Failed to create request", err)
		return "", err
	}

	// Sử dụng API Token
	req.Header.Set("Authorization", "Bearer "+cloudflareToken)
	req.Header.Set("Content-Type", "application/json")

	log.Println("JoinCloudflareRoom: Sending request to Cloudflare API")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("JoinCloudflareRoom: Failed to send request", err)
		return "", err
	}
	defer resp.Body.Close()

	log.Println("JoinCloudflareRoom: Cloudflare API response status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		log.Println("JoinCloudflareRoom: Cloudflare API error response:", errorResponse)
		return "", errors.New("failed to join room")
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println("JoinCloudflareRoom: Cloudflare API success response:", result)

	return result["token"].(string), nil
}
