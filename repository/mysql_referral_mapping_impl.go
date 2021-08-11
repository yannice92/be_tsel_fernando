package repository

import (
	"github.com/yannice92/be_tsel_fernando/model"
)

var (
	insertReferralMapping                = "INSERT INTO referral_mapping (from_id, to_id) VALUES (?,?)"
	insertSummaryReferralMonthly         = "INSERT INTO summary_referral_monthly (customer_id, yearmonth) VALUES (?,?)"
	selectTotalReferralMonthlyCustomerID = "SELECT referral_total FROM summary_referral_monthly WHERE customer_id = ? AND yearmonth = ?"
	updateTotalReferralMonthlyCustomerID = "UPDATE summary_referral_monthly SET referral_total = referral_total + 1 WHERE customer_id = ? AND yearmonth = ?"
	selectRuleRewards                    = "SELECT referral_min_total, referral_max_total, offer_id FROM rule_reward"
	selectSummaryReferralNotZero         = "SELECT customer_id, yearmonth, referral_total FROM summary_referral_monthly WHERE referral_total != 0"
)

const (
	PROVISIONING_PROGRESS = "ONPROGRESS"
	PROVISIONING_SUCCESS = "SUCCESS"
)

func (m *mysqlCustomerRepository) InsertReferralMapping(fromID uint64, toID uint64) error {
	rows, err := m.Conn.Exec(insertReferralMapping, fromID, toID)
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

func (m *mysqlCustomerRepository) InsertSummaryReferralMonthly(customerID uint64, yearmonth uint64) error {

	rows, err := m.Conn.Exec(insertSummaryReferralMonthly, customerID, yearmonth)
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

func (m *mysqlCustomerRepository) GetTotalReferralMonthlyByCustomerID(customerID uint64, yearmonth uint64) (uint64, error) {
	var total uint64
	err := m.Conn.QueryRow(selectTotalReferralMonthlyCustomerID, customerID, yearmonth).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (m *mysqlCustomerRepository) UpdateTotalReferralMonthlyByCustomerID(customerID uint64, yearmonth uint64) error {
	rows, err := m.Conn.Exec(updateTotalReferralMonthlyCustomerID, customerID, yearmonth)
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

func (m *mysqlCustomerRepository) GetRuleRewardList() ([]model.RuleReward, error) {
	var rules []model.RuleReward
	rows, err := m.Conn.Query(selectRuleRewards)
	if err != nil {
		return rules, err
	}

	for rows.Next() {
		var rule model.RuleReward
		err := rows.Scan(&rule.Min, &rule.Max, &rule.OfferID)
		if err != nil {
			return rules, err
		}
		rules = append(rules, rule)
	}
	defer rows.Close()
	return rules, nil
}

func (m *mysqlCustomerRepository) GetSummaryReferralList() ([]model.SummaryReferral, error) {
	var list []model.SummaryReferral
	rows, err := m.Conn.Query(selectSummaryReferralNotZero)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var data model.SummaryReferral
		err := rows.Scan(&data.CustomerID, &data.YearMonth, &data.ReferralTotal)
		if err != nil {
			return list, err
		}
		list = append(list, data)
	}
	defer rows.Close()
	return list, nil
}


