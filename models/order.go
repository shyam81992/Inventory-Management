package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Inventory-Management/db"
)

// Order modal
type Order struct {
	ID          int64 `form:"id" json:"id"`
	InventoryId int64 `form:"inventoryid" json:"inventoryid" binding:"required"`
	UserId      int64 `form:"userid" json:"userid" binding:"required"`
	Count       int64 `form:"count" json:"count" binding:"required"`
}

// CreateOrderTable creates City table if not exits
func CreateOrderTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE TABLE IF NOT EXISTS orders (
		id SERIAL primary key NOT NULL,
		inventoryid Integer NOT NULL,
		userid Integer NOT NULL,
		count Integer NOT NULL,
		created_at timestamptz NOT NULL DEFAULT NOW()
	  )`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating order table")
		fmt.Println(err.Error())
		panic(err)
	}

}

// CreateOrderTrigger creates delete trigger for order table
func CreateOrderInsertTrigger() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `CREATE OR REPLACE FUNCTION process_inventory_update() RETURNS TRIGGER AS 
	$inventory_update$
		DECLARE
    		v_cnt integer;
		BEGIN
		PERFORM 1 FROM inventory WHERE id = NEW.inventoryid FOR UPDATE;
		UPDATE inventory SET count = count - NEW.count WHERE id = NEW.inventoryid AND count >= NEW.count;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		if v_cnt = 0 then
           RAISE EXCEPTION 'inventory not available'; 
        end if;
		RETURN NEW;  
		END;
	$inventory_update$ 
	LANGUAGE plpgsql;
	CREATE TRIGGER process_inventory_update AFTER INSERT ON orders
		FOR EACH ROW EXECUTE PROCEDURE process_inventory_update();`
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating order insert trigger")
		fmt.Println(err.Error())
		//panic(err)
	}
}

//Sales model

type Sales struct {
	InventoryId int64 `form:"inventoryid" json:"inventoryid" binding:"required"`
	Count       int64 `form:"count" json:"count" binding:"required"`
	UserId      int64 `form:"userid" json:"userid"`
}
