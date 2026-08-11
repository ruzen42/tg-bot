package markov

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"strings"
)

func TestChain(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	chain := NewMChain(2)

	/*
		content, err := os.ReadFile("chat_logs.txt")
		if err != nil {
			panic(err)
		}
		text := string(content)
	*/

	text := `привет, как дела?
	нормально, сижу пишу код на голанге.
	а я пошел пить кофе.
	завтра во сколько собираемся?
	думаю часам к восьми.
	сижу пишу код, ничего не получается!`

	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, "!", "")

	chain.Build(text)
	fmt.Printf("prefixes: %d\n\n", len(chain.prefixes))

	for i := 1; i <= 5; i++ {
		fmt.Printf("msg %d: %s\n", i, chain.Generate(15))
	}
}
