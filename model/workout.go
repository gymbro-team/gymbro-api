package model

import "time"

type Workout struct {
	ID             uint64    `json:"id"`
	AthleteID      *uint64   `json:"athlete_id"`
	PersonalID     *uint64   `json:"personal_id"`
	Name           string    `json:"name"`
	Icon           string    `json:"icon"`
	CoverImage     string    `json:"cover_image"`
	WeekDay        byte      `json:"week_day"`
	SetsCount      int       `json:"sets_count"`
	RepsCount      int       `json:"reps_count"`
	ExercisesCount int       `json:"exercises_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
