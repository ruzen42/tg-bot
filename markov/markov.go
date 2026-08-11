package markov

import (
	"encoding/gob"
	"math/rand"
	"os"
	"strings"
)

type MChain struct {
	PrefixLen   int                 // N len
	Transitions map[string][]string // map Transitions
	Prefixes    []string            // all keys
}

func (c *MChain) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(c)
}

func LoadChain(filename string) (*MChain, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var chain MChain
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&chain)
	if err != nil {
		return nil, err
	}

	return &chain, nil
}

func NewMChain(PrefixLen int) *MChain {
	return &MChain{
		PrefixLen:   PrefixLen,
		Transitions: make(map[string][]string),
		Prefixes:    make([]string, 0),
	}
}

func buildKey(words []string) string {
	return strings.Join(words, " ")
}

func (c *MChain) Build(text string) {
	words := strings.Fields(text)

	if len(words) <= c.PrefixLen {
		return
	}

	for i := 0; i < len(words)-c.PrefixLen; i++ {
		prefixSlice := words[i : i+c.PrefixLen]
		prefixKey := buildKey(prefixSlice)

		nextWord := words[i+c.PrefixLen]

		if _, exists := c.Transitions[prefixKey]; !exists {
			c.Prefixes = append(c.Prefixes, prefixKey)
		}

		c.Transitions[prefixKey] = append(c.Transitions[prefixKey], nextWord)
	}
}

func (c *MChain) Generate(n int) string {
	if len(c.Prefixes) == 0 {
		return ""
	}

	startPrefix := c.Prefixes[rand.Intn(len(c.Prefixes))]

	currentWords := strings.Split(startPrefix, " ")

	result := append([]string{}, currentWords...)

	for i := 0; i < n; i++ {
		currentKey := buildKey(currentWords)

		possibleNextWords, exists := c.Transitions[currentKey]
		if !exists || len(possibleNextWords) == 0 {
			break
		}

		nextWord := possibleNextWords[rand.Intn(len(possibleNextWords))]
		result = append(result, nextWord)

		currentWords = append(currentWords[1:], nextWord)
	}

	return strings.Join(result, " ")
}
