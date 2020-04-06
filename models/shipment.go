package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Inventory-Management/db"
)

// Shipment modal
type Shipment struct {
	ID          int64  `form:"id" json:"id"`
	OrderId     int64  `form:"orderid" json:"orderid" binding:"required"`
	UsersId     int64  `form:"userid" json:"userid" binding:"required"`
	WarehouseId int64  `form:"warehouseid" json:"warehouseid" binding:"required"`
	Status      string `form:"status" json:"status" binding:"required"`
}

// CreateShipmentTable creates City table if not exits
func CreateShipmentTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `
	DO $$
	BEGIN
	IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'shipmentstatus_type') THEN
	CREATE TYPE   shipmentstatus_type AS ENUM ('order_placed', 'shipping', 'out_for_delivery');
	END IF;
	END$$;	
	CREATE TABLE IF NOT EXISTS shipment (
		id SERIAL primary key NOT NULL,
		orderid Integer NOT NULL,
		userid Integer NOT NULL,
		warehouseid Integer NOT NULL,
		status shipmentstatus_type NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating shipment table")
		fmt.Println(err.Error())
		panic(err)
	}

}
