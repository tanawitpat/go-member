package member

func validateCreateMemberRequest(req CreateMemberAccountRequest) Error {
	responseError := Error{}
	if req.MobileNumber == "" {
		responseError.AddErrorDetail(ErrorDetail{Field: "mobile_number", Issue: "Field missing"})
	} else if len(req.MobileNumber) > 10 {
		responseError.AddErrorDetail(ErrorDetail{Field: "mobile_number", Issue: "mobile_number format is incorrect"})
	}
	return responseError
}
