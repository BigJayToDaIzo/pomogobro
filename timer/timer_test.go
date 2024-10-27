package timer

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyCountdownOps struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *SpyCountdownOps) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
func TestTimer(t *testing.T) {
	assertCounting := func(t *testing.T, timer Timer, want bool, errMsg string) {
		t.Helper()
		if timer.IsCounting != want {
			t.Errorf("%s, got %v want %v", errMsg, timer.IsCounting, want)
		}
	}
	t.Run("timer returns IsCompleted", func(*testing.T) {
		b := &bytes.Buffer{}
		duration := 0 * time.Second
		durationSpy := SpyTime{}
		s := &ConfigurableSleeper{duration, durationSpy.Sleep}
		timer := NewTimer(b, s)
		timer.Start()
		// TODO: FRAGILE, timer interruption needs to be coded to fix
		if timer.IsCompleted != true {
			t.Errorf("got %v want %v", timer.IsCompleted, true)
		}
	})
	t.Run("timer not counting", func(t *testing.T) {
		b := &bytes.Buffer{}
		duration := 0 * time.Second
		durationSpy := SpyTime{}
		s := &ConfigurableSleeper{duration, durationSpy.Sleep}
		timer := NewTimer(b, s)
		assertCounting(t, timer, false,
			"Timer should not automatically begin counting when instantiated")
		timer.Start()
		timer.Stop()
		assertCounting(t, timer, false,
			"Timer should not be couting after Stop()")
	})
	t.Run("timer should toggle", func(t *testing.T) {
		b := &bytes.Buffer{}
		duration := 0 * time.Second
		durationSpy := SpyTime{}
		s := &ConfigurableSleeper{duration, durationSpy.Sleep}
		timer := NewTimer(b, s)
		timer.Toggle()
		// TODO: FRAGILE, timer interruption needs to be coded to fix
		if timer.IsCompleted != true {
			t.Errorf("got %v want %v", timer.IsCompleted, true)
		}
		timer.Toggle()
		assertCounting(t, timer, false,
			"Timer should stop counting after second Toggle()")
	})
	t.Run("timer should print 321Pomogobro!", func(t *testing.T) {
		b := &bytes.Buffer{}
		duration := 0 * time.Second
		durationSpy := SpyTime{}
		s := &ConfigurableSleeper{duration, durationSpy.Sleep}
		timer := NewTimer(b, s)
		timer.Start()
		got := b.String()
		want := `3
2
1
Pomogobro!
`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("timer order of ops", func(t *testing.T) {
		s := &SpyCountdownOps{}
		timer := NewTimer(s, s)
		timer.Start()
		want := []string{write, sleep, write, sleep, write, sleep, write}
		if !reflect.DeepEqual(s.Calls, want) {
			t.Errorf("got %v want %v", s.Calls, want)
		}

	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("configurable sleeper", func(t *testing.T) {
		duration := 5 * time.Second
		durationSpy := &SpyTime{}
		s := &ConfigurableSleeper{duration, durationSpy.Sleep}
		s.Sleep()
		if durationSpy.durationSlept != duration {
			t.Errorf("Should have slept for %v but slept for %v", duration, durationSpy.durationSlept)
		}
	})
}
