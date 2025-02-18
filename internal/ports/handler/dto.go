package handler

import "time"

// FlightRequest - структура запроса для получения информации о рейсе
type FlightRequest struct {
	FlightID string `json:"flightId" binding:"required"`
}

// FlightResponse - структура ответа для информации о рейсе
type FlightResponse struct {
	FlightID      string    `json:"flightId"`
	Departure     string    `json:"departure"`
	Destination   string    `json:"destination"`
	OrderNumber   string    `json:"orderNumber"`
	Provider      string    `json:"provider"`
	DepartureDate time.Time `json:"departureDate"`
	ArrivalDate   time.Time `json:"arrivalDate"`
	BookingDate   time.Time `json:"bookingDate"`
}

// PassengerRequest - структура запроса для создания или обновления пассажира
type PassengerRequest struct {
	LastName   string `json:"lastName" binding:"required"`
	FirstName  string `json:"firstName" binding:"required"`
	MiddleName string `json:"middleName"`
}

// PassengerResponse - структура ответа для пассажира
type PassengerResponse struct {
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

// DocumentRequest - структура запроса для документа пассажира
type DocumentRequest struct {
	DocumentType   string `json:"documentType" binding:"required"`
	DocumentNumber string `json:"documentNumber" binding:"required"`
}

// DocumentResponse - структура ответа для документа пассажира
type DocumentResponse struct {
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
}

// FlightDetailsResponse - структура ответа для полной информации о билете
type FlightDetailsResponse struct {
	Flight     FlightResponse      `json:"flight"`
	Passengers []PassengerResponse `json:"passengers"`
	Documents  []DocumentResponse  `json:"documents"`
}
