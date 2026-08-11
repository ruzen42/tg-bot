package markov

import (
	"math/rand"
	"strings"
)

type MChain struct {
	prefixLen   int                 // N len
	transitions map[string][]string // map transitions
	prefixes    []string            // all keys
}

func NewMChain(prefixLen int) *MChain {
	return &MChain{
		prefixLen:   prefixLen,
		transitions: make(map[string][]string),
		prefixes:    make([]string, 0),
	}
}

func buildKey(words []string) string {
	return strings.Join(words, " ")
}

func (c *MChain) Build(text string) {
	words := strings.Fields(text)

	if len(words) <= c.prefixLen {
		return
	}

	for i := 0; i < len(words)-c.prefixLen; i++ {
		prefixSlice := words[i : i+c.prefixLen]
		prefixKey := buildKey(prefixSlice)

		nextWord := words[i+c.prefixLen]

		if _, exists := c.transitions[prefixKey]; !exists {
			c.prefixes = append(c.prefixes, prefixKey)
		}

		c.transitions[prefixKey] = append(c.transitions[prefixKey], nextWord)
	}
}

func (c *MChain) Generate(n int) string {
	if len(c.prefixes) == 0 {
		return ""
	}

	startPrefix := c.prefixes[rand.Intn(len(c.prefixes))]

	currentWords := strings.Split(startPrefix, " ")

	result := append([]string{}, currentWords...)

	for i := 0; i < n; i++ {
		currentKey := buildKey(currentWords)

		possibleNextWords, exists := c.transitions[currentKey]
		if !exists || len(possibleNextWords) == 0 {
			break
		}

		nextWord := possibleNextWords[rand.Intn(len(possibleNextWords))]
		result = append(result, nextWord)

		currentWords = append(currentWords[1:], nextWord)
	}

	return strings.Join(result, " ")
}
