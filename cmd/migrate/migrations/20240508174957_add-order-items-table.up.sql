create table if not exists order_items (
    id int unsigned not null auto_increment,
    order_id int unsigned not null,
    product_id int unsigned not null,
    quantity int not null,
    price decimal(10, 2) not null,

    primary key (id),
    foreign key (order_id) references orders(id),
    foreign key (product_id) references products(id)
)
