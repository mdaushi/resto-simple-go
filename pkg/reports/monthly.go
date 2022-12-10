package reports

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/responses"
)

type resultMonth struct {
	Date  time.Time
	Total int64
}

func (h *handler) Monthly(c *gin.Context) {
	var orders []resultMonth

	result := h.DB.Model(&models.Order{}).Select("extract(month from orders.created_at) as month, sum(orders.quantity * menus.harga) as total").Joins("left join menus on menus.id = orders.menu_id").Group("month").Scan(&orders)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	res := responses.ResponsesSuccess("success", &orders)
	c.JSON(http.StatusOK, res)
}
