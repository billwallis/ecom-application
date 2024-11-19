create table if not exists addresses (
    id int unsigned not null auto_increment,
    user_id int unsigned not null,
    is_default boolean not null,
    line_1 varchar(50) not null,
    line_2 varchar(50) not null,
    city varchar(50) not null,
    country varchar(50) not null,
    postcode varchar(12) not null,
    created_at timestamp not null default current_timestamp,

    primary key (id),
    foreign key (user_id) references users(id)
)
