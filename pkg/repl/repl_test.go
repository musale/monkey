package repl

import (
	"strings"
	"sync"
	"testing"
)

type writerSpy struct {
	mutex   sync.Mutex
	counter int
}

func (w *writerSpy) Write(p []byte) (n int, err error) {
	w.mutex.Lock()
	w.counter++
	w.mutex.Unlock()
	return len(p), nil
}

func TestStart(t *testing.T) {
	t.Run("Count exactly 2 calls for PROMPT and monkeyFace", func(t *testing.T) {
		input := strings.NewReader("")
		out := &writerSpy{}
		calls := 2
		Start(input, out)
		compareInts(t, out.counter, calls)
	})

	t.Run("Count exactly 3 calls for PROMPT, monkeyFace and exit", func(t *testing.T) {
		input := strings.NewReader("exit")
		out := &writerSpy{}
		calls := 3
		Start(input, out)
		compareInts(t, out.counter, calls)
	})

	t.Run("Count exactly 5 calls for PROMPT, monkeyFace and statement", func(t *testing.T) {
		input := strings.NewReader("let x = 5;")
		out := &writerSpy{}
		calls := 5
		Start(input, out)
		compareInts(t, out.counter, calls)
	})

	t.Run("Count exactly 4 calls for PROMPT, monkeyFace and errors", func(t *testing.T) {
		input := strings.NewReader("& error")
		out := &writerSpy{}
		calls := 4
		Start(input, out)
		compareInts(t, out.counter, calls)
	})
}

func compareInts(t *testing.T, first, second int) {
	t.Helper()
	if first != second {
		t.Errorf("Expected %d to be equal to %d", first, second)
	}
}
