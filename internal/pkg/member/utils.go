package member

func (responseError *Error) AddError(errorDetail ErrorDetail) []ErrorDetail {
	responseError.Details = append(responseError.Details, errorDetail)
	return responseError.Details
}
