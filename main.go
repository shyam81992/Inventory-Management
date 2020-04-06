package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Inventory-Management/config"
	"github.com/shyam81992/Inventory-Management/controllers"
	"github.com/shyam81992/Inventory-Management/db"
	"github.com/shyam81992/Inventory-Management/models"
)

func main() {

	config.LoadConfig()
	db.InitDb()
	models.Init()

	r := gin.Default()

	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", controllers.CreateUser)
	r.POST("/address", controllers.CreateAddress)

	r.POST("/inventory", controllers.CreateInventory)
	r.PATCH("/inventory/:id", controllers.AddItemsToInventory)
	r.POST("/orders", controllers.CreateOrders)

	r.POST("/warehouse", controllers.CreateWareHouse)
	r.POST("/shipment", controllers.CreateShipment)

	r.GET("/sales/:userid", controllers.GetSalesReport)

	r.Run(":" + config.AppConfig["port"])
}
