package repository

import (
	"database/sql"
	"github.com/yannice92/be_tsel_fernando/model"
	"log"
)

type CustomerRepository interface {
	AddCustomer(request model.CustomerRequest) (customerID uint64, err error)
	CheckReferralExist(referralCode string) (customerID uint64, err error)
	InsertReferralCode(id uint64) bool
	GetCustomerByReferralCode(referral string)

	InsertReferralMapping(fromID uint64, toID uint64) error

	//reward summary
	InsertSummaryReferralMonthly(customerID uint64, yearMonth uint64) error
	GetTotalReferralMonthlyByCustomerID(customerID uint64, yearMonth uint64) (total uint64, err error)
	UpdateTotalReferralMonthlyByCustomerID(customerID uint64, yearMonth uint64) error
	GetSummaryReferralList() ([]model.SummaryReferral, error)

	//rule
	GetRuleRewardList() ([]model.RuleReward,error)

	//provisioning
	InsertTableProvisioningMonthly(input model.Provisioning) error
}

func closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Print(err)
	}
}
