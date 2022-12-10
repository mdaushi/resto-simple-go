package menus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/responses"
)

func (h *handler) Index(c *gin.Context) {
	var menus []models.Menu

	result := h.DB.Find(&menus)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, responses.ResponsesNotFound("Menus not found"))
		return
	}

	res := responses.ResponsesSuccess("success", &menus)
	c.JSON(http.StatusOK, res)
}

type StockSelect struct {
	Nama string
	Stok int
}

func (h *handler) Stock(c *gin.Context) {
	var menus []StockSelect

	result := h.DB.Raw("SELECT nama, stok FROM menus").Scan(&menus)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, responses.ResponsesNotFound("Menus not found"))
		return
	}

	res := responses.ResponsesSuccess("Stok Menu", &menus)
	c.JSON(http.StatusOK, res)

}
