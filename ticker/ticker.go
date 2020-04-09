package ticker

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//RunningConfig ...
type RunningConfig struct {
	SecondPerMinute int
	SecondPerHour   int
	AllowUpdate     time.Duration
	Deadline        time.Duration
	SecondMessage   string
	MinuteMessage   string
	HourMessage     string
	Port            string
}

//ClockWriter ...
func ClockWriter(config *RunningConfig, message chan string) {
	//create a second ticker
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	//deadline timer to kill the clock after 3 hours
	deadline := time.NewTimer(config.Deadline)
	defer deadline.Stop()

	//spinHandler timer to accept the http request
	spinHandler := time.NewTimer(config.AllowUpdate)
	defer spinHandler.Stop()

	//counter..
	counter := 0
	for {
		select {
		case <-ticker.C:
			counter++
			if counter/config.SecondPerHour != 0 && counter%config.SecondPerHour == 0 {
				message <- config.HourMessage
			} else if counter/config.SecondPerMinute != 0 && counter%config.SecondPerMinute == 0 {
				message <- config.MinuteMessage
			} else {
				message <- config.SecondMessage
			}

		case <-deadline.C:
			//close message channel
			close(message)
			return

		case <-spinHandler.C:
			router := mux.NewRouter()
			router.HandleFunc("/secondmessage/{s_Msg}", config.updateSecondMessage).Methods(http.MethodGet)
			router.HandleFunc("/minutemessage/{m_Msg}", config.updateMinuteMessage).Methods(http.MethodGet)
			router.HandleFunc("/hourmessage/{h_Msg}", config.updateHourMessage).Methods(http.MethodGet)

			//spin a server at 8080
			go http.ListenAndServe(":"+config.Port, router)
		}
	}
}

func (config *RunningConfig) updateSecondMessage(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	message := vars["s_Msg"]
	config.SecondMessage = message
	fmt.Fprintf(response, "updated second message...")
}

func (config *RunningConfig) updateMinuteMessage(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	message := vars["m_Msg"]
	config.MinuteMessage = message
	fmt.Fprintf(response, "updated second message...")
}

func (config *RunningConfig) updateHourMessage(response http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	message := vars["h_Msg"]
	config.HourMessage = message
	fmt.Fprintf(response, "updated second message...")
}
