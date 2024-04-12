create sequence gymbro.seq_user_devices;

create table gymbro.user_devices
(
  id         bigint primary key default nextval('gymbro.seq_user_devices'),
  user_id    bigint references gymbro.users(id)   not null,
  device_id  bigint references gymbro.devices(id) not null,
  created_at timestamp with time zone             not null default current_timestamp,
  updated_at timestamp with time zone             not null default current_timestamp,
  deleted_at timestamp with time zone
);