-- +migrate Up


-- Вставка изначальных данных в таблицу products
INSERT INTO products (name, code, size, quantity)
VALUES ('Red T-Shirt', 'RTS1', 10, 100),
       ('Blue Jeans', 'BJN2', 20, 200),
       ('Green Sneakers', 'GSN3', 30, 300),
       ('Black Dress', 'BLD4', 25, 150),
       ('White Shirt', 'WTS5', 15, 180);

-- Вставка изначальных данных в таблицу warehouses
INSERT INTO warehouses (name, available)
VALUES ('Main Warehouse', true),
       ('Secondary Warehouse', false),
       ('Regional Warehouse', true),
       ('Temporary Warehouse', true),
       ('Online Warehouse', true);

-- Вставка изначальных данных в таблицу product_warehouses
INSERT INTO product_warehouses (warehouse_id, product_id, quantity)
VALUES (1, 1, 50),
       (2, 1, 30),
       (3, 2, 20),
       (1, 3, 100),
       (4, 4, 70),
       (5, 5, 90);

-- Вставка изначальных данных в таблицу reservation
INSERT INTO reservation (product_id, quantity, warehouse_id)
VALUES (1, 20, 1),
       (2, 10, 2),
       (3, 50, 3),
       (4, 30, 1),
       (5, 40, 4);