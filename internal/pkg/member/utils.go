package member

func (responseError *Error) AddErrorDetail(errorDetail ErrorDetail) []ErrorDetail {
	responseError.Details = append(responseError.Details, errorDetail)
	return responseError.Details
}
