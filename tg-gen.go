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

	args := os.Args

	if len(args) < 2 {
		log.Fatal("*.gob chain is not specified")
	}

	pref := tg.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err.Error())
	}

	chain, err := markov.LoadChain(args[1])

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("successfully loaded: " + args[1])

	var generated string
	b.Handle(tg.OnText, func(c tg.Context) error {
		if rand.Intn(5) == 0 {
			generated = chain.Generate(rand.Intn(100) + 5)
			log.Debug("generated: " + generated)
			return c.Send(generated)
		}
		return nil
	})

	b.Handle(tg.OnChannelPost, func(c tg.Context) error {
		generated = chain.Generate(1000)
		log.Info("generated: " + generated)
		return c.Send(generated)
	})

	b.Start()
}
