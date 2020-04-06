package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/shyam81992/Inventory-Management/db"
	"github.com/shyam81992/Inventory-Management/models"
)

//CreateUser function
func CreateUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var id int64
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	err := db.Db.QueryRowContext(ctx, `INSERT INTO users(name, role, email) 
	SELECT CAST($1 AS VARCHAR), CAST($2 AS user_type), CAST($3 AS VARCHAR) WHERE NOT EXISTS (
        SELECT 1 FROM users WHERE email=$3
    ) RETURNING id`, user.Name, user.Role, user.Email).Scan(&user.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(403, gin.H{
				"error": "Resource already exits",
			})
			return
		}
		if pqError, ok := err.(*pq.Error); ok {
			if pqError.Code == "23505" {
				c.JSON(403, gin.H{
					"error": "Resource already exits",
				})
			} else {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
		} else {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

	} else {
		c.JSON(200, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"role":  user.Role,
			"email": user.Email,
		})
	}

}

//CreateAddress function
func CreateAddress(c *gin.Context) {

	var addresss models.Address
	if err := c.ShouldBindJSON(&addresss); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `INSERT INTO address(userid, address) 
	SELECT $1, $2 WHERE EXISTS (
        SELECT 1 FROM user WHERE id=$1
    ) RETURNING id`
	err := db.Db.QueryRowContext(ctx, sqlStatement, addresss.UserId, addresss.Address).Scan(&addresss.ID)
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

		c.JSON(200, addresss)
	}

}
