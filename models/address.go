package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Inventory-Management/db"
)

// Address modal
type Address struct {
	ID      int64  `form:"id" json:"id"`
	UserId  int64  `form:"userid" json:"userid" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
}

// CreateAddressTable creates City table if not exits
func CreateAddressTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS address (
		id SERIAL primary key NOT NULL,
		userid Integer NOT NULL,
		Address text NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating warehouse table")
		fmt.Println(err.Error())
		panic(err)
	}

}
