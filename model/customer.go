package model

type Customer struct {
	ID           uint64 `json:"id"`
	Fullname     string `json:"fullname"`
	ReferralCode string `json:"referral_code"`
	Msisdn       uint64 `json:"msisdn"`
}
