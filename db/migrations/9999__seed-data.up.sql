/*
    This isn't a "real" migration, just a way to chuck some data into the
    database while I'm still working on this project
*/

insert into ecom.products(id, name, description, image, price, quantity)
values
    (default, 'Product 1', 'Description 1', 'image1.jpg', 100, 10),
    (default, 'Product 2', 'Description 2', 'image2.jpg', 200, 20),
    (default, 'Product 3', 'Description 3', 'image3.jpg', 300, 30),
    (default, 'Product 4', 'Description 4', 'image4.jpg', 400, 40),
    (default, 'Product 5', 'Description 5', 'image5.jpg', 500, 50),
    (default, 'Product 6', 'Description 6', 'image6.jpg', 600, 60),
    (default, 'Product 7', 'Description 7', 'image7.jpg', 700, 70),
    (default, 'Product 8', 'Description 8', 'image8.jpg', 800, 80),
    (default, 'Product 9', 'Description 9', 'image9.jpg', 900, 90)
