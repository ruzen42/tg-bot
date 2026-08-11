package markov

import (
	"math/rand"
	"strings"
)

type MChain struct {
	transitions map[string][]string
}

func NewMChain() *MChain {
	return &MChain{
		transitions: make(map[string][]string),
	}
}

func (c *MChain) Build(text string) {
	words := strings.Fields(text)

	if len(words) == 0 {
		return
	}

	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]
		c.transitions[currentWord] = append(c.transitions[currentWord], nextWord)
	}
}

func (c *MChain) Generate(seedWord string, n int) string {
	var result []string
	
	currentWord := seedWord
	result = append(result, currentWord)
	
	for i := 1; i < n; i++ {
		possibleNextWords, exists := c.transitions[currentWord]

		// if dict not contains word, stop
		if !exists || len(possibleNextWords) == 0 {
			break
		}

		randomIndex := rand.Intn(len(possibleNextWords))
		currentWord = possibleNextWords[randomIndex]

		result = append(result, currentWord)
	}
	
	return strings.Join(result, " ")
}


