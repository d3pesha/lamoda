package model

type Warehouse struct {
	Id        uint32     `json:"id,omitempty"`
	Name      string     `json:"name"`
	Available *bool      `json:"available"`
	Products  []*Product `json:"products,omitempty"`
}

type Product struct {
	Id         uint32              `json:"id,omitempty"`
	Name       string              `json:"name"`
	Code       string              `json:"code"`
	Quantity   uint32              `json:"quantity"`
	Size       uint16              `json:"size"`
	Warehouses []*ProductWarehouse `json:"warehouse,omitempty"`
}

type ProductWarehouse struct {
	Id          uint32 `json:"id,omitempty"`
	ProductID   uint32 `json:"product_id"`
	Quantity    uint32 `json:"quantity"`
	WarehouseID uint32 `json:"warehouse_id"`
}

type Reservation struct {
	Id          uint32 `json:"-"`
	ProductID   uint32 `json:"product_id"`
	Quantity    uint32 `json:"quantity"`
	WarehouseID uint32 `json:"warehouse_id"`
}

type ProductCreateReq struct {
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Quantity   uint32    `json:"quantity"`
	Size       uint16    `json:"size"`
	Warehouses []*uint32 `json:"warehouse,omitempty"`
}

type WarehouseCreateReq struct {
	Id        uint32 `json:"-"`
	Name      string `json:"name"`
	Available *bool  `json:"available"`
}
