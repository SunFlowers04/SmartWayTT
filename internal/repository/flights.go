package repository

import (
	"context"
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	"github.com/jackc/pgx/v5"
)

type StorageFlight interface {
	Create(ctx context.Context, flight entities.Flight) error
	GetAll(ctx context.Context) ([]entities.Flight, error)
	GetByID(ctx context.Context, flightID string) (entities.Flight, error)
	Update(ctx context.Context, flight entities.Flight) error
	Delete(ctx context.Context, flightID string) error
}

type FlightRepository struct {
	pg *pgx.Conn
}

func NewStorageFlight(pg *pgx.Conn) StorageFlight {
	return &FlightRepository{pg: pg}
}

// Создание нового рейса
func (s *FlightRepository) Create(ctx context.Context, flight entities.Flight) error {
	query := `INSERT INTO flights (
		id, departure_point, destination, order_number, supplier, 
		departure_date, arrival_date, service_date, status
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.pg.Exec(ctx, query,
		flight.FlightID, flight.Departure, flight.Destination,
		flight.OrderNumber, flight.Provider, flight.DepartureDate,
		flight.ArrivalDate, flight.BookingDate, flight.Status,
	)
	return err
}

// Получение всех рейсов
func (s *FlightRepository) GetAll(ctx context.Context) ([]entities.Flight, error) {
	rows, err := s.pg.Query(ctx, `SELECT id, departure_point, destination, order_number, supplier, 
                                         departure_date, arrival_date, service_date, status 
                                  FROM flights`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []entities.Flight
	for rows.Next() {
		var flight entities.Flight
		if err := rows.Scan(
			&flight.FlightID, &flight.Departure, &flight.Destination,
			&flight.OrderNumber, &flight.Provider, &flight.DepartureDate,
			&flight.ArrivalDate, &flight.BookingDate, &flight.Status,
		); err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}
	return flights, nil
}

// Получение рейса по ID
func (s *FlightRepository) GetByID(ctx context.Context, flightID string) (entities.Flight, error) {
	var flight entities.Flight

	query := `SELECT id, departure_point, destination, order_number, supplier, 
                     departure_date, arrival_date, service_date, status 
              FROM flights WHERE id = $1`

	err := s.pg.QueryRow(ctx, query, flightID).Scan(
		&flight.FlightID, &flight.Departure, &flight.Destination,
		&flight.OrderNumber, &flight.Provider, &flight.DepartureDate,
		&flight.ArrivalDate, &flight.BookingDate, &flight.Status,
	)
	if err != nil {
		return entities.Flight{}, err
	}
	return flight, nil
}

// Обновление данных рейса
func (s *FlightRepository) Update(ctx context.Context, flight entities.Flight) error {
	query := `UPDATE flights 
	          SET departure_point = $1, destination = $2, order_number = $3, supplier = $4, 
	              departure_date = $5, arrival_date = $6, service_date = $7, status = $8 
	          WHERE id = $9`

	_, err := s.pg.Exec(ctx, query, flight.Departure, flight.Destination, flight.OrderNumber,
		flight.Provider, flight.DepartureDate, flight.ArrivalDate,
		flight.BookingDate, flight.Status, flight.FlightID)
	return err
}

// Удаление рейса по ID
func (s *FlightRepository) Delete(ctx context.Context, flightID string) error {
	query := `DELETE FROM flights WHERE id = $1`
	_, err := s.pg.Exec(ctx, query, flightID)
	return err
}
