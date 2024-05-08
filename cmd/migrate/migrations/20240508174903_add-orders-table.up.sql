create table if not exists orders (
    id int unsigned not null auto_increment,
    user_id int unsigned not null,
    total decimal(10, 2) not null,
    status enum('pending', 'completed', 'cancelled') not null default 'pending',
    address text not null,
    created_at timestamp not null default current_timestamp,

    primary key (id),
    foreign key (user_id) references users(id)
);
