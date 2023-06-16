/*
MIT License

Copyright (c) 2023 David Lopes

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

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
