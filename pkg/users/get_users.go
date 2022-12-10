package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/responses"
)

func (h *handler) Index(c *gin.Context) {
	var users []models.User

	result := h.DB.Find(&users)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	res := responses.ResponsesSuccess("success", &users)
	c.JSON(http.StatusOK, res)
}
