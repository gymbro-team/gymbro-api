create sequence gymbro.seq_sessions;

create table gymbro.sessions
(
  id         bigint primary key default nextval('gymbro.seq_sessions'),
  user_id    bigint references gymbro.users(id) not null,
  token      text                               not null,
  status     bpchar(1)                          not null default 'A' check ( status in ('A','I') ),
  expires_at timestamp with time zone           not null,
  created_at timestamp with time zone           not null default current_timestamp,
  updated_at timestamp with time zone           not null default current_timestamp
);