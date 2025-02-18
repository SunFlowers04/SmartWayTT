package ports

import (
	"github.com/SunFlowers04/SmartWayTT/internal/ports/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(router *gin.Engine, flightHandler *handler.FlightHandler, passengerHandler *handler.PassengerHandler, documentHandler *handler.DocumentHandler) {

	// Common routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})

	// Routes for flights
	router.GET("/flights", flightHandler.GetFlights)    // Get all flights
	router.POST("/flights", flightHandler.CreateFlight) // Create new flight

	// Routes for passengers
	router.POST("/passengers", passengerHandler.CreatePassenger) // Create new passenger

	// Routes for documents
	router.POST("/documents", documentHandler.CreateDocument) // Create new document
}
