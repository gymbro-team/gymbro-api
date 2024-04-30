package repository

import (
	"database/sql"
	"errors"
	"gymbro-api/model"
)

type WorkoutRepository struct {
	db *sql.DB
}

var ErrWorkoutNotFound = errors.New("workout not found")

func (wr *WorkoutRepository) CreateWorkout(workout *model.Workout) error {
	_, err := wr.db.Exec(`
		insert into gymbro.workouts(id
								   ,athlete_id
								   ,personal_id
								   ,name
								   ,icon
								   ,cover_image
								   ,week_day
								   ,sets_count
								   ,reps_count
								   ,exercises_count
								   ,created_at
								   ,updated_at
								   ,created_by
								   ,updated_by)
					        values (nextval('gymbro.seq_workouts')
						           ,$1::bigint
								   ,$2::bigint
								   ,$3::text
								   ,$4::text
								   ,$5::text
								   ,$6::bpchar
								   ,0
								   ,0
								   ,0
								   ,current_timestamp
								   ,current_timestamp
								   ,null
								   ,null)
	`, workout.AthleteID, workout.PersonalID, workout.Name, workout.Icon, workout.CoverImage, workout.WeekDay)

	return err
}

func (wr *WorkoutRepository) GetWorkoutByID(id uint64) (*model.Workout, error) {
	row := wr.db.QueryRow(`
		select w.id
		      ,w.athlete_id
			  ,w.personal_id
			  ,w.name
			  ,w.icon
			  ,w.cover_image
			  ,w.week_day
			  ,w.sets_count
			  ,w.reps_count
			  ,w.exercises_count
			  ,w.created_at
			  ,w.updated_at
	     from gymbro.workouts w
	    where w.id = $1::bigint
	`, id)

	workout := &model.Workout{}

	err := row.Scan(
		&workout.ID,
		&workout.AthleteID,
		&workout.PersonalID,
		&workout.Name,
		&workout.Icon,
		&workout.CoverImage,
		&workout.WeekDay,
		&workout.SetsCount,
		&workout.RepsCount,
		&workout.ExercisesCount,
		&workout.CreatedAt,
		&workout.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrWorkoutNotFound
	}

	return workout, err
}
