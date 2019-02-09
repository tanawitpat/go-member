package member

func (responseError *Error) AddErrorDetail(errorDetail ErrorDetail) []ErrorDetail {
	responseError.Details = append(responseError.Details, errorDetail)
	return responseError.Details
}

func validateCreateMemberRequest(req CreateMemberAccountRequest) Error {
	responseError := Error{}

	if req.MobileNumber == "" {
		responseError.AddErrorDetail(ErrorDetail{Field: "mobile_number", Issue: "Field missing"})
	} else if len(req.MobileNumber) > 10 {
		responseError.AddErrorDetail(ErrorDetail{Field: "mobile_number", Issue: "mobile_number format is incorrect"})
	}

	return responseError
}
