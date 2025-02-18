package migrations

import (
	"embed"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq" // PostgreSQL driver
)

//go:embed general/*.sql
var embedMigrations embed.FS

var ErrDBInstanceIsNil = errors.New("db instance is nil")

// InitMigrate подключается к базе данных и накатывает миграции.
func InitMigrate(instance *pgx.Conn) error {
	if instance == nil {
		return ErrDBInstanceIsNil
	}
	// Устанавливаем файловую систему для миграций
	goose.SetBaseFS(embedMigrations)

	// Преобразуем *pgx.Conn в *sql.DB
	sqlDB := stdlib.OpenDB(*instance.Config())

	// Устанавливаем диалект базы данных
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	// Накатываем миграции
	if err := goose.Up(sqlDB, "general"); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
