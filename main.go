package main

import (
	tg "gopkg.in/telebot.v4"
	"os"
	log "tg-gen/logger"
	"time"
)

func main() {
	var token = os.Getenv("TOKEN")
	if token == nil {
		log.Fatal("token variable is not specified")
	}

	pref := tg.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	b.Handle("/ping", func(c tg.Context) error {
		return c.Send("pong")
	})

	b.Start()
}
