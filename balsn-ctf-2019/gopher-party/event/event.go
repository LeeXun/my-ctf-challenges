package event

import (
	"balsnctf/gopherparty/config"
	"os"
	"time"
)

// Event should have comment or be unexported
type Event struct {
	StartAt int
	EndAt   int
}

var events = []Event{
	Event{
		StartAt: 0,
		EndAt:   60,
	},
}

func init() {
	if os.Getenv("APP_ENV") == config.App.Production {
		events = generateEvents()
	}
}

func generateEvents() []Event {
	e := []Event{}
	playTime := 2
	brakeTime := 3
	for i := 0; i < 60; i += playTime + brakeTime {
		e = append(e, Event{
			StartAt: i,
			EndAt:   i + playTime,
		})
	}
	return e
}

func IsUp() bool {
	_, minutes, _ := time.Now().Clock()
	var isUp bool
	for _, event := range events {
		if minutes >= event.StartAt && minutes < event.EndAt {
			isUp = true
			break
		}
	}
	return isUp
}

func IsDown() bool {
	return !IsUp()
}
