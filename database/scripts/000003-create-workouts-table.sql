create sequence gymbro.seq_workouts;

create table gymbro.workouts
(
  id          bigint primary key default nextval('gymbro.seq_workouts'),
  athlete_id  bigint references gymbro.users(id),
  personal_id bigint references gymbro.users(id),
  name        text                               not null,
  icon        text                               not null,
  cover_image text                               not null,
  week_day    bpchar                             not null check ( week_day in ('0','1','2','3','4','5','6') ),
  sets_count      integer                        not null,
  reps_count      integer                        not null,
  exercises_count integer                        not null,
  created_at  timestamp with time zone           not null default current_timestamp,
  updated_at  timestamp with time zone           not null default current_timestamp,
  created_by  bigint references gymbro.users(id) not null,
  updated_by  bigint references gymbro.users(id) not null
);
