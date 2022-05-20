package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AmeyCC/go-task/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Users := []model.User{}
	db.Find(&Users)
	respondJSON(w, http.StatusOK, Users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	User := []model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&User); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&User).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, User)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	User := getUserOr404(db, id, w, r)
	if User == nil {
		return
	}
	respondJSON(w, http.StatusOK, User)
}

// getUserOr404 gets a User instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.User {
	User := model.User{}
	if err := db.First(&User, model.User{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &User
}
