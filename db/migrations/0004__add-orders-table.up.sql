create table if not exists ecom.orders (
    id         int primary key generated always as identity,
    user_id    int not null references ecom.users(id),
    total      decimal(10, 2) not null,
    status     varchar not null default 'pending' check (status in ('pending', 'completed', 'cancelled')),
    address    text not null,
    created_at timestamp not null default current_timestamp
)
