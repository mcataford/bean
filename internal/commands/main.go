package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

var Commands = []Command{
	Command{Name: "ruok", Description: "Bean? Are you okay?", Type: 1},
}

// FIXME: Consider moving to a more central location.
var API_BASEURL = "https://discord.com/api"

// Installs global commands as defined under Commands.
func InstallCommands(botToken string, appId string) {
	endpointUrl := fmt.Sprintf("%s/applications/%s/commands", API_BASEURL, appId)

	client := &http.Client{}

	payload, jsonErr := json.Marshal(Commands)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	req, requestErr := http.NewRequest("PUT", endpointUrl, bytes.NewBuffer(payload))

	if requestErr != nil {
		log.Fatal(requestErr)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", botToken))

	_, responseErr := client.Do(req)

	if responseErr != nil {
		log.Println("😿 Failed to install commands.")
		log.Fatal(responseErr)
	}

	log.Println(fmt.Sprintf("😸 Installed %d commands.", len(Commands)))

	for _, command := range Commands {
		log.Println(fmt.Sprintf(">> /%s", command.Name))
	}
}
