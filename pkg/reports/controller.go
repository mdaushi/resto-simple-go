package reports

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{DB: db}

	routes := r.Group("reports")
	routes.GET("/monthly", h.Monthly)
	routes.GET("/yearly", h.Yearly)
}
