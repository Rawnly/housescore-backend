package users_controller

import (
	"encoding/json"
	"housescore/database/models"
	"housescore/util"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request)  {
	body := models.SignUpBody{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&body); err != nil {
		// Da implementare
		log.Fatal(err.Error())
		return
	}

	defer r.Body.Close()

	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
		Phone:     body.Phone,
	}

	if err := user.Create(); err != nil {
		log.Fatal(err.Error())
	}

	util.SendJSON(w, http.StatusCreated, user)
}