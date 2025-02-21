package handler

import (
	"github.com/SunFlowers04/SmartWayTT/internal/entities"
	repo "github.com/SunFlowers04/SmartWayTT/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DocumentHandler struct {
	storage repo.StorageDoc
}

// Конструктор обработчика документов
func NewDocumentHandler(storage repo.StorageDoc) *DocumentHandler {
	return &DocumentHandler{
		storage: storage,
	}
}

// Создать новый документ для пассажира
func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var req DocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	document := entities.Document{
		PassengerID:    req.PassengerID,
		DocumentType:   req.DocumentType,
		DocumentNumber: req.DocumentNumber,
	}

	// Сохраняем документ в хранилище
	if err := h.storage.Create(c, document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения документа"})
		return
	}

	c.JSON(http.StatusCreated, document)
}

// Получить список документов по ID пассажира
func (h *DocumentHandler) GetDocumentsByPassenger(c *gin.Context) {
	passengerID := c.Param("id") // Получаем ID пассажира из запроса

	// Получаем документы пассажира
	documents, err := h.storage.GetByPassengerID(c, passengerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Документы не найдены"})
		return
	}

	c.JSON(http.StatusOK, documents)
}

// Обновить информацию о документе
func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	var req DocumentRequest
	docID := c.Param("id") // Получаем ID документа из запроса

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	updatedDocument := entities.Document{
		DocumentID:     docID,
		PassengerID:    req.PassengerID,
		DocumentType:   req.DocumentType,
		DocumentNumber: req.DocumentNumber,
	}

	// Обновляем данные документа в хранилище
	if err := h.storage.Update(c, updatedDocument); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления документа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Документ успешно обновлен"})
}

// Удалить документ по ID
func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	docID := c.Param("id") // Получаем ID документа из запроса

	// Удаляем документ
	if err := h.storage.Delete(c, docID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления документа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Документ успешно удален"})
}
