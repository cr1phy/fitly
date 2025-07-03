-- Создание таблицы категорий продуктов
CREATE TABLE IF NOT EXISTS product_category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);

-- Создание таблицы продуктов
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES product_category(id),
    name VARCHAR(128) NOT NULL,
    description TEXT,
    calories FLOAT,
    fats FLOAT,
    protein FLOAT,
    carbs FLOAT
);

-- Создание таблицы категорий блюд
CREATE TABLE IF NOT EXISTS dish_category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);

-- Создание таблицы блюд
CREATE TABLE IF NOT EXISTS dishes (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES dish_category(id),
    name VARCHAR(128) NOT NULL,
    description TEXT
);
