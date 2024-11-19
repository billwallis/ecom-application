create table if not exists ecom.order_items (
    id         int primary key generated always as identity,
    order_id   int not null references ecom.orders(id),
    product_id int not null references ecom.products(id),
    quantity   int not null,
    price      decimal(10, 2) not null
)
