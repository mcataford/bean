package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type CreateMessageRequestBody struct {
	Content string `json:"content"`
}

var API_BASEURL = "https://discord.com/api"

// Sends a message to the given channel.
// If an error occurs, the error is returned.
func sendMessage(message string, channel string) error {
	client := &http.Client{}

	body := &CreateMessageRequestBody{Content: message}

	bodyJson, _ := json.Marshal(body)

	req, requestError := http.NewRequest("POST", fmt.Sprintf("%s/channels/%s/messages", API_BASEURL, channel), bytes.NewBuffer(bodyJson))

	if requestError != nil {
		return requestError
	}

	appToken := os.Getenv("DISCORD_APP_TOKEN")

	if len(appToken) == 0 {
		log.Fatal("App token missing, must be provided as DISCORD_APP_TOKEN")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", appToken))

	_, responseError := client.Do(req)

	if responseError != nil {
		return responseError
	}

	return nil
}
