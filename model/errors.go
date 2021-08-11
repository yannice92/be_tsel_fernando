package model

import "errors"

var (
	ErrNoRowsAffected = errors.New("No rows has been affected")
	DuplicateMsisdn = errors.New("msisdn is duplicate")
	ReferralNotExist = errors.New("Referral Code doesn't exist")
	DefaultError = errors.New("Internal Error")
)
