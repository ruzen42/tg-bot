package main

import (
	tg "gopkg.in/telebot.v4"
	"math/rand"
	"os"
	log "tg-gen/logger"
	"tg-gen/markov"
	"time"
)

func main() {
	var token = os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("token variable is not specified")
	}

	pref := tg.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err.Error())
	}

	chain, err := markov.LoadChain("bin/chain.gob")

	if err != nil {
		log.Fatal(err.Error())
	}

	b.Handle(tg.OnText, func(c tg.Context) error {
		return c.Send(chain.Generate(rand.Intn(10) + 5))
	})

	b.Start()
}
