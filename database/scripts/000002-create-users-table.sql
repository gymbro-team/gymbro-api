create sequence gymbro.seq_users;

create table gymbro.users
(
  id         bigint primary key default nextval('gymbro.seq_users'),
  type       varchar(20)              not null check ( type in ('A','N','P') ),
  username   varchar(50) unique       not null,
  name       text                     not null,
  email      varchar(50) unique       not null,
  password   varchar(50)              not null,
  status     bpchar(1)                not null default 'A' check ( status in ('A','I','D') ),
  created_at timestamp with time zone not null default current_timestamp,
  updated_at timestamp with time zone not null default current_timestamp,
  deleted_at timestamp with time zone
);

comment on column gymbro.users.type is 'A: Athlete, N: Nutritionist, P: Personal';
comment on column gymbro.users.status is 'A: Active, I: Inactive, D: Deleted';

create or replace function fnc_trg_users_biu_er() returns trigger as $$
begin
  --
  if (new.password != old.password) then
    --
    new.password := crypt(new.password, gen_salt('bf'));
    --
  end if;
  --
  return new;
  --
end;
$$ language plpgsql;

create trigger trg_users_biu_er
  before insert or update on gymbro.users
  for each row execute function fnc_trg_users_biu_er();
