package main

import (
	"net/http"

	"github.com/elilopezm/facturas-distrital/db"
	"github.com/elilopezm/facturas-distrital/models"
	"github.com/elilopezm/facturas-distrital/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.ConectionDB()

	db.DB.AutoMigrate(models.Cliente{})
	db.DB.AutoMigrate(models.Factura{})
	db.DB.AutoMigrate(models.Producto{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	//users
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")
	router.HandleFunc("/users/{id}", routes.PutUserHandler).Methods("PUT")

	//products
	router.HandleFunc("/products", routes.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", routes.GetProduct).Methods("GET")
	router.HandleFunc("/products", routes.PostProduct).Methods("POST")
	router.HandleFunc("/products/{id}", routes.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/{id}", routes.PutProduct).Methods("PUT")

	//bills
	router.HandleFunc("/bills", routes.GetBills).Methods("GET")
	router.HandleFunc("/bills/{id}", routes.GetBill).Methods("GET")
	router.HandleFunc("/bills", routes.PostBill).Methods("POST")
	router.HandleFunc("/bills/{id}", routes.DeleteBill).Methods("DELETE")
	router.HandleFunc("/bills/{id}", routes.PutBill).Methods("PUT")

	http.ListenAndServe(":2024", router)
}
