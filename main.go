package main

import (
	"flag"
	"fmt"
	"time"

	"tickerclock/ticker"
)

var (
	// all the flags represents second value for time.Duration
	secondPerMin  = flag.Int("secPMin", 60, "represents how many seconds are there in a minute, used to compute tock message")
	secondPerHour = flag.Int("secPHour", 3600, "represents how many seconds are there in a hour, used for bong message")
	deadline      = flag.Int("deadline", 108000, "represents how many seconds clock should run, by default 3 hours")
	allowUpdate   = flag.Int("allowUpdate", 600, "represents when to start letting update the print message")
	port          = flag.String("port", "8080", "expose the http endpoint to server the update message request")
)

func main() {
	flag.Parse()
	//set up running config
	config := &ticker.RunningConfig{
		SecondPerMinute: *secondPerMin,
		SecondPerHour:   *secondPerHour,
		AllowUpdate:     time.Duration(*allowUpdate) * time.Second,
		Deadline:        time.Duration(*deadline) * time.Second,
		SecondMessage:   "tick",
		MinuteMessage:   "tock",
		HourMessage:     "bong",
		Port:            *port,
	}

	message := make(chan string)
	go ticker.ClockWriter(config, message)

	for {
		msg, ok := <-message
		if !ok {
			return
		}
		fmt.Println(msg)
	}
}
