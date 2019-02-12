package member

import (
	"regexp"
	"strings"
)

func validateCreateMemberRequest(req CreateMemberAccountRequest) Error {
	responseError := Error{}
	if req.FirstName == "" {
		responseError.AddErrorDetail(ErrorDetail{Field: "first_name", Issue: "Field missing"})
	}
	if req.LastName == "" {
		responseError.AddErrorDetail(ErrorDetail{Field: "last_name", Issue: "Field missing"})
	}
	if !validateAddress(req.Address) {
		responseError.AddErrorDetail(ErrorDetail{Field: "address", Issue: "Field incomplete"})
	}
	if req.MobileNumber == "" {
		responseError.AddErrorDetail(ErrorDetail{Field: "mobile_number", Issue: "Field missing"})
	} else if len(req.MobileNumber) > 10 || len(req.MobileNumber) < 9 {
		responseError.AddErrorDetail(ErrorDetail{Field: "mobile_number", Issue: "mobile_number format is incorrect"})
	}
	if req.Email == "" {
		responseError.AddErrorDetail(ErrorDetail{Field: "email", Issue: "Field missing"})
	} else if !validateEmailFormat(req.Email) {
		responseError.AddErrorDetail(ErrorDetail{Field: "email", Issue: "email format is incorrect"})
	}
	return responseError
}

func validateEmailFormat(input string) bool {
	input = strings.ToLower(input)
	regex := regexp.MustCompile("^[a-z0-9_%+]+[.\\-]?[a-z0-9]+@[a-z0-9\\-]+\\.[a-z]{2,4}$")
	return regex.MatchString(input)
}

func validateAddress(address Address) bool {
	if address.StreetAddress == "" ||
		address.Subdistrict == "" ||
		address.District == "" ||
		address.Province == "" ||
		address.PostalCode == "" {
		return false
	}
	return true
}
