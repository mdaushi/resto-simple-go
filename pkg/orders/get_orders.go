package orders

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/responses"
)

func (h *handler) Index(c *gin.Context) {
	var orders []models.Order

	result := h.DB.Preload("User").Preload("Menu").Find(&orders)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	res := responses.ResponsesSuccess("success", &orders)
	c.JSON(http.StatusOK, res)
}
