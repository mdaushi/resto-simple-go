package orders

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/responses"
)

type OrderRequestBody struct {
	IdUser   int
	IdMenu   int
	Quantity int
}

func (h *handler) CreateOrder(c *gin.Context) {
	body := OrderRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// cek stok
	var stockExists models.Menu

	h.DB.First(&stockExists, body.IdMenu)

	if stockExists.Stok == 0 {
		c.JSON(http.StatusNotFound, responses.ResponsesNotFound("Stok sudah habis"))
		return
	}

	// buat orderan
	var order models.Order

	order.UserID = body.IdUser
	order.MenuID = body.IdMenu
	order.Quantity = body.Quantity

	result := h.DB.Create(&order)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// kurangi stok pesanan
	h.DB.First(&stockExists, body.IdMenu)
	stockExists.Stok = stockExists.Stok - 1
	h.DB.Save(&stockExists)

	res := responses.ResponsesSuccess("success", "")
	c.JSON(http.StatusOK, res)

}
