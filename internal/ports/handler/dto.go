package handler

import "time"

// FlightRequest - структура запроса для создания или обновления билета
type FlightRequest struct {
	FlightID      string    `json:"flightId" binding:"required"`      // Уникальный идентификатор рейса
	Departure     string    `json:"departure" binding:"required"`     // Пункт отправления
	Destination   string    `json:"destination" binding:"required"`   // Пункт назначения
	OrderNumber   string    `json:"orderNumber" binding:"required"`   // Номер заказа
	Provider      string    `json:"provider" binding:"required"`      // Провайдер услуги
	DepartureDate time.Time `json:"departureDate" binding:"required"` // Дата вылета
	ArrivalDate   time.Time `json:"arrivalDate" binding:"required"`   // Дата прилета
	BookingDate   time.Time `json:"bookingDate" binding:"required"`   // Дата оформления
	Status        string    `json:"status" binding:"required"`        // booked, completed, cancelled
}

// FlightResponse - структура ответа с информацией о рейсе
type FlightResponse struct {
	FlightID      string    `json:"flightId"`      // Уникальный идентификатор рейса
	Departure     string    `json:"departure"`     // Пункт отправления
	Destination   string    `json:"destination"`   // Пункт назначения
	OrderNumber   string    `json:"orderNumber"`   // Номер заказа
	Provider      string    `json:"provider"`      // Провайдер услуги
	DepartureDate time.Time `json:"departureDate"` // Дата вылета
	ArrivalDate   time.Time `json:"arrivalDate"`   // Дата прилета
	BookingDate   time.Time `json:"bookingDate"`   // Дата оформления
	Status        string    `json:"status"`        // booked, completed, cancelled
}

// PassengerRequest - структура запроса для создания или обновления пассажира
type PassengerRequest struct {
	ID         string `json:"id" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	FirstName  string `json:"firstName" binding:"required"`
	MiddleName string `json:"middleName"`
}

// PassengerResponse - структура ответа для пассажира
type PassengerResponse struct {
	PassengerID string `json:"passengerId"` // ID пассажира
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
}

// DocumentRequest - структура запроса для документа пассажира
type DocumentRequest struct {
	PassengerID    string `json:"passengerId" binding:"required"`  // ID пассажира
	DocumentType   string `json:"documentType" binding:"required"` // Тип документа (паспорт, виза и т. д.)
	DocumentNumber string `json:"documentNumber" binding:"required"`
}

// DocumentResponse - структура ответа для документа пассажира
type DocumentResponse struct {
	DocumentID     string `json:"documentId"`     // ID документа
	PassengerID    string `json:"passengerId"`    // ID пассажира
	DocumentType   string `json:"documentType"`   // Тип документа
	DocumentNumber string `json:"documentNumber"` // Номер документа
}

// FlightDetailsResponse - структура ответа для полной информации о билете
type FlightDetailsResponse struct {
	Flight     FlightResponse      `json:"flight"`     // Информация о рейсе
	Passengers []PassengerWithDocs `json:"passengers"` // Пассажиры с документами
}

// PassengerWithDocs - структура для представления пассажира с его документами
type PassengerWithDocs struct {
	Passenger PassengerResponse  `json:"passenger"`
	Documents []DocumentResponse `json:"documents"`
}

// PassengerReportRequest - запрос на получение отчета по пассажиру
type PassengerReportRequest struct {
	PassengerID string    `json:"passengerId" binding:"required"` // ID пассажира
	StartDate   time.Time `json:"startDate" binding:"required"`   // Начало периода
	EndDate     time.Time `json:"endDate" binding:"required"`     // Конец периода
}

// PassengerReportResponse - ответ с отчетом по пассажиру за период
type PassengerReportResponse struct {
	PassengerID string         `json:"passengerId"` // ID пассажира
	FirstName   string         `json:"firstName"`   // Имя
	LastName    string         `json:"lastName"`    // Фамилия
	MiddleName  string         `json:"middleName"`  // Отчество (если есть)
	Flights     []ReportFlight `json:"flights"`     // Список перелетов
}

// ReportFlight - структура для представления перелетов в отчете
type ReportFlight struct {
	BookingDate     time.Time `json:"bookingDate"`     // Дата оформления
	DepartureDate   time.Time `json:"departureDate"`   // Дата вылета
	OrderNumber     string    `json:"orderNumber"`     // Номер заказа
	Departure       string    `json:"departure"`       // Пункт отправления
	Destination     string    `json:"destination"`     // Пункт назначения
	ServiceProvided bool      `json:"serviceProvided"` // Услуга оказана (true/false)
}
