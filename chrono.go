package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

func chrono() {
	tickInterval := time.Millisecond
	ticker := time.NewTicker(tickInterval)

	spinner := wow.New(os.Stdout, spin.Get(spin.Clock), "")
	spinner.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		spinner.Spinner(spin.Spinner{Frames: []string{""}, Interval: 10})
		spinner.Persist()
		spinner.Stop()
		os.Exit(0)
	}()

	var t time.Duration
	for _ = range ticker.C {
		t += tickInterval
		spinner.Text(t.String())
	}
}
