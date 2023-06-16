/*
MIT License

Copyright (c) 2023 Cisco Systems, Inc. and its affiliates

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
	"strconv"
)

// Backend represents a single Backend within AppDynamics application
// Note that the REST version only has ID, Name and Description
type Backend struct {
	ID                         int    `json:"id"`
	TierID                     int    `json:"tierId"`
	Name                       string `json:"name"`
	ApplicationComponentNodeID int    `json:"applicationComponentNodeId"`
	Properties                 []struct {
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"properties"`
}

// BackendService intermediates Application requests
type BackendService service

// GetBackends obtains all backends for an application from a controller
func (s *BackendService) GetBackends(app string) ([]*Backend, error) {

	url := "controller/rest/applications/" + app + "/backends?output=json"

	var backends []*Backend
	err := s.client.Rest("GET", url, &backends, nil)
	if err != nil {
		return nil, err
	}

	return backends, nil
}

// ResolveBackendToTier - resolves Backend to an application Tier
// It might break in future versions of AppDynamics
func (s *BackendService) ResolveBackendToTier(backendID int, tierID int) error {

	backendIDstr := strconv.Itoa(backendID)
	tierIDstr := strconv.Itoa(tierID)
	url := fmt.Sprintf("controller/restui/backendUiService/resolveBackendToExistingTier/" + backendIDstr + "/" + tierIDstr)

	err := s.client.RestInternal("POST", url, nil, nil)
	if err != nil {
		if fmt.Sprintf("%s", err) == "EOF" { // successful call returns EOF error -> empty body
			return nil
		}
		fmt.Println(err)
		return err
	}

	return nil
}

// UnresolveBackendToTier - resolves Backend to an application Tier
// It might break in future versions of AppDynamics
// !!! this is unfinished !!! - explore how to send body request with POST
// in a form [<<backendID>>]
func (s *BackendService) UnresolveBackendToTier(backendID int) error {

	url := fmt.Sprintf("controller/restui/backendUiService/deleteBackends")

	body := []int{backendID}
	err := s.client.RestInternal("POST", url, nil, &body)
	if err != nil {
		return err
	}

	return nil
}
