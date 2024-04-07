package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=eslopezm password=12345 dbname=facturacion_distrital port=5432"
var DB *gorm.DB //metodo par hacer migraciones

func ConectionDB() { //uso de gorm para la conexion
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{}) //AQUI SE INTENTA CONECTAR Y MANDA LA BD O UN ERR
	if error != nil {
		log.Fatal(error) // puede usarse un panic
	} else {
		log.Println("DB CONECTADA....")
	}
}
