package member

import (
	"go-member/internal/app"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func CreateMemberAccount(w http.ResponseWriter, r *http.Request) {
	req := CreateMemberAccountRequest{}
	res := CreateMemberAccountResponse{}
	responseError := Error{}

	requestID := r.Header.Get("request_id")
	if requestID == "" {
		log := app.InitLogger()
		log.Printf("request_id missing")
		responseError.AddErrorDetail(ErrorDetail{Field: "request_id", Issue: "Field missing"})
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("NULL - Request: %+v", req)
		log.Printf("NULL - Response: %+v", res)
		return
	}

	log := app.InitLogger().WithFields(logrus.Fields{"request_id": requestID})

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Printf("Cannot decode json")
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("Response: %#v", res)
		return
	}
	log.Printf("Request: %+v", req)

	if responseError := validateCreateMemberRequest(req); len(responseError.Details) != 0 {
		log.Printf("validateCreateMemberRequest: Failed")
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("Response: %+v", res)
		return
	}

	customerID, err := genCustomerID()
	if err != nil {
		log.Printf("Cannot generate customer ID: %+v", err)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.InternalServerError.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Printf("Response: %+v", res)
		return
	}

	member := Member{
		CustomerID:   customerID,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		MobileNumber: req.MobileNumber,
		Address: Address{
			StreetAddress: req.Address.StreetAddress,
			Subdistrict:   req.Address.Subdistrict,
			District:      req.Address.District,
			Province:      req.Address.Province,
			PostalCode:    req.Address.PostalCode,
		},
		AccountStatus: accountStatusActive,
	}

	if err := insertMemberDB(member); err != nil {
		log.Printf("Cannot insert member to the database: %+v", err)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.InternalServerError.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Printf("Response: %+v", res)
		return
	}

	res = CreateMemberAccountResponse{
		Status:        statusSuccess,
		CustomerID:    member.CustomerID,
		AccountStatus: member.AccountStatus,
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
	log.Printf("Response: %+v", res)
	return
}
