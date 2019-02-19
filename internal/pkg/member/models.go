package member

const (
	statusSuccess       = "SUCCEEDED"
	statusFail          = "FAILED"
	accountStatusActive = "ACTIVE"
	accountStatusBanned = "BANNED"
	databaseMember      = "go-member"
)

type Member struct {
	MemberID    string  `bson:"member_id"`
	FirstName     string  `bson:"first_name"`
	LastName      string  `bson:"last_name"`
	Email         string  `bson:"email"`
	MobileNumber  string  `bson:"mobile_number"`
	Address       Address `bson:"address"`
	AccountStatus string  `bson:"account_status"`
}

type CreateMemberAccountRequest struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	MobileNumber string  `json:"mobile_number"`
	Address      Address `json:"address"`
}

type CreateMemberAccountResponse struct {
	Status        string `json:"status"`
	MemberID    string `json:"member_id"`
	AccountStatus string `json:"account_status"`
	Error         *Error `json:"error,omitempty"`
}

type InquiryMemberAccountResponse struct {
	Status        string  `json:"status"`
	MemberID    string  `json:"member_id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Email         string  `json:"email"`
	MobileNumber  string  `json:"mobile_number"`
	Address       Address `json:"address"`
	AccountStatus string  `json:"account_status"`
	Error         *Error  `json:"error,omitempty"`
}

type Address struct {
	StreetAddress string `bson:"street_address" json:"street_address"`
	Subdistrict   string `bson:"subdistrict" json:"subdistrict"`
	District      string `bson:"district" json:"district"`
	Province      string `bson:"province" json:"province"`
	PostalCode    string `bson:"postal_code" json:"postal_code"`
}

type Error struct {
	Name    string        `bson:"name,omitempty" json:"name,omitempty"`
	Details []ErrorDetail `bson:"details,omitempty" json:"details,omitempty"`
}

type ErrorDetail struct {
	Field string `bson:"field,omitempty" json:"field,omitempty"`
	Issue string `bson:"issue,omitempty" json:"issue,omitempty"`
}

type IncrementIndex struct {
	MemberID int `bson:"member_id"`
}
