package service

import (
	"github.com/yannice92/be_tsel_fernando/model"
	"log"
)

func (c *customerService) GetTotalReferral(input model.ReferralRequest) (uint64, error) {
	total, err := c.repo.GetTotalReferralMonthlyByCustomerID(input.CustomerID, yearmonth)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (c *customerService) GetMonthlyTotalReward(input model.ReferralRequest) (string, error) {
	//get rule reward
	rules, err := c.repo.GetRuleRewardList()
	if err != nil {
		return "", err
	}
	total, err := c.GetTotalReferral(input)
	for _, rule := range rules {
		if total >= rule.Min && total <= rule.Max {
			return rule.OfferID, nil
		}
	}
	return "", nil
}

//private provisioning
func (c *customerService) Provisioning() error {
	list, err := c.repo.GetSummaryReferralList()
	if err != nil {
		log.Println("failed get Summary Rewards List:", err)
	}
	for _, referral := range list {
		referralReq := model.ReferralRequest{CustomerID: referral.CustomerID}
		reward, err := c.GetMonthlyTotalReward(referralReq)
		if err != nil {
			log.Println("Failed get Total Rewards:", err)
		} else {
			input := model.Provisioning{
				SummaryReferral: referral,
				OfferID:         reward,
			}
			err := c.repo.InsertTableProvisioningMonthly(input)
			if err != nil {
				log.Println("Failed insert Provisioning:", input.CustomerID, err)
			} else {
				log.Println("Success insert Provisioning:", input.CustomerID)
			}
		}
	}
	return nil
}
