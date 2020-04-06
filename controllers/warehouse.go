package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Inventory-Management/config"
	"github.com/shyam81992/Inventory-Management/db"
	"github.com/shyam81992/Inventory-Management/models"
	"github.com/streadway/amqp"
)

func CreateWareHouse(c *gin.Context) {

	var wh models.WareHouse
	if err := c.ShouldBindJSON(&wh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT INTO warehouse(name, address, userid) values($1, $2, $3) 
	RETURNING id`, wh.Name, wh.Address, wh.UserId).Scan(&wh.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Record Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {

		c.JSON(200, wh)
	}

}

//CreateShipment function
func CreateShipment(c *gin.Context) {

	var shipment models.Shipment
	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT INTO shipment(orderid, userid, warehouseid, status)
	 values($1, $2, $3, $4) RETURNING id`, shipment.OrderId, shipment.UsersId,
		shipment.WarehouseId, shipment.Status).Scan(&shipment.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Record Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		msg, _ := json.Marshal(shipment)
		err = publishmsg(msg)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, shipment)
		}

	}

}

func publishmsg(msg []byte) error {
	conn, err := amqp.Dial(config.RabbitConfig["uri"])
	if err != nil {
		fmt.Println(err, "Failed to connect to RabbitMQ")
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err, "Failed to open a channel")
		return err
	}
	defer ch.Close()
	err = ch.Publish(
		"",                               // exchange
		config.RabbitConfig["queuename"], // routing key
		false,                            // mandatory
		false,                            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})

	if err != nil {
		fmt.Println(err, "Failed to publish a message")
	}
	return err
}
