package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Inventory-Management/db"
)

// InventoryLog modal
type InventoryLog struct {
	ID          int64  `form:"id" json:"id"`
	InventoryId string `form:"inventoryid" json:"inventoryid"`
	UserId      int64  `form:"userid" json:"userid"`
	Count       string `form:"count" json:"count"`
}

// CreateInventoryLogTable creates City table if not exits
func CreateInventoryLogTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS inventorylog (
		id SERIAL primary key NOT NULL,
		inventoryid Integer NOT NULL,
		userid Integer NOT NULL,
		count Integer NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating inventorylog table")
		fmt.Println(err.Error())
		panic(err)
	}

}
