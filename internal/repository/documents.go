package repository

import (
	"context"
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	"github.com/jackc/pgx/v5"
)

type StorageDoc interface {
	Create(document entities.Document) error
	GetAll() ([]entities.Document, error)
}

type storageDoc struct {
	pg *pgx.Conn
}

func NewStorageDocument(pg *pgx.Conn) StorageDoc {
	return &storageDoc{pg: pg}
}

func (s *storageDoc) Create(document entities.Document) error {
	_, err := s.pg.Exec(context.Background(), "INSERT INTO documents (document_type, document_num, ...) VALUES ($1, $2, ...)", document.DocumentType)
	return err
}

func (s *storageDoc) GetAll() ([]entities.Document, error) {
	rows, err := s.pg.Query(context.Background(), "SELECT * FROM documents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []entities.Document
	for rows.Next() {
		var document entities.Document
		if err := rows.Scan(&document.DocumentType); err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}
