package model

type Warehouse struct {
	ID        uint32     `jsonapi:"primary,warehouse"`
	Name      string     `jsonapi:"attr,name"`
	Available *bool      `jsonapi:"attr,available"`
	Products  []*Product `jsonapi:"relation,product,omitempty"`
}

type Product struct {
	ID       uint32 `jsonapi:"primary,product"`
	Name     string `jsonapi:"attr,name"`
	Code     string `jsonapi:"attr,code"`
	Quantity uint32 `jsonapi:"attr,quantity"`
	Size     uint16 `jsonapi:"attr,size"`
}

type ProductWarehouse struct {
	ID          uint32 `jsonapi:"primary,product_warehouse"`
	ProductID   uint32 `jsonapi:"attr,product_id"`
	Quantity    uint32 `jsonapi:"attr,quantity"`
	WarehouseID uint32 `jsonapi:"attr,warehouse_id"`
}

type Reservation struct {
	ID          uint32 `jsonapi:"primary,reservation"`
	ProductID   uint32 `jsonapi:"attr,product_id"`
	Quantity    uint32 `jsonapi:"attr,quantity"`
	WarehouseID uint32 `jsonapi:"attr,warehouse_id"`
}

type ReservationReq struct {
	Data []*Reservation `jsonapi:"relation,reservation"`
}

type ProductCreateReq struct {
	ID         uint32       `jsonapi:"primary,product"`
	Name       string       `jsonapi:"attr,name"`
	Code       string       `jsonapi:"attr,code"`
	Quantity   uint32       `jsonapi:"attr,quantity"`
	Size       uint16       `jsonapi:"attr,size"`
	Warehouses []*Warehouse `jsonapi:"relation,warehouse,omitempty"`
}

type WarehouseCreateReq struct {
	ID        uint32 `jsonapi:"primary,warehouse"`
	Name      string `jsonapi:"attr,name"`
	Available *bool  `jsonapi:"attr,available"`
}
