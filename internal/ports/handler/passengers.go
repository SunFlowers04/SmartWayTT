package handler

import (
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	repo "github.com/SunFlowers04/SmartWayTT/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PassengerHandler struct {
	storage repo.StoragePass
}

func NewPassengerHandler(storage repo.StoragePass) *PassengerHandler {
	return &PassengerHandler{
		storage: storage,
	}
}

// Получить список всех пассажиров
func (h *PassengerHandler) GetPassengers(c *gin.Context) {
	ctx := c.Request.Context()
	passengers, err := h.storage.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пассажиры не найдены"})
		return
	}
	c.JSON(http.StatusOK, passengers)
}

// Получить список пассажиров по ID рейса
func (h *PassengerHandler) GetPassengersByFlight(c *gin.Context) {
	ctx := c.Request.Context()
	flightID := c.Param("id")

	passengers, err := h.storage.GetByFlightID(ctx, flightID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пассажиры не найдены"})
		return
	}

	c.JSON(http.StatusOK, passengers)
}

// Создать нового пассажира
func (h *PassengerHandler) CreatePassenger(c *gin.Context) {
	ctx := c.Request.Context()
	var req PassengerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	passenger := entities.Passenger{
		FlightID:   req.ID,
		LastName:   req.LastName,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
	}

	if err := h.storage.Create(ctx, passenger); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пассажира"})
		return
	}

	c.JSON(http.StatusCreated, passenger)
}

// Обновить информацию о пассажире
func (h *PassengerHandler) UpdatePassenger(c *gin.Context) {
	ctx := c.Request.Context()
	passengerID := c.Param("id")

	var req PassengerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	updatedPassenger := entities.Passenger{
		PassengerID: passengerID,
		FlightID:    req.ID,
		LastName:    req.LastName,
		FirstName:   req.FirstName,
		MiddleName:  req.MiddleName,
	}

	if err := h.storage.Update(ctx, updatedPassenger); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления пассажира"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пассажир успешно обновлен"})
}

// Удалить пассажира
func (h *PassengerHandler) DeletePassenger(c *gin.Context) {
	ctx := c.Request.Context()
	passengerID := c.Param("id")

	if err := h.storage.Delete(ctx, passengerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления пассажира"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пассажир успешно удален"})
}

// GetPassengerReport - обработчик запроса отчета по пассажиру за период
func (h *PassengerHandler) GetPassengerReport(c *gin.Context) {
	ctx := c.Request.Context()
	// Привязываем JSON-запрос к структуре DTO
	var req PassengerReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	passenger, err := h.storage.GetByID(c, req.PassengerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения пассажира "})
		return
	}

	// Вызов репозитория для получения данных о перелетах пассажира
	flights, err := h.storage.GetReport(ctx, req.PassengerID, req.StartDate, req.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения отчета"})
		return
	}

	// Если перелетов нет, возвращаем пустой массив
	if len(flights) == 0 {
		c.JSON(http.StatusOK, PassengerReportResponse{
			PassengerID: req.PassengerID,
			Flights:     []ReportFlight{},
		})
		return
	}

	// Формируем ответ
	var reportFlights []ReportFlight
	for _, flight := range flights {
		reportFlights = append(reportFlights, ReportFlight{
			BookingDate:     flight.BookingDate,
			DepartureDate:   flight.DepartureDate,
			OrderNumber:     flight.OrderNumber,
			Departure:       flight.Departure,
			Destination:     flight.Destination,
			ServiceProvided: flight.Status == "completed", // Если рейс завершен, услуга оказана
		})
	}

	// Формируем полный отчет по пассажиру
	response := PassengerReportResponse{
		PassengerID: req.PassengerID,
		FirstName:   passenger.FirstName,
		LastName:    passenger.LastName,
		MiddleName:  passenger.MiddleName,
		Flights:     reportFlights,
	}

	// Возвращаем JSON-ответ
	c.JSON(http.StatusOK, response)
}
