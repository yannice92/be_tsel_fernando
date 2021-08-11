package service

import (
	"github.com/yannice92/be_tsel_fernando/model"
	"github.com/yannice92/be_tsel_fernando/utils"
)

var (
	yearmonth = utils.GetYearMonth()
)

type CustomerService interface {
	AddCustomer(request model.CustomerRequest) error
	GetTotalReferral(input model.ReferralRequest) (uint64, error)
	GetMonthlyTotalReward(input model.ReferralRequest) (reward string, err error)

	//provisioning
	Provisioning() error
}