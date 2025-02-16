package handlers

import (
	"net/http"
	"strings"

	"github.com/azoma13/avito/internal/dataBase"
	"github.com/azoma13/avito/internal/utils"
	"github.com/azoma13/avito/models"
)

func BuyItemHandler(w http.ResponseWriter, r *http.Request) {
	merchName := strings.TrimPrefix(r.URL.Path, "/api/buy/")
	if merchName == "" {
		responseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error item name is missing",
		})
		return
	}

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

	merch, err := dataBase.GetMerchDB(merchName)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error item not found",
		})
		return
	}

	if employee.Balance < merch.Price {
		responseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error not enough balance",
		})
		return
	}

	err = dataBase.PayBuyMerchDB(employee, merch)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Errors: "error buy merch: " + err.Error(),
		})
		return
	}

}
