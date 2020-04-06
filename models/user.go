package models

import (
	"context"
	"fmt"
	"time"

	"github.com/shyam81992/Inventory-Management/db"
)

// User modal
type User struct {
	ID    int64  `form:"id" json:"id"`
	Name  string `form:"name" json:"name" binding:"required"`
	Role  string `form:"role" json:"role" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
}

// CreateUserTable creates City table if not exits
func CreateUserTable() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	sqlStatement := `
	DO $$
	BEGIN
	IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'user_type') THEN
	CREATE TYPE   user_type AS ENUM ('vendor', 'buyer', 'shipper');
	END IF;
	END$$;	
	CREATE TABLE IF NOT EXISTS users (
			id SERIAL primary key NOT NULL,
			name VARCHAR(2083) NOT NULL,
			role user_type NOT NULL,
			email VARCHAR(2083) NOT NULL,
			created_at timestamptz NOT NULL DEFAULT NOW(),
			UNIQUE(email)
		  )  `
	_, err := db.Db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("error in creating user table")
		fmt.Println(err.Error())
		panic(err)
	}

}
