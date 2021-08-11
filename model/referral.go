package model

type RuleReward struct {
	Min uint64
	Max uint64
	OfferID string
}

type SummaryReferral struct {
	CustomerID uint64
	YearMonth uint64
	ReferralTotal uint64
}

type Provisioning struct {
	SummaryReferral
	OfferID string
}