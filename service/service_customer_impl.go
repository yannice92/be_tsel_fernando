package service

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/yannice92/be_tsel_fernando/model"
	"github.com/yannice92/be_tsel_fernando/repository"
	"github.com/yannice92/be_tsel_fernando/utils"
)

var mysqlErr *mysql.MySQLError

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(
	c repository.CustomerRepository,
) CustomerService {
	return &customerService{
		repo: c,
	}
}

func (c *customerService) AddCustomer(input model.CustomerRequest) error {
	var fromID uint64
	if !utils.IsEmpty(input.Referral) {
		fromID, _ = c.repo.CheckReferralExist(input.Referral)
		if fromID == 0 {
			return model.ReferralNotExist
		}
	}
	id, err := c.repo.AddCustomer(input)

	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return model.DuplicateMsisdn
		}
		return model.DefaultError
	}
	c.repo.InsertReferralCode(id)
	//add mapping, from_id always 0 if not have referral, maybe it will be useful for audit
	err = c.repo.InsertReferralMapping(fromID, id)
	if err != nil {
		return err
	}
	//always initialized data when user register
	err = c.repo.InsertSummaryReferralMonthly(id, yearmonth)
	if err != nil {
		return err
	}
	//update the summary table monthly
	if fromID > 0 {
		err = c.repo.UpdateTotalReferralMonthlyByCustomerID(fromID, yearmonth)
	}
	if err != nil {
		return err
	}
	return nil
}
