package markov

import (
	"math/rand"
	"testing"
	"time"
	"os"
	"fmt"
//	"encoding/gob"
)

func TestChain(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	chain := NewMChain(2)
	data, err := os.ReadFile("export.txt")
	
	if err != nil {
		return
	}

	chain.Build(string(data))

	fmt.Println(chain.Generate(10))
	
	/*
	file, err := os.Create("chain.gob")

	if err != nil {
		return
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	if err := enc.Encode(chain); err != nil {
		return
	}*/
}
