package routes

import (
	"encoding/json"
	"net/http"

	"github.com/elilopezm/facturas-distrital/db"
	"github.com/elilopezm/facturas-distrital/models"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Producto
	db.DB.Find(&products)

	json.NewEncoder(w).Encode(&products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Producto
	paramsp := mux.Vars(r)
	db.DB.First(&product, paramsp["id"])
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product Not Found"))
		return
	}

	db.DB.Model(&product).Association("Facturas").Find(&product.Facturas)

	json.NewEncoder(w).Encode(&product)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Producto
	json.NewDecoder(r.Body).Decode(&product)

	createdProducts := db.DB.Create(&product)
	err := createdProducts.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Producto
	paramsp := mux.Vars(r)
	db.DB.First(&product, paramsp["id"])

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {

	paramsp := mux.Vars(r) // Recuperar el ID de usuario de los parámetros de solicitud
	productId := paramsp["id"]
	var product models.Producto

	db.DB.First(&product, productId)

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product Not Found"))
		return
	}

	json.NewDecoder(r.Body).Decode(&product)

	db.DB.Save(&product)

	if db.DB.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(db.DB.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&product)
}
