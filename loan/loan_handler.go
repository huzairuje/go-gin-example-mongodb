package loan

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"github.com/go-gin-example-mongodb/response"
	"github.com/go-gin-example-mongodb/utils"
	"github.com/go-gin-example-mongodb/validator"
)

type Handler struct {
	context        context.Context
	loanRepository Repository
	db             *mongo.Database
}

func NewHandler(db *mongo.Database) *Handler {
	var ctx context.Context
	repo := NewLoanService(db, ctx)
	return &Handler{
		context:        ctx,
		loanRepository: repo,
		db:             db,
	}
}

// Create godoc
// @Summary Create Loan Process
// @Description Post Loan Process
// @Tags loans
// @Accept */*
// @Produce  json
// @Success 200 {object} response.Single
// @Failure 400 {object} response.Single
// @Failure 422 {object} response.Single
// @Failure 500 {object} response.Single
// @Router /loans [post]
func (p Handler) Create(ctx *gin.Context) {
	var req CreateLoanRequest
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, utils.BadRequest, err.Error())
		return
	}
	isValidGender := validator.IsValidGender(req.Gender)
	if !isValidGender {
		response.BadRequest(ctx, utils.BadRequest, utils.GenderIsNotValid)
		return
	}
	isValidProvince := validator.IsValidProvince(req.AddressProvince)
	if !isValidProvince {
		response.BadRequest(ctx, utils.BadRequest, utils.ProvinceIsNotValid)
		return
	}
	isValidTenor := validator.IsValidTenor(req.Tenor)
	if !isValidTenor {
		response.BadRequest(ctx, utils.BadRequest, utils.TenorIsNotValid)
		return
	}
	dob, err := time.Parse(utils.DateFormat, req.DateOfBirth)
	if err != nil {
		response.BadRequest(ctx, utils.BadRequest, utils.AgeIsNotValidFormat)
		return
	}
	isValidAge := validator.IsValidAge(dob)
	if !isValidAge {
		response.BadRequest(ctx, utils.BadRequest, utils.AgeIsNotValid)
		return
	}
	isValidAmount := validator.IsValidLoanAmount(req.LoanAmount)
	if !isValidAmount {
		response.BadRequest(ctx, utils.BadRequest, utils.AmountLoanIsNotValid)
		return
	}
	isKtpNumberExist := p.loanRepository.IsKtpNumberExist(req.KTPNumber)
	if isKtpNumberExist {
		response.BadRequest(ctx, utils.BadRequest, utils.KTPNumberIsExist)
		return
	}
	isMaxLimitRequestLoan := p.loanRepository.isSurpassLimit()
	if isMaxLimitRequestLoan {
		response.BadRequest(ctx, utils.BadRequest, utils.MaxLimitRequestLoanProcessPerDay)
		return
	}
	data, err := p.loanRepository.Store(req)
	if err != nil {
		response.InternalServerError(ctx, utils.SomethingWentWrong, err.Error())
		return
	}
	response.SingleData(ctx, utils.LoanProcessIsSuccess, data)
	return
}

// List Loans godoc
// @Summary List Loans
// @Description get List Loans
// @Tags loans
// @Accept */*
// @Produce  json
// @Success 200 {object} response.Single
// @Failure 500 {object} response.Single
// @Produce  json
// @Router /loans [get]
func (p Handler) List(ctx *gin.Context) {
	data, err := p.loanRepository.List()
	if err != nil {
		response.InternalServerError(ctx, utils.SomethingWentWrong, err.Error())
		return
	}
	if len(data) == 0 {
		response.ListData(ctx, utils.OK, gin.H{})
		return
	}
	response.ListData(ctx, utils.OK, data)
	return
}
