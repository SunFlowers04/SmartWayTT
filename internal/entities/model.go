package entities

import "time"

// Flight - сущность для таблицы Flight (билеты)
type Flight struct {
	ID            string    `gorm:"primaryKey;type:varchar(255)" json:"flightId"`
	Departure     string    `gorm:"type:varchar(255)" json:"departure"`
	Destination   string    `gorm:"type:varchar(255)" json:"destination"`
	OrderNumber   string    `gorm:"type:varchar(255)" json:"orderNumber"`
	Provider      string    `gorm:"type:varchar(255)" json:"provider"`
	DepartureDate time.Time `gorm:"type:timestamp" json:"departureDate"`
	ArrivalDate   time.Time `gorm:"type:timestamp" json:"arrivalDate"`
	BookingDate   time.Time `gorm:"type:timestamp" json:"bookingDate"`
}

// Passenger - сущность для таблицы Passenger (пассажиры)
type Passenger struct {
	ID         string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	FlightID   string `gorm:"type:varchar(255);not null" json:"flightId"`
	LastName   string `gorm:"type:varchar(255);not null" json:"lastName"`
	FirstName  string `gorm:"type:varchar(255);not null" json:"firstName"`
	MiddleName string `gorm:"type:varchar(255)" json:"middleName"`
	Flight     Flight `gorm:"foreignKey:FlightID" json:"flight"`
}

// Document - сущность для таблицы Document (документы пассажира)
type Document struct {
	ID             string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	PassengerID    string    `gorm:"type:varchar(255);not null" json:"passengerId"`
	DocumentType   string    `gorm:"type:varchar(255);not null" json:"documentType"`
	DocumentNumber string    `gorm:"type:varchar(255);not null" json:"documentNumber"`
	Passenger      Passenger `gorm:"foreignKey:PassengerID" json:"passenger"`
}
