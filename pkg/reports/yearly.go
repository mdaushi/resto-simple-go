package reports

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/responses"
)

type resultYear struct {
	Date  time.Time
	Total int64
}

func (h *handler) Yearly(c *gin.Context) {
	var orders []resultYear

	result := h.DB.Model(&models.Order{}).Select("extract(year from orders.created_at) as year, sum(orders.quantity * menus.harga) as total").Joins("left join menus on menus.id = orders.menu_id").Group("year").Scan(&orders)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	res := responses.ResponsesSuccess("success", &orders)
	c.JSON(http.StatusOK, res)
}
