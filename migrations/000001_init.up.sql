-- +migrate Up
CREATE TABLE products
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    code     VARCHAR(50) UNIQUE,
    size     INT,
    quantity INT DEFAULT 0
);

CREATE TABLE warehouses
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(255),
    available BOOLEAN
);

CREATE TABLE product_warehouses
(
    id           SERIAL PRIMARY KEY,
    warehouse_id INTEGER REFERENCES warehouses (id),
    product_id   INTEGER REFERENCES products (id),
    quantity     INT DEFAULT 0,
    UNIQUE (product_id, warehouse_id)
);

CREATE TABLE reservation
(
    id           SERIAL PRIMARY KEY,
    product_id   INTEGER REFERENCES products (id),
    quantity     INT DEFAULT 0,
    warehouse_id INTEGER REFERENCES warehouses (id)
);
