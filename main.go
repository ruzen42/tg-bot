package main

import (
	log "tg-gen/logger"
	"time"
	"os"
	tg "gopkg.in/telebot.v4"
)

func main() {
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
