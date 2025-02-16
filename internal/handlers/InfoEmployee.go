package handlers

import (
	"net/http"

	"github.com/azoma13/avito/internal/dataBase"
	"github.com/azoma13/avito/internal/utils"
	"github.com/azoma13/avito/models"
)

func InfoEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	username, err := utils.ValidateJWT(r)
	if err != nil {
		responseJSON(w, http.StatusUnauthorized, models.ErrorResponse{
			Errors: "error invailed token",
		})
		return
	}

	employee, err := dataBase.GetEmployeeDB(username)
	if err != nil {
		responseJSON(w, http.StatusUnauthorized, models.ErrorResponse{
			Errors: "error to fetch employee",
		})
		return
	}

	infoRes, err := dataBase.GetInfoEmployeeDB(employee)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Errors: "error send coin to user: " + err.Error(),
		})
		return
	}

	responseJSON(w, http.StatusOK, infoRes)
}
