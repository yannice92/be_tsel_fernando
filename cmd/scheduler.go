package main

import (
	"github.com/yannice92/be_tsel_fernando/config"
	"github.com/yannice92/be_tsel_fernando/repository"
	"github.com/yannice92/be_tsel_fernando/service"
	"log"
)

func main() {
	db := config.GetConnection()
	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}()
	repo := repository.NewMysqlCustomerRepository(db)
	cService := service.NewCustomerService(repo)
	err := cService.Provisioning()
	if err != nil {
		log.Println("scheduler error : ", err)
	}

}
