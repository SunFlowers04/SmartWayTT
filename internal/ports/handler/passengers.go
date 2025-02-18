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

func (h *PassengerHandler) CreatePassenger(c *gin.Context) {
	var req PassengerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	passenger := entities.Passenger{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		// и другие поля
	}

	if err := h.storage.Create(passenger); err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusCreated, passenger)
}

func (h *PassengerHandler) GetPassengers(c *gin.Context) {
	passengers, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, passengers)
		return
	}
	c.JSON(http.StatusOK, passengers)
}
