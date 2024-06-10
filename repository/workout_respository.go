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

func NewWorkoutRepository(db *sql.DB) *WorkoutRepository {
	return &WorkoutRepository{db}
}

func (wr *WorkoutRepository) CreateWorkout(workout *model.Workout, userId uint64) error {
	_, err := wr.db.Exec(`
		insert into gymbro.workouts(id
								   ,athlete_id
								   ,personal_id
								   ,name
								   ,icon
								   ,cover_image
								   ,description
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
								   ,$6::text
								   ,$7::bpchar
								   ,0
								   ,0
								   ,0
								   ,current_timestamp
								   ,current_timestamp
								   ,$8::bigint
								   ,$8::bigint)
	`, workout.AthleteID, workout.PersonalID, workout.Name, workout.Icon, workout.CoverImage, workout.Description, workout.WeekDay, userId)

	return err
}

func (wr *WorkoutRepository) GetWorkoutByID(id uint64, userId uint64) (*model.Workout, error) {
	row := wr.db.QueryRow(`
		select w.id
		      ,w.athlete_id
			  ,w.personal_id
			  ,w.name
			  ,w.icon
			  ,w.description
			  ,w.cover_image
			  ,w.week_day
			  ,w.sets_count
			  ,w.reps_count
			  ,w.exercises_count
			  ,w.created_at
			  ,w.updated_at
	     from gymbro.workouts w
	    where w.id = $1::bigint
		  and (w.athlete_id = $2::bigint
		   or w.personal_id = $2::bigint
		   or w.created_by = $2::bigint)
	`, id, userId)

	workout := &model.Workout{}

	err := row.Scan(
		&workout.ID,
		&workout.AthleteID,
		&workout.PersonalID,
		&workout.Name,
		&workout.Icon,
		&workout.Description,
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

func (wr *WorkoutRepository) GetWorkouts(userId uint64) ([]model.Workout, error) {
	rows, err := wr.db.Query(`
		select w.id
		      ,w.athlete_id
			  ,w.personal_id
			  ,w.name
			  ,w.icon
			  ,w.cover_image
			  ,w.description
			  ,w.week_day
			  ,w.sets_count
			  ,w.reps_count
			  ,w.exercises_count
			  ,w.created_at
			  ,w.updated_at
	     from gymbro.workouts w
		where w.athlete_id = $1::bigint
		   or w.personal_id = $1::bigint
		   or w.created_by = $1::bigint
	`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := []model.Workout{}

	for rows.Next() {
		workout := model.Workout{}
		err := rows.Scan(
			&workout.ID,
			&workout.AthleteID,
			&workout.PersonalID,
			&workout.Name,
			&workout.Icon,
			&workout.CoverImage,
			&workout.Description,
			&workout.WeekDay,
			&workout.SetsCount,
			&workout.RepsCount,
			&workout.ExercisesCount,
			&workout.CreatedAt,
			&workout.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		workouts = append(workouts, workout)
	}

	return workouts, nil
}

func (wr *WorkoutRepository) UpdateWorkout(workout *model.Workout, userId uint64) error {
	_, err := wr.db.Exec(`
		update gymbro.workouts
		   set athlete_id  = $1::bigint
		      ,personal_id = $2::bigint
			  ,name        = $3::text
			  ,icon        = $4::text
			  ,cover_image = $5::text
			  ,description = $6::text
			  ,week_day    = $7::bpchar
			  ,updated_at  = current_timestamp
			  ,updated_by  = null
		 where id          = $8::bigint
		   and (athlete_id = $9::bigint
			or personal_id = $9::bigint
			or created_by = $9::bigint) 
	`, workout.AthleteID, workout.PersonalID, workout.Name, workout.Icon, workout.CoverImage, workout.Description, workout.WeekDay, workout.ID, userId)

	if err == sql.ErrNoRows {
		return ErrWorkoutNotFound
	}

	return err
}

func (wr *WorkoutRepository) DeleteWorkout(id uint64, userId uint64) error {
	_, err := wr.db.Exec(`
		delete
		  from gymbro.workouts
		 where id = $1::bigint
		   and (athlete_id = $2::bigint
			or personal_id = $2::bigint
			or created_by = $2::bigint)
	`, id, userId)

	if err == sql.ErrNoRows {
		return ErrWorkoutNotFound
	}

	return err
}
