package webhook

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type UserData struct {
	Avatar           string `json:"avatar"`
	AvatarDecoration string `json:"avatar_decoration"`
	Discriminator    string `json:"discriminator"`
	DisplayName      string `json:"display_name"`
	GlobalName       string `json:"global_name"`
	Id               string `json:"id"`
	PublicFlags      int    `json:"public_flags"`
	Username         string `json:"username"`
}

type InteractionData struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

type WebhookRequestBody struct {
	ApplicationId string          `json:"application_id"`
	Entitlements  []string        `json:"entitlements"`
	Id            string          `json:"id"`
	Token         string          `json:"token"`
	Type          int             `json:"type"`
	User          UserData        `json:"user"`
	Version       int             `json:"version"`
	Data          InteractionData `json:"data"`
	ChannelId     string          `json:"channel_id"`
}

type InteractionResponseData struct {
	Content string `json:"content"`
}

type InteractionResponse struct {
	Type int                     `json:"type"`
	Data InteractionResponseData `json:"data"`
}

func Respond(responseWriter http.ResponseWriter, request *http.Request, message []byte, statusCode int) {
	requestUrl := request.URL

	log.Println(fmt.Sprintf("%s - %d", requestUrl, statusCode))
	if statusCode != 200 {
		responseWriter.WriteHeader(statusCode)
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(message)
}

func verifyInteractionSignature(publicKey string, message string, signature string) bool {
	decodedKey, keyErr := hex.DecodeString(publicKey)

	if keyErr != nil {
		log.Fatal(keyErr)
	}

	decodedSignature, sigErr := hex.DecodeString(signature)

	if sigErr != nil {
		log.Fatal(sigErr)
	}

	return ed25519.Verify(decodedKey, []byte(message), decodedSignature)
}

// Handles verification interactions and pings from Discord.
func handlePing(response http.ResponseWriter, request *http.Request, body string) {
	responseBody := map[string]int{"type": 1}
	responseJson, _ := json.Marshal(responseBody)
	signature := request.Header["X-Signature-Ed25519"][0]
	timestamp := request.Header["X-Signature-Timestamp"][0]

	message := timestamp + string(body)

	publicKey := os.Getenv("DISCORD_APP_PUBLIC_KEY")

	if len(publicKey) == 0 {
		log.Fatal("No public key supplied, must be available as DISCORD_APP_PUBLIC_KEY")
	}

	if !verifyInteractionSignature(publicKey, message, signature) {
		Respond(response, request, []byte{}, 401)
	}

	Respond(response, request, responseJson, 200)
}

func handleRUOk(response http.ResponseWriter, request *http.Request) {
	responseBody := InteractionResponse{
		Type: 4,
		Data: InteractionResponseData{Content: "Meow"},
	}

	jsonJson, _ := json.Marshal(responseBody)

	Respond(response, request, jsonJson, 200)
}

func handleInteraction(response http.ResponseWriter, request *http.Request, interactionData InteractionData) {
	log.Println(fmt.Sprintf("Handling interaction: %s", interactionData.Name))
	switch interactionData.Name {
	case "ruok":
		handleRUOk(response, request)
	}
}

func handleWebhook(response http.ResponseWriter, request *http.Request) {
	// FIXME: Should do something about consuming the body. :(
	body, _ := ioutil.ReadAll(request.Body)

	requestBody := WebhookRequestBody{}
	jsonError := json.Unmarshal(body, &requestBody)

	if jsonError != nil {
		log.Fatal(jsonError)
	}

	switch requestBody.Type {
	case 1:
		handlePing(response, request, string(body))
	case 2:
		handleInteraction(response, request, requestBody.Data)
	default:
		Respond(response, request, []byte{}, 400)
	}
}

func ListenToWebhook(host string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleWebhook)

	return http.ListenAndServe(host, mux)
}
