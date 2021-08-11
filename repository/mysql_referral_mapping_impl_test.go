package repository

import (
	"log"
	configPkg "github.com/yannice92/be_tsel_fernando/config"
	"testing"
)

func TestMysqlCustomerRepository_GetRuleRewardList(t *testing.T) {
	configPkg.GetEnv()
	db := configPkg.GetConnection()
	repo := NewMysqlCustomerRepository(db)
	check, _ := repo.GetRuleRewardList()
	log.Println("rules:",check)
}