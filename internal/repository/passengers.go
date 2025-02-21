package repository

import (
	"context"
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	"github.com/jackc/pgx/v5"
	"time"
)

type StoragePass interface {
	Create(ctx context.Context, passenger entities.Passenger) error
	GetAll(ctx context.Context) ([]entities.Passenger, error)
	GetByID(ctx context.Context, passengerID string) (entities.Passenger, error)
	GetByFlightID(ctx context.Context, flightID string) ([]entities.Passenger, error)
	Update(ctx context.Context, passenger entities.Passenger) error
	Delete(ctx context.Context, passengerID string) error
	GetReport(ctx context.Context, passengerID string, startDate, endDate time.Time) ([]entities.Flight, error)
}

type PassengerRepository struct {
	pg *pgx.Conn
}

func NewStoragePass(pg *pgx.Conn) StoragePass {
	return &PassengerRepository{pg: pg}
}

// Создание пассажира
func (s *PassengerRepository) Create(ctx context.Context, passenger entities.Passenger) error {
	_, err := s.pg.Exec(ctx, "INSERT INTO passengers (flight_id, last_name, first_name, middle_name) VALUES ($1, $2, $3, $4)",
		passenger.FlightID, passenger.LastName, passenger.FirstName, passenger.MiddleName)
	return err
}

func (s *PassengerRepository) GetAll(ctx context.Context) ([]entities.Passenger, error) {
	rows, err := s.pg.Query(context.Background(), "SELECT * FROM passengers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passengers []entities.Passenger
	for rows.Next() {
		var passenger entities.Passenger
		if err = rows.Scan(&passenger.PassengerID, &passenger.FlightID, &passenger.LastName, &passenger.FirstName, &passenger.MiddleName); err != nil {
			return nil, err
		}
		passengers = append(passengers, passenger)
	}
	return passengers, nil
}

// Получение пассажиров по ID рейса
func (s *PassengerRepository) GetByFlightID(ctx context.Context, flightID string) ([]entities.Passenger, error) {
	rows, err := s.pg.Query(ctx, "SELECT id, flight_id, last_name, first_name, middle_name FROM passengers WHERE flight_id = $1", flightID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passengers []entities.Passenger
	for rows.Next() {
		var p entities.Passenger
		if err := rows.Scan(&p.PassengerID, &p.FlightID, &p.LastName, &p.FirstName, &p.MiddleName); err != nil {
			return nil, err
		}
		passengers = append(passengers, p)
	}
	return passengers, nil
}

// Получение пассажиров по ID пассажира
func (s *PassengerRepository) GetByID(ctx context.Context, passengerID string) (entities.Passenger, error) {
	rows, err := s.pg.Query(ctx, "SELECT id, flight_id, last_name, first_name, middle_name FROM passengers WHERE id = $1", passengerID)
	if err != nil {
		return entities.Passenger{}, err
	}
	defer rows.Close()

	var p entities.Passenger
	if err := rows.Scan(&p.PassengerID, &p.FlightID, &p.LastName, &p.FirstName, &p.MiddleName); err != nil {
		return entities.Passenger{}, err
	}
	return p, nil
}

// Обновление пассажира
func (s *PassengerRepository) Update(ctx context.Context, passenger entities.Passenger) error {
	_, err := s.pg.Exec(ctx, "UPDATE passengers SET flight_id=$1, last_name=$2, first_name=$3, middle_name=$4 WHERE id=$5",
		passenger.FlightID, passenger.LastName, passenger.FirstName, passenger.MiddleName, passenger.PassengerID)
	return err
}

// Удаление пассажира
func (s *PassengerRepository) Delete(ctx context.Context, passengerID string) error {
	_, err := s.pg.Exec(ctx, "DELETE FROM passengers WHERE id = $1", passengerID)
	return err
}

func (s *PassengerRepository) GetReport(ctx context.Context, passengerID string, startDate, endDate time.Time) ([]entities.Flight, error) {
	query := `SELECT 
                f.id AS flight_id,
                f.departure_point,
                f.destination,
                f.order_number,
                f.departure_date,
                f.service_date AS booking_date,
                f.status
            FROM flights f
            JOIN passengers p ON f.id = p.flight_id
            WHERE p.id = $1
            AND (
                (f.service_date < $2 AND f.departure_date BETWEEN $2 AND $3)  
                OR 
                (f.service_date BETWEEN $2 AND $3 AND f.status = 'booked')    
                OR 
                (f.service_date BETWEEN $2 AND $3 AND f.status = 'completed')
            )
            ORDER BY f.departure_date;`

	rows, err := s.pg.Query(ctx, query, passengerID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []entities.Flight
	for rows.Next() {
		var report entities.Flight
		if err := rows.Scan(
			&report.FlightID, &report.Departure, &report.Destination,
			&report.OrderNumber, &report.DepartureDate, &report.BookingDate, &report.Status,
		); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}
