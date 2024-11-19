create table if not exists ecom.users (
    id         int primary key generated always as identity,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp not null default current_timestamp
)
