package repository

import (
	"context"
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	"github.com/jackc/pgx/v5"
)

type StoragePass interface {
	Create(passenger entities.Passenger) error
	GetAll() ([]entities.Passenger, error)
}

type storagePass struct {
	pg *pgx.Conn
}

func NewStoragePass(pg *pgx.Conn) StoragePass {
	return &storagePass{pg: pg}
}

func (s *storagePass) Create(passenger entities.Passenger) error {
	_, err := s.pg.Exec(context.Background(), "INSERT INTO passengers (last_name, first_name, ...) VALUES ($1, $2, ...)", passenger.LastName, passenger.FirstName)
	return err
}

func (s *storagePass) GetAll() ([]entities.Passenger, error) {
	rows, err := s.pg.Query(context.Background(), "SELECT * FROM passengers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passengers []entities.Passenger
	for rows.Next() {
		var passenger entities.Passenger
		if err := rows.Scan(&passenger.LastName, &passenger.FirstName); err != nil {
			return nil, err
		}
		passengers = append(passengers, passenger)
	}
	return passengers, nil
}
