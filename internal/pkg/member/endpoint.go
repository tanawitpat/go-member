package member

import (
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
	log.Printf("%s - Request: %+v", requestID, req)

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

	res.Status = statusSuccess
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
	log.Printf("%s - Response: %+v", requestID, res)
	return
}
