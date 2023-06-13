package appdrest

import (
	"fmt"
)

// BusinessTransaction represents one BT within one Application
type BusinessTransaction struct {
	InternalName   string `json:"internalName"`
	TierID         int    `json:"tierId"`
	EntryPointType string `json:"entryPointType"`
	Background     bool   `json:"background"`
	TierName       string `json:"tierName"`
	Name           string `json:"name"`
	ID             int    `json:"id"`
}

// BusinessTransactionService intermediates BusinessTransaction requests
type BusinessTransactionService service

// GetBusinessTransactions obtains all BTs from an application
func (s *BusinessTransactionService) GetBusinessTransactions(appID int) ([]*BusinessTransaction, error) {

	url := fmt.Sprintf("controller/rest/applications/%d/business-transactions?output=json", appID)

	var bts []*BusinessTransaction
	err := s.client.Rest("GET", url, &bts, nil)
	if err != nil {
		return nil, err
	}

	return bts, nil
}
