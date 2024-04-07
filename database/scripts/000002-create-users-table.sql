create sequence gymbro.seq_users;

create table gymbro.users
(
  id         bigint primary key default nextval('gymbro.seq_users'),
  type       varchar(20)              not null check ( type in ('A','N','P') ),
  username   varchar(50) unique       not null,
  name       text                     not null,
  email      varchar(50) unique       not null,
  password   varchar(50)              not null,
  created_at timestamp with time zone not null default current_timestamp,
  updated_at timestamp with time zone not null default current_timestamp
);

comment on column gymbro.users.type is 'A: Athlete, N: Nutritionist, P: Personal';