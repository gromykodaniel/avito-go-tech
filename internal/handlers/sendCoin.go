package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/azoma13/avito/internal/dataBase"
	"github.com/azoma13/avito/internal/utils"
	"github.com/azoma13/avito/models"
)

func SendCoinHandler(w http.ResponseWriter, r *http.Request) {
	username, err := utils.ValidateJWT(r)
	if err != nil {
		responseJSON(w, http.StatusUnauthorized, models.ErrorResponse{
			Errors: "error invailed token",
		})
		return
	}

	var sendCoin models.SentCoinRequest
	err = json.NewDecoder(r.Body).Decode(&sendCoin)
	if err != nil {

		responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Errors: "json deserialization error: " + err.Error(),
		})

		return
	}

	employeeToUser, err := dataBase.GetEmployeeDB(sendCoin.ToUser)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error to fetch toUser",
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

	if employee.ID == employeeToUser.ID {
		responseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error selection to user",
		})
		return
	}

	if employee.Balance < sendCoin.Amount {
		responseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error not enough balance",
		})
		return
	}

	err = dataBase.SendCoinDB(employee, employeeToUser, sendCoin)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Errors: "error send coin to user: " + err.Error(),
		})
		return
	}

}
