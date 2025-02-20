package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	fmt.Println("slack file upload")

	//
	if err := godotenv.Load(".env"); err != nil {

		log.Fatal("couldn't load the envs")
		return
	}

	API := slack.New(os.Getenv("SLACK_TOKEN"))
	channelsArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"paper.pdf"} //? if it doesnt work change the file or import a pdf file.

	for i := range fileArr {

		fileContent, err := os.Open(fileArr[i])
		if err != nil {
			log.Fatal("Error opening file: ", err)
			return
		}
		defer fileContent.Close()

		params := slack.UploadFileV2Parameters{
			Channel:  channelsArr[0],
			File:     fileArr[i],
			Reader:   fileContent,
			Filename: fileArr[i],
			FileSize: 1400,
		}

		file, err := API.UploadFileV2(params)

		fmt.Println(file)
		if err != nil {
			log.Fatal("there is an error ", err)
			return
		}

		fmt.Println("file name and url: ", file.Title, file.ID)
	}
}
