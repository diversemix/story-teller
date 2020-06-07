package markovchain

import (
	"strings"
	"testing"
)

func TestFromReader(t *testing.T) {
	want := 7
	builder := new(Builder)
	reader := strings.NewReader("once upon a time there were three little pigs")

	m, err := builder.FromReader(reader)
	if err != nil {
		t.Errorf("Create() returned: %q", err)
	}

	if got := m.Length(); got != want {
		t.Errorf("Length() results in length = %d, want %d", got, want)
	}

	wantString := "{\"a_time\":{\"there\":1},\"once_upon\":{\"a\":1},\"there_were\":{\"three\":1},\"three_little\":{\"pigs\":1},\"time_there\":{\"were\":1},\"upon_a\":{\"time\":1},\"were_three\":{\"little\":1}}"
	if got, _ := m.String(); got != wantString {
		t.Errorf("String() results in = %q, want %q", got, wantString)
	}
}
