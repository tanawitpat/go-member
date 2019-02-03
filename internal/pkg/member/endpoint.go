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
		responseError.AddError(ErrorDetail{Field: "request_id", Issue: "Field missing"})
		res.Status = statusFail
		res.Error = &Error{
			Name:    responseNameBadRequest,
			Details: responseError.Details,
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		log.Printf("NULL - Response: %#v", res)
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

	res.Status = statusSuccess
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
	log.Printf("%s - Response: %#v", requestID, res)
	return
}
