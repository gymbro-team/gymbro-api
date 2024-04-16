package model

import "time"

type Exercise struct {
	ID        int64     `json:id`
	WorkoutID int64     `json:workout_id`
	Name      string    `json:name`
	Icon      string    `json:icon`
	Sets      int       `json:sets`
	Reps      int       `json:reps`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
