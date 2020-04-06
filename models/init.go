package models

func Init() {
	//user
	CreateUserTable()
	CreateAddressTable()

	// Inventory
	CreateInventoryTable()
	CreateInventoryInsertUpdateTrigger()
	CreateInventoryLogTable()

	CreateOrderTable()
	CreateOrderInsertTrigger()
	CreateWareHouseTable()

	CreateShipmentTable()
}
