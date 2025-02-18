-- +goose Up
-- Создание таблицы passengers
CREATE TABLE passengers (
                            id          VARCHAR(255) PRIMARY KEY,
                            last_name   VARCHAR(255) NOT NULL,
                            first_name  VARCHAR(255) NOT NULL,
                            middle_name VARCHAR(255) DEFAULT NULL,
                            flight_id   VARCHAR(255) NOT NULL,
                            FOREIGN KEY (flight_id) REFERENCES flights(id) ON DELETE CASCADE
);

-- Создание индекса
CREATE INDEX passengers_flight_id_idx ON passengers (flight_id);

-- +goose Down
-- Удаление индекса
DROP INDEX IF EXISTS passengers_flight_id_idx;

-- Удаление таблицы passengers
DROP TABLE IF EXISTS passengers;
