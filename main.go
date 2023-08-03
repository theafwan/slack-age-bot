package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOCKEN", "xoxb-4668272650624-5656435173346-3w2Bkyigj68eZ7Zzc6wPBX6y")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05KAB15J58-5680081784368-6a45dfca66ee5d1c2c2a50c07b1fd13e8d2aac12fea5a99bcec72a4cbc2496cc")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOCKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println(err)
			}
			age := 2023 - yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
