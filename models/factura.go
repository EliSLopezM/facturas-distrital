package models

import (
	"time"

	"gorm.io/gorm"
)

type Factura struct {
	gorm.Model

	ID        uint       `json:"id"`
	Fecha     time.Time  `json:"fecha"`
	Productos []Producto `gorm:"many2many:facturas_productos;"`
	Total     float64    `gorm:"not null" json:"total"`
	ClienteID int        `json:"cliente_id"` // Clave externa que hace referencia a la tabla "clientes"gorm:"foreignKey:ID;references:clientes.id"
}

/* Calcular el total al crear o actualizar una factura
func (f *Factura) BeforeSave(tx *gorm.DB) error {
	for _, producto := range f.Productos {
		f.Total += producto.PrecioUnitario * float64(producto.Cantidad)
	}
	return nil
}*/

// Función para obtener el total de la factura
func (f *Factura) GetTotal() float64 {
	var total float64
	for _, producto := range f.Productos {
		total += producto.PrecioUnitario * float64(producto.Cantidad)
	}
	return total
}
