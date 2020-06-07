package main

import (
	"fmt"
	"markovchain"
	"os"
)

const MAX_WORDS = 200
const FILENAME = "story.txt"
const WORD1 = "then"
const WORD2 = "Alice"

func main() {
	fmt.Println("MarkovChain Story Teller, Peter Hooper, 2020 MIT Licence")
	builder := new(markovchain.Builder)

	reader, err := os.Open(FILENAME)

	if err != nil {
		fmt.Errorf("Could not read %q - %q\n", FILENAME, err)
		os.Exit(1)
	}

	chain, err := builder.FromReader(reader)
	if err != nil {
		fmt.Errorf("Unable to build chain %q", err)
		os.Exit(1)
	}
	reader.Close()

	fmt.Printf("Formed chain with length %d\n", chain.Length())

	story, err := builder.GetStory(chain, WORD1, WORD2, MAX_WORDS)
	if err != nil {
		fmt.Errorf("Unable to build story %q", err)
		os.Exit(1)
	}

	fmt.Println("\n\n ~~~ Are you ready for the story? ~~~\n")
	fmt.Printf("%s\n\n ~~~ Thats all folks! ~~~ \n\n", story)

}
