package repository

import (
	"context"
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	"github.com/jackc/pgx/v5"
)

type StorageFlight interface {
	Create(flight entities.Flight) error
	GetAll() ([]entities.Flight, error)
}

type storageFlight struct {
	pg *pgx.Conn
}

func NewStorageFlight(pg *pgx.Conn) StorageFlight {
	return &storageFlight{pg: pg}
}

func (s *storageFlight) Create(flight entities.Flight) error {
	// Write to DB
	_, err := s.pg.Exec(context.Background(),
		`INSERT INTO flights (order_number, departure_point) 
     VALUES ($1, $2)`,
		flight.OrderNumber,
		flight.Destination,
		// добавьте остальные поля
	)
	return err
}

func (s *storageFlight) GetAll() ([]entities.Flight, error) {
	rows, err := s.pg.Query(context.Background(), "SELECT * FROM flights")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []entities.Flight
	for rows.Next() {
		var flight entities.Flight
		if err := rows.Scan(&flight.OrderNumber, &flight.Destination); err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}
	return flights, nil
}
