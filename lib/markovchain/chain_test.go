package markovchain

import "testing"

func TestInit(t *testing.T) {
	want := 0
	m := new(MarkovChain)
	m.Init()
	if got := m.Length(); got != want {
		t.Errorf("Init() results in length = %q, want %q", got, want)
	}
}

func TestInitAfterAdd(t *testing.T) {
	want := 0
	m := new(MarkovChain)
	m.Init()
	m.Add("a", "b", "c")
	m.Init()
	if got := m.Length(); got != want {
		t.Errorf("Init() results in length = %q, want %q", got, want)
	}
}

func TestAddChangesLength(t *testing.T) {
	want := 1
	m := new(MarkovChain)
	m.Init()
	m.Add("a", "b", "c")
	if got := m.Length(); got != want {
		t.Errorf("Add() results in length = %q, want %q", got, want)
	}
	want = 2
	m.Add("a", "c", "c")
	if got := m.Length(); got != want {
		t.Errorf("Add() results in length = %q, want %q", got, want)
	}
}

func TestAddChangesPossibilities(t *testing.T) {
	want := 0
	m := new(MarkovChain)
	m.Init()
	if got := m.Possibilities("a", "b")["c"]; got != want {
		t.Errorf("Possibilities() results in length = %q, want %q", got, want)
	}

	want = 1
	m.Add("a", "b", "c")
	if got := m.Possibilities("a", "b")["c"]; got != want {
		t.Errorf("Possibilities() results in length = %q, want %q", got, want)
	}

	want = 2
	m.Add("a", "b", "c")
	if got := m.Possibilities("a", "b")["c"]; got != want {
		t.Errorf("Possibilities() results in length = %q, want %q", got, want)
	}
}

func TestAChain(t *testing.T) {
	want := "{\"a_a\":{\"b\":2},\"a_b\":{\"c\":1,\"d\":1}}"
	m := new(MarkovChain)
	m.Init()
	m.Add("a", "B", "c")
	m.Add("a", "b", "d")
	m.Add("A", "a", "b")
	m.Add("a", "a", "B")
	if got, err := m.String(); got != want {
		t.Errorf("String() results = %q, want %q error %q", got, want, err)
	}
}
