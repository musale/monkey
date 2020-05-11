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
	t.Run("Count exactly 1 call for PROMPT", func(t *testing.T) {
		input := strings.NewReader("")
		out := &writerSpy{}
		calls := 1
		Start(input, out)
		compareInts(t, out.counter, calls)
	})

	t.Run("Count exactly 2 calls for PROMPT and exit", func(t *testing.T) {
		input := strings.NewReader("exit")
		out := &writerSpy{}
		calls := 2
		Start(input, out)
		compareInts(t, out.counter, calls)
	})

	t.Run("Count 14 calls to print tokens and 2 for the PROMPT", func(t *testing.T) {
		expect := "let add = fn(x, y){x + y}"
		input := strings.NewReader(expect)

		out := &writerSpy{}
		// 14 calls to token and 2 calls to PROMPT
		calls := 16
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
