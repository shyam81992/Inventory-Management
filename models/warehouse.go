package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Inventory-Management/db"
)

// WareHouse modal
type WareHouse struct {
	ID      int64  `form:"id" json:"id"`
	Name    string `form:"name" json:"name" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
	UserId  int64  `form:"userid" json:"userid" binding:"required"`
}

// CreateWareHouseTable creates City table if not exits
func CreateWareHouseTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS warehouse (
		id SERIAL primary key NOT NULL,
		name VARCHAR(2083) NOT NULL,
		Address text NOT NULL,
		userid Integer NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating warehouse table")
		fmt.Println(err.Error())
		panic(err)
	}

}
