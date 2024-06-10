package repository

import (
	"database/sql"
	"errors"
	"gymbro-api/model"
)

type ExerciseRepository struct {
	db *sql.DB
}

var ErrExerciseNotFound = errors.New("exercise not found")

func NewExerciseRepository(db *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{db}
}

func (wr *ExerciseRepository) CreateExercise(exercise *model.Exercise, userId uint64) error {
	_, err := wr.db.Exec(`
		insert into gymbro.exercises(id
								    ,workout_id
								    ,name
								    ,icon
								    ,sets
								    ,reps
								    ,created_at
								    ,updated_at
								    ,created_by
								    ,updated_by)
					         values (nextval('gymbro.seq_exercises')
						            ,$1::bigint
							 	    ,$2::text
							 	    ,$3::text
							 	    ,$4::integer
								    ,$5::integer
								    ,current_timestamp
								    ,current_timestamp
								    ,$6::bigint
								    ,$6::bigint)
	`, exercise.WorkoutID, exercise.Name, exercise.Icon, exercise.Sets, exercise.Reps, userId)

	return err
}

func (wr *ExerciseRepository) GetExerciseByID(id uint64) (*model.Exercise, error) {
	row := wr.db.QueryRow(`
		select e.id
		      ,e.workout_id
			  ,e.name
			  ,e.icon
			  ,e.sets
			  ,e.reps
			  ,e.created_at
			  ,e.updated_at
	     from gymbro.exercises e
	    where e.id = $1::bigint
	`, id)

	exercise := &model.Exercise{}

	err := row.Scan(
		&exercise.ID,
		&exercise.WorkoutID,
		&exercise.Name,
		&exercise.Icon,
		&exercise.Sets,
		&exercise.Reps,
		&exercise.CreatedAt,
		&exercise.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrExerciseNotFound
	}

	return exercise, err
}

func (wr *ExerciseRepository) GetExercises() ([]model.Exercise, error) {
	rows, err := wr.db.Query(`
		select e.id
			  ,e.workout_id
		   	  ,e.name
		   	  ,e.icon
			  ,e.sets
			  ,e.reps
			  ,e.created_at
			  ,e.updated_at
		  from gymbro.exercises e
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	exercises := []model.Exercise{}

	for rows.Next() {
		exercise := model.Exercise{}
		err := rows.Scan(
			&exercise.ID,
			&exercise.WorkoutID,
			&exercise.Name,
			&exercise.Icon,
			&exercise.Sets,
			&exercise.Reps,
			&exercise.CreatedAt,
			&exercise.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

func (wr *ExerciseRepository) UpdateExercise(exercise *model.Exercise) error {
	_, err := wr.db.Exec(`
		update gymbro.exercises
		   set workout_id  = $1::bigint
			  ,name        = $2::text
			  ,icon        = $3::text
			  ,sets        = $4::integer
			  ,reps        = $5::integer
			  ,updated_at  = current_timestamp
			  ,updated_by  = null
		 where id          = $6::bigint
	`, exercise.WorkoutID, exercise.Name, exercise.Icon, exercise.Sets, exercise.Reps, exercise.ID)

	if err == sql.ErrNoRows {
		return ErrExerciseNotFound
	}

	return err
}

func (wr *ExerciseRepository) DeleteExercise(id uint64) error {
	_, err := wr.db.Exec(`
		delete
		  from gymbro.exercises
		 where id = $1::bigint
	`, id)

	if err == sql.ErrNoRows {
		return ErrExerciseNotFound
	}

	return err
}
