package repository

import (
	"context"
	"errors"
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	"github.com/jackc/pgx/v5"
)

type StorageDoc interface {
	Create(ctx context.Context, doc entities.Document) error
	GetByPassengerID(ctx context.Context, passengerID string) ([]entities.Document, error)
	Update(ctx context.Context, doc entities.Document) error
	Delete(ctx context.Context, docID string) error
}

// Реализация метода в конкретном хранилище (например, с БД)
func (s *DocumentRepository) GetByPassengerID(ctx context.Context, passengerID string) ([]entities.Document, error) {
	var documents []entities.Document

	// SQL-запрос для получения документов
	query := `SELECT id, passenger_id, document_type, document_number FROM documents WHERE passenger_id = $1`
	rows, err := s.pg.Query(ctx, query, passengerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Заполняем список документов
	for rows.Next() {
		var doc entities.Document
		if err := rows.Scan(&doc.DocumentID, &doc.PassengerID, &doc.DocumentType, &doc.DocumentNumber); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	if len(documents) == 0 {
		return nil, errors.New("документы не найдены")
	}

	return documents, nil
}

func (s *DocumentRepository) Update(ctx context.Context, doc entities.Document) error {
	query := `UPDATE documents SET document_type = $1, document_number = $2 WHERE id = $3`
	_, err := s.pg.Exec(ctx, query, doc.DocumentType, doc.DocumentNumber, doc.DocumentID)
	if err != nil {
		return err
	}
	return nil
}

type DocumentRepository struct {
	pg *pgx.Conn
}

func (s *DocumentRepository) Delete(ctx context.Context, docID string) error {
	query := `DELETE FROM documents WHERE id = $1`
	_, err := s.pg.Exec(ctx, query, docID)
	if err != nil {
		return err
	}
	return nil
}

func NewStorageDocument(pg *pgx.Conn) StorageDoc {
	return &DocumentRepository{pg: pg}
}

func (s *DocumentRepository) Create(ctx context.Context, document entities.Document) error { //todo исправить репозиторий
	_, err := s.pg.Exec(context.Background(), "INSERT INTO documents (document_type, document_num, ...) VALUES ($1, $2, ...)", document.DocumentType)
	return err
}
