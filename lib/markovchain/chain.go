package markovchain

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

type Occurrences map[string]int

type Chain map[string]Occurrences

type MarkovChain struct {
	chain Chain
}

func (m *MarkovChain) Init() {
	m.chain = make(Chain)
}

func (m *MarkovChain) Length() int {
	return len(m.chain)
}

func (m *MarkovChain) generateKey(p1 string, p2 string) string {
	return strings.ToLower(p1) + "_" + strings.ToLower(p2)
}

func (m *MarkovChain) Add(p1 string, p2 string, next string) {
	key := m.generateKey(p1, p2)
	if _, ok := m.chain[key]; !ok {
		m.chain[key] = make(Occurrences)
	}

	m.chain[key][strings.ToLower(next)]++
}

func (m *MarkovChain) Possibilities(p1 string, p2 string) Occurrences {
	key := p1 + "_" + p2
	return m.chain[key]
}

func (m *MarkovChain) String() (string, error) {
	val, err := json.Marshal(m.chain)
	return string(val), err
}

func (m *MarkovChain) SuggestNext(word1 string, word2 string) (string, error) {
	key := m.generateKey(word1, word2)
	nextOccurrences := m.chain[key]
	opts := make([]string, 0, len(nextOccurrences))
	for k := range nextOccurrences {
		opts = append(opts, k)
	}

	if len(opts) > 0 {
		next := rand.Intn(len(opts))
		//fmt.Printf("\n%s + %s = %q\n", word1, word2, opts[next])
		return opts[next], nil
	}

	return "", fmt.Errorf("Could not get next word for '%s %s'", word1, word2)
}
