create table if not exists ecom.products (
    id          int primary key generated always as identity,
    name        varchar(255) not null,
    description text not null,
    image       varchar(255) not null,
    price       decimal(10, 2) not null,
    quantity    int not null,
    created_at  timestamp not null default current_timestamp
)
