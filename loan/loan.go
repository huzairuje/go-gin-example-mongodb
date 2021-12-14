package loan

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Loan struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	FullName        string             `json:"full_name" bson:"full_name"`
	Gender          string             `json:"gender" bson:"gender"`
	KTPNumber       string             `json:"ktp_number" bson:"ktp_number"`
	ImageOfKTP      string             `json:"image_of_ktp" bson:"image_of_ktp"`
	ImageOfSelfie   string             `json:"image_of_selfie" bson:"image_of_selfie"`
	DateOfBirth     string             `json:"date_of_birth" bson:"date_of_birth"`
	Address         string             `json:"address" bson:"address"`
	AddressProvince string             `json:"address_province" bson:"address_province"`
	PhoneNumber     string             `json:"phone_number" bson:"phone_number"`
	Email           string             `json:"email" bson:"email"`
	Nationality     string             `json:"nationality" bson:"nationality"`
	LoanAmount      string             `json:"loan_amount" bson:"loan_amount"`
	Status          string             `json:"status" bson:"status"`
	Tenor           string             `json:"tenor" bson:"tenor"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	Installment     *Installment       `json:"installment" bson:"installment"`
}

type Installment struct {
	Id                     primitive.ObjectID `json:"id" bson:"_id"`
	LoanId                 primitive.ObjectID `json:"loan_id" bson:"loan_id"`
	AllPaidOff             bool               `json:"all_paid_off" bson:"all_paid_off"`
	TenorRemaining         int32              `json:"tenor_remaining" bson:"tenor_remaining"`
	TotalTenor             int32              `json:"total_tenor" bson:"total_tenor"`
	InstallmentAmount      string             `json:"installment_amount" bson:"installment_amount"`
	TotalInstallmentAmount string             `json:"total_installment_amount" bson:"total_installment_amount"`
	CreatedAt              time.Time          `json:"created_at" bson:"created_at"`
}
