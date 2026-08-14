package main

import (
	"hash/fnv"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	tg "gopkg.in/telebot.v4"

	log "tg-gen/logger"
	"tg-gen/markov"
)

// queryLength decides how many words to generate based on the inline
// query text: if it's a plain number, that number is used directly;
// otherwise the text is hashed into a number in [1, 100].
func queryLength(text string) int {
	text = strings.TrimSpace(text)

	n, err := strconv.Atoi(text)
	if err != nil {
		h := fnv.New32a()
		h.Write([]byte(text))
		n = int(h.Sum32()%100) + 1
	}

	switch {
	case n < 1:
		n = 1
	case n > 100:
		n = 100
	}
	return n
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var token = os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("token variable is not specified")
	}
	log.Debug("init")

	args := os.Args

	if len(args) < 2 {
		log.Fatal("*.gob chain is not specified")
	}

	chain, err := markov.LoadChain(args[1])

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("successfully loaded: " + args[1])

	pref := tg.Settings{
		Token:  token,
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err.Error())
	}

	b.Handle(tg.OnText, func(c tg.Context) error {
		// always reply when someone replies directly to the bot's message
		if c.Message().ReplyTo != nil && c.Message().ReplyTo.Sender.ID == b.Me.ID {
			generated := chain.Generate(queryLength(c.Text()))
			log.Debug("generated (reply): " + generated)
			return c.Send(generated)
		}

		if rand.Intn(3) == 0 {
			generated := chain.Generate(rand.Intn(10) + 5)
			log.Debug("generated: " + generated)
			return c.Send(generated)
		}
		return nil
	})

	b.Handle("/fetch", func(c tg.Context) error {
		out, err := fastfetch()
		if err != nil {
			log.Warn("fetch error: " + err.Error())
		}
		return c.Reply(out)
	})

	b.Handle(tg.OnChannelPost, func(c tg.Context) error {
		generated := chain.Generate(100)
		log.Info("generated: " + generated)
		return c.Send(generated)
	})

	b.Handle(tg.OnQuery, func(c tg.Context) error {
		query := c.Query()
		n := queryLength(query.Text)

		generated := chain.Generate(n)
		if generated == "" {
			log.Debug("empty generation for inline query, skipping answer")
			return c.Answer(&tg.QueryResponse{Results: tg.Results{}})
		}
		log.Info("generated inline: " + generated)

		result := &tg.ArticleResult{
			Title: generated,
			Text:  generated,
		}
		result.SetResultID(strconv.FormatInt(time.Now().UnixNano(), 10))

		return c.Answer(&tg.QueryResponse{
			Results:   tg.Results{result},
			CacheTime: 0,
		})
	})

	b.Start()
}
