package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"log"
	configPkg "github.com/yannice92/be_tsel_fernando/config"
	"github.com/yannice92/be_tsel_fernando/model"
	"testing"
)

type config struct {
	DB *sql.DB
}

func Conn() {

}

func TestAddCustomer(t *testing.T) {
	configPkg.GetEnv()
	db := configPkg.GetConnection()
	repo := NewMysqlCustomerRepository(db)
	input := model.CustomerRequest{
		Fullname: "Test",
		Msisdn:   6282123333333,
	}
	_, err := repo.AddCustomer(input)
	if err != nil {
		t.Error("gagal insert :", err)
	} else {
		log.Println("success")
	}
}

func TestCheckReferralExist(t *testing.T) {
	configPkg.GetEnv()
	db := configPkg.GetConnection()
	repo := NewMysqlCustomerRepository(db)
	check, _ := repo.CheckReferralExist("TEST123s")
	log.Println("ID:",check)
}