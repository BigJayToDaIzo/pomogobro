package timer

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.SleepFunc(s.Duration)
}

type Timer struct {
	IsCounting  bool
	IsCompleted bool
	Writer      io.Writer
	Sleeper     Sleeper
}

// configuration points temp set to const
const countdownDurationInSeconds = 3
const completeMsg = "Pomogobro!"

func NewTimer(w io.Writer, s Sleeper) Timer {
	return Timer{
		IsCounting:  false,
		IsCompleted: false,
		Writer:      w,
		Sleeper:     s,
	}
}

func (t *Timer) Start() {
	t.IsCounting = true
	for t.IsCounting {
		for i := countdownDurationInSeconds; i > 0; i-- {
			t.Writer.Write([]byte(fmt.Sprintf("%d\n", i)))
			t.Sleeper.Sleep()
		}
		t.Writer.Write([]byte(fmt.Sprintf("%s\n", completeMsg)))
		t.IsCompleted = true
		t.Stop()

	}
}

func (t *Timer) Stop() {
	t.IsCounting = false
}

func (t *Timer) Toggle() {
	switch t.IsCounting {
	case true:
		t.Stop()
	case false:
		t.Start()
	}
}
