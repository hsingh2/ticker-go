package ticker

import (
	"testing"
	"time"
)

func TestClockWriter(t *testing.T) {
	message := make(chan string)

	tests := []struct {
		config         *RunningConfig
		expectedOutput []string
		actual         chan string
	}{
		{config: &RunningConfig{AllowUpdate: 1 * time.Minute, SecondPerMinute: 3, SecondPerHour: 3 * 3, Deadline: 9 * time.Second, SecondMessage: "tick", MinuteMessage: "tock", HourMessage: "bong"},
			expectedOutput: []string{"tick", "tick", "tock", "tick", "tick", "tock", "tick", "tick", "bong"},
			actual:         message},
	}

	for _, test := range tests {
		i := 0
		go ClockWriter(test.config, message)

		for msg := range test.actual {
			if msg != test.expectedOutput[i] {
				t.Errorf("Test failed, expectedvalue: %v, actualvalue: %v", test.expectedOutput[i], msg)
			}
			i++
		}
	}
}
