-- +goose Up
-- Создание типов
CREATE TYPE flight_status AS ENUM ('booked', 'completed', 'cancelled');

-- Создание таблицы flights
CREATE TABLE flights (
                         id              VARCHAR(255) PRIMARY KEY,
                         departure_point VARCHAR(255) NOT NULL,
                         destination     VARCHAR(255) NOT NULL,
                         order_number    VARCHAR(255) NOT NULL,
                         supplier        VARCHAR(255) NOT NULL,
                         departure_date  TIMESTAMP NOT NULL,
                         arrival_date    TIMESTAMP NOT NULL,
                         service_date    TIMESTAMP NOT NULL,
                         status          flight_status NOT NULL DEFAULT 'booked'
);

-- Создание индекса
CREATE INDEX flights_order_number_idx ON flights (order_number);

-- Создание индекса для поиска по датам
CREATE INDEX flights_departure_date_idx ON flights (departure_date);

-- +goose Down
-- Удаление индексов
DROP INDEX IF EXISTS flights_order_number_idx;
DROP INDEX IF EXISTS flights_departure_date_idx;

-- Удаление таблицы flights
DROP TABLE IF EXISTS flights;

-- Удаление типа flight_status
DROP TYPE IF EXISTS flight_status;
