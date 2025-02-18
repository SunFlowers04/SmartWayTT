package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	_ "log"
)

// DB обертка для подключения к базе данных через pgx
type DB struct {
	conn *pgx.Conn
}

// NewDB создает новое подключение к базе данных PostgreSQL с использованием pgx
func NewDB(dsn string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Возвращаем структуру DB с открытым соединением
	return &DB{conn: conn}, nil
}

// Close закрывает соединение с базой данных
func (db *DB) Close() error {
	return db.conn.Close(context.Background())
}

// Ping проверяет соединение с базой данных
func (db *DB) Ping() error {
	err := db.conn.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}
	return nil
}

// Exec выполняет SQL-запросы (например, INSERT, UPDATE, DELETE)
func (db *DB) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.conn.Exec(context.Background(), query, args...)
}

// Query выполняет SQL-запросы для получения данных
func (db *DB) Query(query string, args ...interface{}) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), query, args...)
}

// QueryRow выполняет SQL-запрос и возвращает одну строку
func (db *DB) QueryRow(query string, args ...interface{}) pgx.Row {
	return db.conn.QueryRow(context.Background(), query, args...)
}

// CopyFrom копирует данные из файла в таблицу
