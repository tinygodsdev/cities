package cities

import "time"

type City struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}
