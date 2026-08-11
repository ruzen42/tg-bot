package markov

import (
	"time"
	"math/rand"
	"testing"
	"fmt"
)

func TestChain(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	chain := NewMChain()

	text := `в лесу родилась елочка в лесу она росла зимой и летом стройная зеленая была 
	метель ей пела песенку спи елочка бай бай мороз снежком укутывал смотри не замерзай`
	
	chain.Build(text)

	generatedText := chain.Generate("в", 500)

	fmt.Println(generatedText)
}
