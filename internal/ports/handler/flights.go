package handler

import (
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	repo "github.com/SunFlowers04/SmartWayTT/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FlightHandler struct {
	storage repo.StorageFlight
}

func NewFlightHandler(storage repo.StorageFlight) *FlightHandler {
	return &FlightHandler{
		storage: storage,
	}
}

// Получить список всех билетов
func (h *FlightHandler) GetFlights(c *gin.Context) {
	flights, err := h.storage.GetAll(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Билеты не найдены"})
		return
	}
	c.JSON(http.StatusOK, flights)
}

// Получить полную информацию по билету
func (h *FlightHandler) GetFlightDetails(c *gin.Context) {
	flightID := c.Param("id")
	flight, err := h.storage.GetByID(c, flightID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Билет не найден"})
		return
	}
	c.JSON(http.StatusOK, flight)
}

func (h *FlightHandler) CreateFlight(c *gin.Context) {
	var req FlightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err) // парсинг джейсон клиента в стркутуру сервиса
		return
	}

	flight := entities.Flight{
		FlightID:      req.FlightID,
		Departure:     req.Departure,
		Destination:   req.Destination,
		OrderNumber:   req.OrderNumber,
		Provider:      req.Provider,
		DepartureDate: req.DepartureDate,
		ArrivalDate:   req.ArrivalDate,
		BookingDate:   req.BookingDate,
		Status:        req.Status,
		// и другие поля
	}

	if err := h.storage.Create(c, flight); err != nil {
		c.JSON(http.StatusInternalServerError, flight)
		return
	}

	c.JSON(http.StatusCreated, flight)
}

// Обновить информацию о билете
func (h *FlightHandler) UpdateFlight(c *gin.Context) {
	flightID := c.Param("id")
	var req FlightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	updatedFlight := entities.Flight{
		FlightID:      flightID,
		Departure:     req.Departure,
		Destination:   req.Destination,
		OrderNumber:   req.OrderNumber,
		Provider:      req.Provider,
		DepartureDate: req.DepartureDate,
		ArrivalDate:   req.ArrivalDate,
		BookingDate:   req.BookingDate,
		Status:        req.Status,
	}

	if err := h.storage.Update(c, updatedFlight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления билета"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Билет успешно обновлен"})
}

// Удалить билет
func (h *FlightHandler) DeleteFlight(c *gin.Context) {
	flightID := c.Param("id")

	if err := h.storage.Delete(c, flightID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления билета"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Билет успешно удален"})
}
