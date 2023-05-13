package main

import (
	commands "beanbot/internal/commands"
	webhook "beanbot/internal/webhook"
	"log"
    "fmt"
    "strings"
	"os"
)

func validateEnvironment() {
    expectedEnv := []string{
        "DISCORD_APP_TOKEN",
        "DISCORD_APP_ID",
        "DISCORD_APP_PUBLIC_KEY",
    }

    missingEnv := []string{}

    for _, expectedEnvKey := range expectedEnv {
        _, exists := os.LookupEnv(expectedEnvKey)

        if !exists {
            missingEnv = append(missingEnv, expectedEnvKey)
        }
    }

    if len(missingEnv) != 0 {
        log.Fatal(fmt.Sprintf("Missing some required environment variables: %s", strings.Join(missingEnv, ", ")))
    }
}

func main() {
	validateEnvironment()

    appToken := os.Getenv("DISCORD_APP_TOKEN")
	appId := os.Getenv("DISCORD_APP_ID")

	commands.InstallCommands(appToken, appId)
	log.Println("😸 Listening to webhook requests...")
	err := webhook.ListenToWebhook(":8080")

	if err != nil {
		log.Fatal(err)
	}
}
