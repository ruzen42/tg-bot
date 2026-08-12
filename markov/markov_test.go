package markov

import (
	"testing"
	"os"
	"fmt"
)

func TestChain(t *testing.T) {
	chain := NewMChain(4)
	data, err := os.ReadFile("export.txt")
	
	if err != nil {
		return
	}

	chain.Build(string(data))

	fmt.Println(chain.Generate(10))
	
	chain.Save("../bin/slackware.gob")
}
