package routes

import (
	"encoding/json"
	"net/http"

	"github.com/elilopezm/facturas-distrital/db"
	"github.com/elilopezm/facturas-distrital/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Cliente
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Cliente
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	db.DB.Model(&user).Association("Facturas").Find(&user.Facturas)
	json.NewEncoder(w).Encode(&user)
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Cliente
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Cliente
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) // Recuperar el ID de usuario de los parámetros de solicitud
	userId := params["id"]
	var user models.Cliente

	db.DB.First(&user, userId)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	json.NewDecoder(r.Body).Decode(&user)

	db.DB.Save(&user)

	if db.DB.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(db.DB.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}
