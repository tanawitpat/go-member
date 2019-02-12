package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCreateMemberRequestPositive(t *testing.T) {
	req := CreateMemberAccountRequest{
		FirstName:    "Tanawit",
		LastName:     "Pattanaveerangkoon",
		MobileNumber: "028888888",
		Email:        "abc@gmail.com",
		Address: Address{
			StreetAddress: "100/100 Yotha Rd.",
			Subdistrict:   "Talad Noi",
			District:      "Samphanthawong",
			Province:      "Bangkok",
			PostalCode:    "10100",
		},
	}
	responseError := validateCreateMemberRequest(req)
	assert.Equal(t, 0, len(responseError.Details))
}

func TestValidateCreateMemberRequestInvalidMobileNumber(t *testing.T) {
	req := CreateMemberAccountRequest{
		FirstName:    "Tanawit",
		LastName:     "Pattanaveerangkoon",
		MobileNumber: "02999",
		Email:        "abc@gmail.com",
		Address: Address{
			StreetAddress: "100/100 Yotha Rd.",
			Subdistrict:   "Talad Noi",
			District:      "Samphanthawong",
			Province:      "Bangkok",
			PostalCode:    "10100",
		},
	}
	responseError := validateCreateMemberRequest(req)
	assert.Equal(t, 1, len(responseError.Details))
	if len(responseError.Details) != 0 {
		assert.Equal(t, "mobile_number", responseError.Details[0].Field)
		assert.Equal(t, "mobile_number format is incorrect", responseError.Details[0].Issue)
	}
}

func TestValidateCreateMemberRequestFirstNameMissing(t *testing.T) {
	req := CreateMemberAccountRequest{
		LastName:     "Pattanaveerangkoon",
		MobileNumber: "028888888",
		Email:        "abc@gmail.com",
		Address: Address{
			StreetAddress: "100/100 Yotha Rd.",
			Subdistrict:   "Talad Noi",
			District:      "Samphanthawong",
			Province:      "Bangkok",
			PostalCode:    "10100",
		},
	}
	responseError := validateCreateMemberRequest(req)
	assert.Equal(t, 1, len(responseError.Details))
	if len(responseError.Details) != 0 {
		assert.Equal(t, "first_name", responseError.Details[0].Field)
		assert.Equal(t, "Field missing", responseError.Details[0].Issue)
	}
}

func TestValidateCreateMemberRequestFirstNameLastNameEmailMissing(t *testing.T) {
	req := CreateMemberAccountRequest{
		MobileNumber: "023331111",
		Address: Address{
			StreetAddress: "100/100 Yotha Rd.",
			Subdistrict:   "Talad Noi",
			District:      "Samphanthawong",
			Province:      "Bangkok",
			PostalCode:    "10100",
		},
	}
	responseError := validateCreateMemberRequest(req)
	t.Log(responseError)
	assert.Equal(t, 3, len(responseError.Details))
}

func TestValidateEmailFormatPositive(t *testing.T) {
	testEmailInvalid := []string{"tanawit.p", "tanawit.pat@"}
	for _, email := range testEmailInvalid {
		isValid := validateEmailFormat(email)
		assert.Equal(t, false, isValid)
	}
}

func TestValidateEmailFormatNegative(t *testing.T) {
	testEmailValid := []string{"tanawit.p@google.tech", "tanawit@google.com", "tanawit@google.in.th"}
	for _, email := range testEmailValid {
		isValid := validateEmailFormat(email)
		assert.Equal(t, true, isValid)
	}
}
