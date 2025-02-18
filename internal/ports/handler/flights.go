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

func (h *FlightHandler) CreateFlight(c *gin.Context) {
	var req FlightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusCreated, err)
		return
	}

	flight := entities.Flight{
		OrderNumber: req.FlightID,
		// и другие поля
	}

	if err := h.storage.Create(flight); err != nil {
		c.JSON(http.StatusCreated, flight)
		return
	}

	c.JSON(http.StatusCreated, flight)
}

func (h *FlightHandler) GetFlights(c *gin.Context) {
	flights, err := h.storage.GetAll()
	if err != nil {

		c.JSON(http.StatusCreated, err)
		return
	}
	c.JSON(http.StatusOK, flights)
}
