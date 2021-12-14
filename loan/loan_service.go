package loan

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
	"time"

	"github.com/go-gin-example-mongodb/utils"
)

type Service struct {
	ctx context.Context
	db  *mongo.Database
}

func NewLoanService(db *mongo.Database, ctx context.Context) Repository {
	return Service{
		ctx: ctx,
		db:  db}
}

func (p Service) Store(req CreateLoanRequest) (*Loan, error) {
	status := utils.RandomizeStatus()
	province := strings.ToUpper(req.AddressProvince)
	gender := strings.ToUpper(req.Gender)
	nationality := strings.ToUpper(req.Nationality)
	collection := p.db.Collection(utils.CollectionName)
	newObjID := primitive.NewObjectID()
	query := bson.M{
		"_id":              newObjID,
		"full_name":        req.FullName,
		"gender":           gender,
		"ktp_number":       req.KTPNumber,
		"image_of_ktp":     req.ImageOfKTP,
		"image_of_selfie":  req.ImageOfSelfie,
		"date_of_birth":    req.DateOfBirth,
		"address":          req.Address,
		"address_province": province,
		"phone_number":     req.PhoneNumber,
		"email":            req.Email,
		"nationality":      nationality,
		"loan_amount":      req.LoanAmount,
		"tenor":            req.Tenor,
		"status":           status,
		"created_at":       time.Now(),
	}

	result, err := collection.InsertOne(p.ctx, query)
	if err != nil {
		return nil, err
	}
	objLoanId := result.InsertedID
	loanId := objLoanId.(primitive.ObjectID)
	resultLoan, err := p.findLoanById(loanId)
	if err != nil {
		return nil, err
	}

	//insert into installment if the status is accepted
	if resultLoan.Status == utils.Accepted {
		installment, err := p.insertInstallment(loanId, req)
		if err != nil {
			return nil, err
		}
		resultLoan.Installment = installment
		return resultLoan, nil
	}
	return resultLoan, nil
}

func (p Service) List() ([]*Loan, error) {
	collection := p.db.Collection(utils.CollectionName)
	res, err := collection.Find(p.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var loans []*Loan
	for res.Next(p.ctx) {
		loan := &Loan{}
		err := res.Decode(loan)
		if err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}
	return loans, nil
}

func (p Service) isSurpassLimit() bool {
	row := p.db.Collection(utils.CollectionName)
	var count int64
	count, err := row.CountDocuments(p.ctx, bson.M{})
	if err != nil {
		logrus.Error(err)
		return false
	}
	if count >= utils.MaxLimitLoanProcessPerDay {
		return true
	}
	return false
}

func (p Service) IsKtpNumberExist(ktpNumber string) bool {
	row := p.db.Collection(utils.CollectionName)
	var count int64
	count, err := row.CountDocuments(p.ctx, bson.M{
		"ktp_number": ktpNumber,
	})
	if err != nil {
		logrus.Error(err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func (p Service) findLoanById(loanID primitive.ObjectID) (*Loan, error) {
	collection := p.db.Collection(utils.CollectionName)
	objIdString := loanID.Hex()
	objID, err := primitive.ObjectIDFromHex(objIdString)
	if err != nil {
		return nil, err
	}
	result := collection.FindOne(p.ctx, bson.M{
		"_id": objID,
	})
	var loan Loan
	err = result.Decode(&loan)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (p Service) insertInstallment(loanId primitive.ObjectID, req CreateLoanRequest) (*Installment, error) {
	intLoanAmount, _ := strconv.Atoi(req.LoanAmount)
	intTenor, _ := strconv.Atoi(req.Tenor)
	paymentAmount := intLoanAmount / intTenor
	paymentAmountFloat := float64(paymentAmount)
	paymentAmountFinalize := paymentAmountFloat + (paymentAmountFloat * utils.InterestRate)
	paymentAmountFinalizeString := strconv.FormatFloat(paymentAmountFinalize, 'f', 2, 64)
	row := p.db.Collection(utils.CollectionName)
	newObjID := primitive.NewObjectID()
	query := bson.M{
		"_id": loanId,
	}
	update := bson.M{
		"$set": bson.M{
			"installment": bson.M{
				"_id":                      newObjID,
				"loan_id":                  loanId,
				"all_paid_off":             false,
				"tenor_remaining":          intTenor,
				"total_tenor":              intTenor,
				"installment_amount":       paymentAmountFinalizeString,
				"total_installment_amount": req.LoanAmount,
				"created_at":               time.Now(),
			},
			"update_at": time.Now(),
		},
	}
	_, err := row.UpdateOne(p.ctx, query, update)
	if err != nil {
		return nil, err
	}
	installment, err := p.getInstallment(loanId)
	if err != nil {
		return nil, err
	}
	return installment, nil
}

func (p Service) getInstallment(loanId primitive.ObjectID) (*Installment, error) {
	collection := p.db.Collection(utils.CollectionName)
	objIdString := loanId.Hex()
	objID, err := primitive.ObjectIDFromHex(objIdString)
	if err != nil {
		return nil, err
	}
	result := collection.FindOne(p.ctx, bson.M{
		"_id": objID,
	})
	var installment Installment
	err = result.Decode(&installment)
	if err != nil {
		return nil, err
	}
	return &installment, nil
}
