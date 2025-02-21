package ports

import (
	"github.com/SunFlowers04/SmartWayTT/internal/ports/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Функция Router регистрирует маршруты API
func Router(router *gin.Engine, flightHandler *handler.FlightHandler, passengerHandler *handler.PassengerHandler, documentHandler *handler.DocumentHandler) {
	// Обработчик для несуществующих маршрутов (404)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Страница не найдена"})
	})

	// Маршруты для билетов (flights)
	router.GET("/flights", flightHandler.GetFlights)           // Получить список билетов
	router.GET("/flights/:id", flightHandler.GetFlightDetails) // Получить полную информацию по билету
	router.POST("/flights", flightHandler.CreateFlight)        // Создать новый билет
	router.PUT("/flights/:id", flightHandler.UpdateFlight)     // Обновить информацию о билете
	router.DELETE("/flights/:id", flightHandler.DeleteFlight)  // Удалить билет

	// Маршруты для пассажиров (passengers)
	router.GET("/flights/:id/passengers", passengerHandler.GetPassengersByFlight) // Получить список пассажиров по билету
	router.POST("/passengers", passengerHandler.CreatePassenger)                  // Создать пассажира
	router.PUT("/passengers/:id", passengerHandler.UpdatePassenger)               // Обновить информацию о пассажире
	router.DELETE("/passengers/:id", passengerHandler.DeletePassenger)            // Удалить пассажира
	router.GET("/passengers/:id/report", passengerHandler.GetPassengerReport)     // Получить отчет по пассажиру за период

	// Маршруты для документов (documents)
	router.GET("/passengers/:id/documents", documentHandler.GetDocumentsByPassenger) // Получить список документов по пассажиру
	router.POST("/documents", documentHandler.CreateDocument)                        // Создать документ
	router.PUT("/documents/:id", documentHandler.UpdateDocument)                     // Обновить информацию о документе
	router.DELETE("/documents/:id", documentHandler.DeleteDocument)                  // Удалить документ
}
