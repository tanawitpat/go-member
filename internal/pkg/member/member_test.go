package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		MobileNumber: "0890000000",
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
