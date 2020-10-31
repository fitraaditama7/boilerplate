package dashboard

import (
	"arka/cmd/lib/authentication"
	"arka/cmd/lib/customError"
	"arka/pkg/helper"
	"arka/pkg/response"
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (dashboard *authDashboard) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	var data string
	var params loginRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	err := helper.Validate(params)
	if err != nil {
		logrus.Error(err)
		response.Error(w, customError.ErrInvalidBodyRequest)
		return
	}

	if params.Email != "" {
		if authentication.IsEmail(params.Email) == true {
			data = params.Email
		} else {
			logrus.Error("error body request")
			response.Error(w, customError.ErrInvalidBodyRequest)
			return
		}
	} else if params.PhoneNumber != "" {
		data = params.PhoneNumber
	} else {
		logrus.Error("error body request")
		response.Error(w, customError.ErrInvalidBodyRequest)
		return
	}

	res, err := dashboard.service.Login(ctx, data, params.Password)
	if err != nil {
		logrus.Error(err)
		response.Error(w, err)
		return
	}

	response.Success(w, res)
}
