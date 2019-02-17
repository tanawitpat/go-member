package member

import (
	"go-member/internal/app"
	"net/http"

	"github.com/go-chi/render"
)

func CreateMemberAccount(w http.ResponseWriter, r *http.Request) {
	log := app.InitLogger()
	req := CreateMemberAccountRequest{}
	res := CreateMemberAccountResponse{}
	responseError := Error{}
	requestID := r.Header.Get("request_id")

	if requestID == "" {
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

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Printf("%s - Cannot decode json", requestID)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("%s - Response: %#v", requestID, res)
		return
	}
	log.Printf("%s - Request: %+v", requestID, req)

	if responseError := validateCreateMemberRequest(req); len(responseError.Details) != 0 {
		log.Printf("%s - validateCreateMemberRequest: Failed", requestID)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("%s - Response: %+v", requestID, res)
		return
	}

	customerID, err := genCustomerID()
	if err != nil {
		log.Printf("%s - Cannot generate customer ID: %+v", requestID, err)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.InternalServerError.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Printf("%s - Response: %+v", requestID, res)
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
		log.Printf("%s - Cannot insert member to the database: %+v", requestID, err)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.InternalServerError.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Printf("%s - Response: %+v", requestID, res)
		return
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
