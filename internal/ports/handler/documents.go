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

func NewDocumentHandler(storage repo.StorageDoc) *DocumentHandler {
	return &DocumentHandler{
		storage: storage,
	}
}

func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var req DocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusPaymentRequired, err)
		return
	}

	document := entities.Document{
		DocumentType: req.DocumentType,
		// и другие поля
	}

	if err := h.storage.Create(document); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusCreated, document)
}

func (h *DocumentHandler) GetDocuments(c *gin.Context) {
	documents, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, documents)
}
