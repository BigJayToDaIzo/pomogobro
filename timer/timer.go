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
	SleepDuration time.Duration
	SleepFunc     func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.SleepFunc(s.SleepDuration)
}

type Timer struct {
	IsCounting     bool
	RemainingTicks int
	IsCompleted    bool
	Writer         io.Writer
	Sleeper        Sleeper
}

// configuration points temp set to const
const completeMsg = "Pomogobro!"

func NewTimer(w io.Writer, s Sleeper, t int) *Timer {
	return &Timer{
		IsCounting:     false,
		RemainingTicks: t,
		IsCompleted:    false,
		Writer:         w,
		Sleeper:        s,
	}
}

func (t *Timer) Start() {
	t.IsCounting = true
	for t.IsCounting {
		for i := t.RemainingTicks; i > 0; i-- {
			t.Writer.Write([]byte(fmt.Sprintf("%d\n", i)))
			t.RemainingTicks--
			t.Sleeper.Sleep()
		}
		t.Stop()
		t.IsCompleted = true
		t.Writer.Write([]byte(fmt.Sprintf("%s\n", completeMsg)))
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
