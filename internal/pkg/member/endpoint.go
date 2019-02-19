package member

import (
	"go-member/internal/app"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
)

func CreateMemberAccount(w http.ResponseWriter, r *http.Request) {
	req := CreateMemberAccountRequest{}
	res := CreateMemberAccountResponse{}
	responseError := Error{}

	requestID := r.Header.Get("request_id")
	log := app.InitLoggerEndpoint(r)

	if requestID == "" {
		log.Errorf("request_id missing")
		responseError.AddErrorDetail(ErrorDetail{Field: "request_id", Issue: "Field missing"})
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Infof("Request: %+v", req)
		log.Infof("Response: %+v", res)
		return
	}

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Errorf("Cannot decode json")
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Infof("Response: %#v", res)
		return
	}
	log.Infof("Request: %+v", req)

	session, err := app.GetMongoSession()
	if err != nil {
		log.Errorf("Cannot get mongo session")
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.InternalServerError.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Infof("Response: %#v", res)
		return
	}
	defer session.Close()
	db := session.DB(databaseMember)

	if responseError := validateCreateMemberRequest(req); len(responseError.Details) != 0 {
		log.Errorf("validateCreateMemberRequest failed")
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.BadRequest.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Infof("Response: %+v", res)
		return
	}

	customerID, err := genCustomerID(db)
	if err != nil {
		log.Errorf("Cannot generate customer ID: %+v", err)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.InternalServerError.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Infof("Response: %+v", res)
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

	if err := db.C("member").Insert(member); err != nil {
		log.Errorf("Cannot insert member to the database: %+v", err)
		res.Status = statusFail
		res.Error = &Error{
			Name:    app.EM.Internal.InternalServerError.Name,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Infof("Response: %+v", res)
		return
	}

	res = CreateMemberAccountResponse{
		Status:        statusSuccess,
		CustomerID:    member.CustomerID,
		AccountStatus: member.AccountStatus,
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
	log.Infof("Response: %+v", res)
	return
}

func InquiryMemberAccount(w http.ResponseWriter, r *http.Request) {
	res := InquiryMemberAccountResponse{}
	log := app.InitLoggerEndpoint(r)
	urlParameter := mux.Vars(r)
	res.CustomerID = urlParameter["customerID"]

	session, err := app.GetMongoSession()
	if err != nil {
		log.Errorf("Cannot get mongo session: %+v", err)
		res.Status = statusFail
		res.Error = &Error{
			Name: app.EM.Internal.InternalServerError.Name,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Infof("Response: %#v", res)
		return
	}
	defer session.Close()
	db := session.DB(databaseMember)

	member := Member{}
	if err := db.C("member").Find(bson.M{"customer_id": urlParameter["customerID"]}).One(&member); err != nil {
		if err.Error() == "not found" {
			log.Errorf("Account not found: %+v", err)
			res.Status = statusFail
			res.Error = &Error{
				Name: app.EM.Internal.AccountNotFound.Name,
			}
			render.Status(r, http.StatusOK)
			render.JSON(w, r, res)
			log.Infof("Response: %#v", res)
			return
		}
		log.Errorf("Cannot get member data from the database: %+v", err)
		res.Status = statusFail
		res.Error = &Error{
			Name: app.EM.Internal.InternalServerError.Name,
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		log.Infof("Response: %#v", res)
		return
	}

	res = InquiryMemberAccountResponse{
		Status:        statusSuccess,
		CustomerID:    urlParameter["customerID"],
		FirstName:     member.FirstName,
		LastName:      member.LastName,
		Email:         member.Email,
		MobileNumber:  member.MobileNumber,
		Address:       member.Address,
		AccountStatus: member.AccountStatus,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
	log.Infof("Response: %+v", res)
	return
}
