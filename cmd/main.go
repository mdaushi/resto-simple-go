package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdaushi/resto-simple-go/pkg/common/db"
	"github.com/mdaushi/resto-simple-go/pkg/menus"
	"github.com/mdaushi/resto-simple-go/pkg/orders"
	"github.com/mdaushi/resto-simple-go/pkg/receipts"
	"github.com/mdaushi/resto-simple-go/pkg/reports"
	"github.com/mdaushi/resto-simple-go/pkg/users"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	// daftar routes disini
	orders.RegisterRoutes(r, h)
	menus.RegisterRoutes(r, h)
	users.RegisterRoutes(r, h)
	receipts.RegisterRoutes(r, h)
	reports.RegisterRoutes(r, h)

	r.Run(port)
}
