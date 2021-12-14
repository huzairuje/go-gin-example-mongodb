package loan

type CreateLoanRequest struct {
	FullName        string `json:"full_name" validate:"required"`
	Gender          string `json:"gender" validate:"required"`
	KTPNumber       string `json:"ktp_number" validate:"required"`
	ImageOfKTP      string `json:"image_of_ktp" validate:"required"`
	ImageOfSelfie   string `json:"image_of_selfie" validate:"required"`
	DateOfBirth     string `json:"date_of_birth" validate:"required"`
	Address         string `json:"address" validate:"required"`
	AddressProvince string `json:"address_province" validate:"required"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Nationality     string `json:"nationality" validate:"required"`
	LoanAmount      string `json:"loan_amount" validate:"required"`
	Tenor           string `json:"tenor" validate:"required"`
}
