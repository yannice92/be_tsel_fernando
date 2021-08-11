package repository

import (
	"database/sql"
	"github.com/yannice92/be_tsel_fernando/model"
	"github.com/yannice92/be_tsel_fernando/utils"
	"log"
)

var (
	insertCustomer      = "INSERT INTO customer (fullname, msisdn) VALUES(?,?)"
	selectReferralExist = "SELECT id FROM customer WHERE referral_code = ? LIMIT 1"
	updateReferralCode  = "UPDATE customer SET referral_code = ? WHERE id = ?"
)

type mysqlCustomerRepository struct {
	Conn *sql.DB
}

func NewMysqlCustomerRepository(Conn *sql.DB) CustomerRepository {
	return &mysqlCustomerRepository{Conn}
}

//CheckReferralExist if true, referral is existed or error, otherwise is not existed
func (m *mysqlCustomerRepository) CheckReferralExist(referralCode string) (uint64, error) {
	var id uint64
	err := m.Conn.QueryRow(selectReferralExist, referralCode).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *mysqlCustomerRepository) AddCustomer(request model.CustomerRequest) (uint64, error) {
	rows, err := m.Conn.Exec(insertCustomer, request.Fullname, request.Msisdn)
	if err != nil {
		return 0, err
	}
	count, err := rows.RowsAffected()
	if count <= 0 {
		return 0, model.ErrNoRowsAffected
	}
	id, err := rows.LastInsertId()
	if err != nil {
		log.Println("error:", err)
		return 0, err
	}
	return uint64(id), nil
}

func (m *mysqlCustomerRepository) InsertReferralCode(id uint64) bool {
	referral := utils.ReferralCode(5)
	log.Println("Referral:", referral)
	_, err := m.Conn.Exec(updateReferralCode, referral, id)

	if err != nil {
		log.Println("error:", err)
		return false
	}
	return true
}

func (m *mysqlCustomerRepository) GetCustomerByReferralCode(referral string) {
	panic("implement me")
}