package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Inventory-Management/db"
	"github.com/shyam81992/Inventory-Management/models"
)

//CreateInventory function
func CreateInventory(c *gin.Context) {

	var inventory models.Inventory

	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT INTO inventory (name, count, userid) 
	values ($1, $2, $3) RETURNING id`, inventory.Name, inventory.Count, inventory.UserId).Scan(&inventory.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Resource Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, inventory)
	}

}

func AddItemsToInventory(c *gin.Context) {

	var inventory models.Inventory
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var reqbody models.InventoryUpdate
	if err := c.ShouldBindJSON(&reqbody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inventory.ID = id

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `UPDATE inventory 
	SET count = count + $2
	WHERE id = $1 RETURNING name,count,userid`, inventory.ID,
		reqbody.Count).Scan(&inventory.Name, &inventory.Count, &inventory.UserId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Resource Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, inventory)
	}

}

func CreateOrders(c *gin.Context) {

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT into orders (inventoryid, userid, count) 
	values ($1, $2, $3) RETURNING id`, order.InventoryId, order.UserId, order.Count).Scan(&order.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" || err.Error() == "pq: inventory not available" {
			c.JSON(404, gin.H{
				"error": "Resource Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {

		msg, _ := json.Marshal(gin.H{
			"orderid": order.ID,
			"userid":  order.UserId,
			"status":  "order_placed",
		})
		publishmsg(msg)
		c.JSON(200, order)
	}

}

func GetSalesReport(c *gin.Context) {
	userid, _ := strconv.ParseInt(c.Param("userid"), 10, 64)
	var sales models.Sales
	sales.UserId = userid
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `SELECT inventoryid, sum(count) FROM orders
	WHERE userid=$1 GROUP BY inventoryid ORDER BY sum(count) DESC limit 100`, userid).Scan(&sales.InventoryId, &sales.Count)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(404, gin.H{
				"error": "No Records Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, sales)
	}
}
