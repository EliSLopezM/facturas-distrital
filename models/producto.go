package models

import "gorm.io/gorm"

type Producto struct {
	gorm.Model

	ID             uint      `json:"id"`
	Nombre         string    `gorm:"type:varchar(100);not null" json:"nombre"`
	Facturas       []Factura `gorm:"many2many:facturas_productos;"`
	PrecioUnitario float64   `gorm:"not null" json:"precio_unitario"`
	Cantidad       int       `gorm:"not null" json:"cantidad"`
	Descripción    string    `json:"descripcion"`
}
