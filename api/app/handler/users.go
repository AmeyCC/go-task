package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AmeyCC/go-task/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, &users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	User := model.User{}

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
	name := vars["name"]
	user := getUserOr404(db, name, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)

}

// getUserOr404 gets a User instance if exists, or respond the 404 error otherwise
/*func getUserOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.User {
	User := model.User{}
	if err := db.First(&User, model.User{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &User
}*/

func getUserOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.User {
	User := model.User{}
	vars := mux.Vars(r)
	in := vars["in"]
	if err := db.Where("name LIKE ? ", "%"+in+"%").Find(&User, model.User{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &User
}
