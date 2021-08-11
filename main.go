package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
	"github.com/yannice92/be_tsel_fernando/config"
	httpHandler "github.com/yannice92/be_tsel_fernando/delivery/http"
	"github.com/yannice92/be_tsel_fernando/repository"
	"github.com/yannice92/be_tsel_fernando/service"
	"log"
	"net/http"
	"os"
)

func main() {
	db := config.GetConnection()
	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
		log.Println("close")
	}()
	repo := repository.NewMysqlCustomerRepository(db)
	cService := service.NewCustomerService(repo)
	handler := httpHandler.NewCustomerHandler(cService)
	serveHttp(handler)
}

func serveHttp(router *httprouter.Router) {
	addr := os.Getenv("SERVER_ADDR")
	log.Println("server is listening on:", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalln("server failed listening on : ", addr)
	}

}
