package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Inventory-Management/db"
)

// Inventory modal
type Inventory struct {
	ID     int64  `form:"id" json:"id"`
	Name   string `form:"name" json:"name" binding:"required"`
	Count  int64  `form:"count" json:"count" binding:"required"`
	UserId int64  `form:"userid" json:"userid" binding:"required"`
}

// Inventory modal
type InventoryUpdate struct {
	Count int64 `form:"count" json:"count" binding:"required"`
}

// CreateInventoryTable creates City table if not exits
func CreateInventoryTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS inventory (
		id SERIAL primary key NOT NULL,
		name VARCHAR(2083) NOT NULL,
		count Integer NOT NULL,
		userid Integer NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating inventory table")
		fmt.Println(err.Error())
		panic(err)
	}

}

// CreateInventoryInsertTrigger creates delete trigger for city table
func CreateInventoryInsertUpdateTrigger() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE OR REPLACE FUNCTION process_inventory_insert_update() RETURNS TRIGGER AS 
	$inventory_insert_update$
		BEGIN
			INSERT INTO inventorylog(inventoryid, userid, count) SELECT NEW.id, NEW.userid, NEW.count;
			RETURN NEW;
		END;
	$inventory_insert_update$ 
	LANGUAGE plpgsql;
	CREATE TRIGGER process_inventory_insert_update AFTER INSERT OR UPDATE ON inventory
		FOR EACH ROW EXECUTE PROCEDURE process_inventory_insert_update();`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating inventory insert trigger")
		fmt.Println(err.Error())
		//panic(err)
	}
}
