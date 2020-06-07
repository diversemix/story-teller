package markovchain

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang-collections/go-datastructures/queue"
)

const MIN_WORDS = 5

type Builder struct{}

func (b *Builder) FromReader(r io.Reader) (*MarkovChain, error) {
	chain := new(MarkovChain)
	chain.Init()

	_, err := b.PopulateChain(chain, r)
	if err != nil {
		return nil, err
	}

	return chain, nil
}

func (b *Builder) PopulateChain(m *MarkovChain, r io.Reader) (int, error) {
	text := ""
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		text += string(buf[:n])
		if err == io.EOF {
			break
		}
	}
	return b.WordsToChain(m, text)
}

func (b *Builder) WordsToChain(m *MarkovChain, text string) (int, error) {
	words := strings.Fields(text)
	if len(words) < MIN_WORDS {
		return 0, fmt.Errorf("cannot process less than %q words", MIN_WORDS)
	}

	queue := new(queue.Queue)
	for _, word := range words {
		queue.Put(word)
		if queue.Len() == 3 {
			result, _ := queue.Get(3)
			m.Add(result[0].(string), result[1].(string), result[2].(string))
			queue.Put(result[1])
			queue.Put(result[2])
		}
	}

	return len(words), nil
}

func (b *Builder) GetStory(chain *MarkovChain, word1 string, word2 string, max int) (string, error) {
	story := word1 + " " + word2
	for i := 0; i < max; i++ {
		word, err := chain.SuggestNext(word1, word2)
		if err != nil {
			fmt.Errorf("Unable to build story %q", err)
			os.Exit(1)
		}
		story += " " + word
		word1 = word2
		word2 = word
	}
	return story, nil
}
