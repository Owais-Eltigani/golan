package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	fmt.Println("slack_bot ... ")

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("couldn't load file")
	}
	fmt.Println(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// setting the .env
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	fmt.Println(bot)

	// goroutine.
	go printCommandEvent(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Examples:    []string{"my yob is 2000"},
		Handler: func(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {

			year := req.Param("year")
			yr, err := strconv.Atoi(year)

			if err != nil {
				fmt.Println(err)
			}

			age := 2025 - yr
			fmt.Println("age is: ", age)

			res.Reply(strconv.Itoa(age))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {

		log.Fatal(err)
	}

}

func printCommandEvent(analyticalData <-chan *slacker.CommandEvent) {

	// displaying event info
	for event := range analyticalData {

		fmt.Println("event information: ")
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
	}
}
