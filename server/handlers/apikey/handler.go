package apikey

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"notification-service/internal/repositories/apikey_repository"
	"github.com/google/uuid"
)

type Handler struct {
	apikeyRepo apikey_repository.APIKeyRepository
}

func NewHandler(apikeyRepo apikey_repository.APIKeyRepository) *Handler {
	return &Handler{
		apikeyRepo: apikeyRepo,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Generate a new API key
	key := uuid.New().String()

	apiKey, err := h.apikeyRepo.CreateAPIKey(c.Request.Context(), userID, key, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create API key"})
		return
	}

	c.JSON(http.StatusCreated, apiKey)
}

func (h *Handler) Delete(c *gin.Context) {
	keyID := c.Param("id")
	userID := c.GetString("user_id")

	if err := h.apikeyRepo.DeactivateAPIKey(c.Request.Context(), keyID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete API key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "API key deleted successfully"})
}