package receipts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/responses"
)

func (h *handler) Print(c *gin.Context) {
	id := c.Param("id")

	// cek apakah sudah terprint atau belum
	var receiptExists models.Receipt
	receiptReady := h.DB.Preload("Order.Menu").Preload("Order.User").First(&receiptExists, "order_id = ?", id)

	// jika sudah terprint langsung tampilkan
	if receiptReady.Error == nil {
		c.JSON(http.StatusOK, responses.ResponsesSuccess("success", formatReceipt(receiptExists)))
		return
	}
	// cek orderan apakah id ada atau tidak
	var order models.Order

	result := h.DB.First(&order, id)

	if result.Error != nil {
		c.JSON(http.StatusOK, responses.ResponsesNotFound("ID Not Found"))
		return
	}

	// print receipt(store ke db)
	receiptStore := models.Receipt{
		OrderID: int(order.ID),
	}

	receiptResult := h.DB.Create(&receiptStore)
	if receiptResult.Error != nil {
		c.AbortWithError(http.StatusNotFound, receiptResult.Error)
		return
	}

	h.DB.Preload("Order.Menu").Preload("Order.User").First(&receiptExists, "order_id = ?", id)

	res := responses.ResponsesSuccess("success", formatReceipt(receiptExists))
	c.JSON(http.StatusOK, res)
}

type strukFormat struct {
	Title      string      `json:"title"`
	TotalItems int         `json:"total"`
	Pesanan    interface{} `json:"pesanan"`
	TotalHarga int         `json:"total_harga"`
	Tanggal    time.Time   `json:"tanggal"`
}

func formatReceipt(data models.Receipt) strukFormat {
	s := strukFormat{}
	s.Title = "Struk Pesanan"
	s.TotalItems = data.Order.Quantity
	s.Pesanan = data
	s.TotalHarga = data.Order.Quantity * int(data.Order.Menu.Harga)
	s.Tanggal = time.Now()

	return s
}
