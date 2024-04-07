create sequence gymbro.seq_exercises;

create table gymbro.exercises
(
  id          bigint primary key default nextval('gymbro.seq_exercises'),
  workout_id  bigint references gymbro.workouts(id),
  name        text                               not null,
  icon        text                               not null,
  sets        integer                            not null,
  reps        integer                            not null,
  created_at  timestamp with time zone           not null default current_timestamp,
  updated_at  timestamp with time zone           not null default current_timestamp,
  created_by  bigint references gymbro.users(id) not null,
  updated_by  bigint references gymbro.users(id) not null
);