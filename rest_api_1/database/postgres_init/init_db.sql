create table if not exists product(
	id SERIAL primary key,
	product_name VARCHAR(50) not null,
	price numeric(10,2) not null
);


INSERT INTO product (product_name, price)
VALUES
('Smartphone X', 999.99),
('Laptop Pro 15', 1499.99),
('Wireless Earbuds', 199.95),
('Gaming Console', 499.99),
('Smartwatch Series 5', 299.99);
