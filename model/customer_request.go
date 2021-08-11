package model

type CustomerRequest struct {
	Fullname string `json:"fullname"`
	Msisdn   uint64 `json:"msisdn"`
	Referral string `json:"referral"`
}