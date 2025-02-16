package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/azoma13/avito/internal/dataBase"
	"github.com/azoma13/avito/internal/utils"
	"github.com/azoma13/avito/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func AuthEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	var reqEmployee models.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&reqEmployee)
	if err != nil {

		responseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "json deserialization error: " + err.Error(),
		})

		return
	}

	validate := validator.New()
	err = validate.Struct(reqEmployee)
	if err != nil {

		responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Errors: "failed validation during registration employee: " + err.Error(),
		})

		return
	}

	employee, err := dataBase.GetEmployeeDB(reqEmployee.Username)
	if err != nil {

		err := utils.RegisterEmployee(reqEmployee.Username, reqEmployee.Password)
		if err != nil {

			responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
				Errors: "failed to register employee: " + err.Error(),
			})

			return
		}
		employee, err = dataBase.GetEmployeeDB(reqEmployee.Username)
		if err != nil {

			responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
				Errors: "error to fetch employee after registration: " + err.Error(),
			})

			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(reqEmployee.Password))
	if err != nil {

		responseJSON(w, http.StatusUnauthorized, models.ErrorResponse{
			Errors: "invalid password: " + err.Error(),
		})

		return
	}

	token, err := utils.GenerateJWT(employee.Username)
	if err != nil {

		responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Errors: "error generate jwt token: " + err.Error(),
		})

		return
	}
	r.Header.Set("Authorization", "Bearer "+token)

	responseJSON(w, http.StatusOK, models.AuthResponse{
		Token: token,
	})
}
