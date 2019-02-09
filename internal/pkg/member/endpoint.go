package member

import (
	"go-member/internal/app"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

func CreateMemberAccount(w http.ResponseWriter, r *http.Request) {
	req := CreateMemberAccountRequest{}
	res := CreateMemberAccountResponse{}
	responseError := Error{}
	requestID := r.Header.Get("request_id")

	if requestID == "" {
		log.Printf("request_id missing")
		responseError.AddErrorDetail(ErrorDetail{Field: "request_id", Issue: "Field missing"})
		res.Status = statusFail
		res.Error = &Error{
			Name:    responseNameBadRequest,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("NULL - Request: %+v", req)
		log.Printf("NULL - Response: %+v", res)
		return
	}

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Printf("Cannot decode json")
		res.Status = statusFail
		res.Error = &Error{
			Name:    responseNameBadRequest,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("%s - Response: %#v", requestID, res)
		return
	}
	log.Printf("%s - Request: %+v", requestID, req)

	db, err := app.GetMongoSession()
	if err != nil {
		log.Printf("%+v", err)
		res.Status = statusFail
		res.Error = &Error{
			Name:    responseNameInternalServerError,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Printf("%s - Response: %+v", requestID, res)
		return
	}

	if responseError := validateCreateMemberRequest(req); len(responseError.Details) != 0 {
		log.Printf("%s - validateCreateMemberRequest: Failed", requestID)
		res.Status = statusFail
		res.Error = &Error{
			Name:    responseNameBadRequest,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("%s - Response: %+v", requestID, res)
		return
	}

	member := Member{
		CustomerID:   "1",
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
		AccountStatus: "ACTIVE",
	}

	if err := db.C("member").Insert(member); err != nil {
		log.Printf("%+v", err)
	}

	res = CreateMemberAccountResponse{
		Status:        statusSuccess,
		CustomerID:    member.CustomerID,
		AccountStatus: member.AccountStatus,
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
	log.Printf("%s - Response: %+v", requestID, res)
	return
}
