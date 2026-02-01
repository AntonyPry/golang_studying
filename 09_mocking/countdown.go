package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord = "Go!"
	sleep     = "sleep"
	write     = "write"

	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

// incorrect version to fail the test
//func Countdown(out io.Writer, sleeper Sleeper) {
//	for i := countdownStart; i > 0; i-- {
//		sleeper.Sleep()
//	}
//
//	for i := countdownStart; i > 0; i-- {
//		fmt.Fprintln(out, i)
//	}
//
//	fmt.Fprint(out, finalWord)
//}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 5 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
