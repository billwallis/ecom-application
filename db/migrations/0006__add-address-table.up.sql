create table if not exists ecom.addresses (
    id         int primary key generated always as identity,
    user_id    int not null references ecom.users(id),
    is_default boolean not null,
    line_1     varchar(50) not null,
    line_2     varchar(50) not null,
    city       varchar(50) not null,
    country    varchar(50) not null,
    postcode   varchar(12) not null,
    created_at timestamp not null default current_timestamp
)
