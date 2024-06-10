create sequence gymbro.seq_users;

create table gymbro.users
(
  id         bigint primary key default nextval('gymbro.seq_users'),
  type       varchar(20)              not null check ( type in ('A','N','P') ),
  username   varchar(50) unique       not null,
  name       text                     not null,
  email      varchar(50) unique       not null,
  password   varchar(200)              not null,
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
  if (coalesce(new.password, '') != coalesce(old.password, '')) then
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

insert into gymbro.users (id, type, username, name, email, password, status, created_at, updated_at, deleted_at)
values (DEFAULT, 'A'::varchar(20), 'teste'::varchar(50), 'Usuário Teste'::text, 'teste@teste.com'::varchar(50),
        '1234'::varchar(200), 'A'::char, DEFAULT, DEFAULT, null::timestamp with time zone);
