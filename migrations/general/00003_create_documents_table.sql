-- +goose Up
-- Создание таблицы documents
CREATE TABLE documents (
                           id            VARCHAR(255) PRIMARY KEY,
                           document_type VARCHAR(255) NOT NULL,
                           document_num  VARCHAR(255) NOT NULL,
                           passenger_id  VARCHAR(255) NOT NULL,
                           FOREIGN KEY (passenger_id) REFERENCES passengers(id) ON DELETE CASCADE
);

-- Создание индекса
CREATE INDEX documents_passenger_id_idx ON documents (passenger_id);

-- +goose Down
-- Удаление индекса
DROP INDEX IF EXISTS documents_passenger_id_idx;

-- Удаление таблицы documents
DROP TABLE IF EXISTS documents;
