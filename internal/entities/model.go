package entities

import "time"

// Flight - сущность для таблицы Flight (билеты)
type Flight struct {
	FlightID      string    `json:"flightId"`      // Уникальный идентификатор рейса
	Departure     string    `json:"departure"`     // Пункт отправления
	Destination   string    `json:"destination"`   // Пункт назначения
	OrderNumber   string    `json:"orderNumber"`   // Номер заказа
	Provider      string    `json:"provider"`      // Поставщик (supplier)
	DepartureDate time.Time `json:"departureDate"` // Дата вылета
	ArrivalDate   time.Time `json:"arrivalDate"`   // Дата прилета
	BookingDate   time.Time `json:"bookingDate"`   // Дата оформления рейса (service_date)
	Status        string    `json:"status"`        // Статус рейса (booked, completed, cancelled)
}

// Passenger - сущность для таблицы Passenger (пассажиры)
type Passenger struct {
	PassengerID string `json:"id"`         // Уникальный идентификатор пассажира
	FlightID    string `json:"flightId"`   // ID рейса, к которому привязан пассажир
	LastName    string `json:"lastName"`   // Фамилия пассажира
	FirstName   string `json:"firstName"`  // Имя пассажира
	MiddleName  string `json:"middleName"` // Отчество (может быть пустым)
}

// Document - сущность для таблицы Document (документы пассажира)
type Document struct {
	DocumentID     string `json:"id"`             // Уникальный идентификатор документа
	PassengerID    string `json:"passengerId"`    // ID пассажира, к которому привязан документ
	DocumentType   string `json:"documentType"`   // Тип документа (паспорт, виза и т. д.)
	DocumentNumber string `json:"documentNumber"` // Номер документа
}
