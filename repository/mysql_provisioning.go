package repository

import (
	"github.com/yannice92/be_tsel_fernando/model"
	"github.com/yannice92/be_tsel_fernando/utils"
)

var (
	insertProvisioning = "INSERT INTO provisioning_reward (customer_id, offer_id, status, yearmonth) VALUES (?,?,?,?)"
)

func (m *mysqlCustomerRepository) InsertTableProvisioningMonthly(input model.Provisioning) error {
	yearmonth := utils.GetYearMonth()
	rows, err := m.Conn.Exec(insertProvisioning, input.CustomerID, input.OfferID, PROVISIONING_PROGRESS, yearmonth)
	if err != nil {
		return err
	}
	count, err := rows.RowsAffected()
	if count <= 0 {
		return model.ErrNoRowsAffected
	}
	_, err = rows.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
