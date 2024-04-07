package routes

import (
	"encoding/json"
	"net/http"

	"github.com/elilopezm/facturas-distrital/db"
	"github.com/elilopezm/facturas-distrital/models"
	"github.com/gorilla/mux"
)

func GetBills(w http.ResponseWriter, r *http.Request) {
	var bills []models.Factura
	db.DB.Find(&bills)

	json.NewEncoder(w).Encode(&bills)
}

func GetBill(w http.ResponseWriter, r *http.Request) {
	var bill models.Factura
	paramsb := mux.Vars(r)
	db.DB.First(&bill, paramsb["id"])
	if bill.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bill Not Found"))
		return
	}

	db.DB.Model(&bill).Association("Productos").Find(&bill.Productos)
	json.NewEncoder(w).Encode(&bill)
}

func PostBill(w http.ResponseWriter, r *http.Request) {
	var bill models.Factura
	json.NewDecoder(r.Body).Decode(&bill)

	createdUser := db.DB.Create(&bill)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&bill)
}

func DeleteBill(w http.ResponseWriter, r *http.Request) {
	var bill models.Factura
	paramsb := mux.Vars(r)
	db.DB.First(&bill, paramsb["id"])

	if bill.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bill Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&bill)
	w.WriteHeader(http.StatusOK)
}

func PutBill(w http.ResponseWriter, r *http.Request) {

	paramsb := mux.Vars(r) // Recuperar el ID de usuario de los parámetros de solicitud
	billId := paramsb["id"]
	var bill models.Cliente

	db.DB.First(&bill, billId)

	if bill.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bill Not Found"))
		return
	}

	json.NewDecoder(r.Body).Decode(&bill)

	db.DB.Save(&bill)

	if db.DB.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(db.DB.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&bill)
}
