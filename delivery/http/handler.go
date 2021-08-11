package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yannice92/be_tsel_fernando/model"
	"github.com/yannice92/be_tsel_fernando/service"
	"log"
	"net/http"
)

type CustomerHandler struct {
	CUseCase service.CustomerService
}

func NewCustomerHandler(customerSrv service.CustomerService) *httprouter.Router {
	handler := &CustomerHandler{
		CUseCase: customerSrv,
	}

	myRouter := httprouter.New()
	myRouter.POST("/customers", handler.AddCustomer)
	myRouter.GET("/customers/referral", handler.GetReferral)
	myRouter.GET("/customers/reward", handler.GetReward)
	return myRouter
}

func (handler *CustomerHandler) AddCustomer(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	request := model.CustomerRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		response := &model.DataResponse{Message: err.Error()}
		ResponseWriter(req, rw, response, http.StatusBadRequest)
		return
	}
	err = handler.CUseCase.AddCustomer(request)
	if err != nil {
		response := &model.DataResponse{Code: "0001", Message: err.Error()}
		log.Println("error AddCustomer: ", err)
		ResponseWriter(req, rw, response, http.StatusBadRequest)
		return
	}
	reponseModel := model.DataResponse{
		Code:    "0000",
		Message: "Success",
		Data:    nil,
	}
	ResponseWriter(req, rw, reponseModel, 200)
}

func (handler *CustomerHandler) GetReferral(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	request := model.ReferralRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		response := &model.DataResponse{Code: "0001", Message: err.Error()}
		ResponseWriter(req, rw, response, http.StatusBadRequest)
		return
	}
	total, err := handler.CUseCase.GetTotalReferral(request)
	if err != nil {
		response := &model.DataResponse{Code: "0005", Message: model.DefaultError.Error()}
		ResponseWriter(req, rw, response, http.StatusInternalServerError)
		return
	}
	response := model.DataResponse{
		Code:    "0000",
		Message: "Success",
		Data: model.ReferralResponse{
			Total: total,
		},
	}
	ResponseWriter(req, rw, response, http.StatusOK)
	return
}

func (handler *CustomerHandler) GetReward(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	request := model.ReferralRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		response := &model.DataResponse{Code: "0001", Message: err.Error()}
		ResponseWriter(req, rw, response, http.StatusBadRequest)
		return
	}
	reward, err := handler.CUseCase.GetMonthlyTotalReward(request)
	if err != nil {
		response := &model.DataResponse{Code: "0005", Message: model.DefaultError.Error()}
		ResponseWriter(req, rw, response, http.StatusInternalServerError)
		return
	}
	response := model.DataResponse{
		Code:    "0000",
		Message: "Success",
		Data: model.RewardResponse{
			Reward: reward,
		},
	}
	ResponseWriter(req, rw, response, http.StatusOK)
	return
}

func ResponseWriter(r *http.Request, rw http.ResponseWriter, response interface{}, statusCode int) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(response)
}
