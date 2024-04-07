package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model

	Cedula         string    `gorm:"not null;unique_index" json:"cedula"`
	PrimerNombre   string    `gorm:"type:varchar(100);not null" json:"primer_nombre"`
	PrimerApellido string    `gorm:"type:varchar(100);not null" json:"primer_apellido"`
	Dirección      string    `gorm:"type:varchar(100);" json:"direccion"`
	Celular        string    `json:"celular"`
	Email          string    `gorm:"type:varchar(100);not null" json:"email"`
	Facturas       []Factura `json:"facturas"`
}
