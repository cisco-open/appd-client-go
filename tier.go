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

// Tier represents one tier within one Application
type Tier struct {
	AgentType     string `json:"agentType"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ID            int    `json:"id"`
	NumberOfNodes int    `json:"numberOfNodes"`
	Type          string `json:"type"`
}

// TierService intermediates Tier requests
type TierService service

// GetTiers obtains all Tiers from an Application
func (s *TierService) GetTiers(appID int) ([]*Tier, error) {

	url := fmt.Sprintf("controller/rest/applications/%d/tiers?output=json", appID)

	var tiers []*Tier
	err := s.client.Rest("GET", url, &tiers, nil)
	if err != nil {
		return nil, err
	}

	return tiers, nil
}
