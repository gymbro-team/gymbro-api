package model

import "time"

type Exercise struct {
	ID        uint64    `json:"id"`
	WorkoutID uint64    `json:"workout_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Sets      int       `json:"sets"`
	Reps      int       `json:"reps"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
