package main

import (
	"os"
	"time"

	"example.com/pomogobro/timer"
)

func main() {
	// test stdout
	s := &timer.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	// s := timer.ConfigurableSleeper{Duration: 333 * time.Millisecond, SleepFunc: time.Sleep}
	t := timer.NewTimer(os.Stdout, s)
	t.Start()
	// test buffer
	// b := &bytes.Buffer{}
	// t2 := timer.NewTimer(b, s)
	// t2.Start()
	// fmt.Println(b.String())
	// api endpoint?

	// disk /tmp writes?
}
