package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Could not load .env file.")
		os.Exit(1)
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(botToken)

	channelId := os.Getenv("CHANNEL_ID")
	channelArr := []string{channelId}

	fileArr := []string{"test.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
