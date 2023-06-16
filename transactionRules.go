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
	"bytes"
	"encoding/xml"
	"fmt"
	"mime/multipart"
)

type Scope struct {
	Text             string `xml:",chardata"`
	ScopeDescription string `xml:"scope-description,attr"`
	ScopeName        string `xml:"scope-name,attr"`
	ScopeType        string `xml:"scope-type,attr"`
	ScopeVersion     string `xml:"scope-version,attr"`
}

type ScopeList struct {
	Text  string  `xml:",chardata"`
	Scope []Scope `xml:"scope"`
}

type TxMatchRuleStr struct {
	Text string `xml:",innerxml"`
}

type Rule struct {
	Text            string         `xml:",chardata"`
	AgentType       string         `xml:"agent-type,attr"`
	Enabled         bool           `xml:"enabled,attr"`
	Priority        string         `xml:"priority,attr"`
	RuleDescription string         `xml:"rule-description,attr"`
	RuleName        string         `xml:"rule-name,attr"`
	RuleType        string         `xml:"rule-type,attr"`
	Version         string         `xml:"version,attr"`
	TxMatchRule     TxMatchRuleStr `xml:"tx-match-rule"`
}

type RuleList struct {
	Text string `xml:",chardata"`
	Rule []Rule `xml:"rule"`
}

type ScopeRule struct {
	Text            string `xml:",chardata"`
	RuleDescription string `xml:"rule-description,attr"`
	RuleName        string `xml:"rule-name,attr"`
}

type ScopeRuleMapping struct {
	Text      string      `xml:",chardata"`
	ScopeName string      `xml:"scope-name,attr"`
	ScopeRule []ScopeRule `xml:"rule"`
}

type ScopeRuleMappingList struct {
	Text             string             `xml:",chardata"`
	ScopeRuleMapping []ScopeRuleMapping `xml:"scope-rule-mapping"`
}

type MdsData struct {
	XMLName              xml.Name             `xml:"mds-data"`
	Text                 string               `xml:",chardata"`
	ControllerVersion    string               `xml:"controller-version,attr"`
	ScopeList            ScopeList            `xml:"scope-list"`
	RuleList             RuleList             `xml:"rule-list"`
	ScopeRuleMappingList ScopeRuleMappingList `xml:"scope-rule-mapping-list"`
}

////////////////

type RuleMatchTxMatchRule struct {
	Type                string                       `json:"type"`
	Txautodiscoveryrule RuleMatchTxautodiscoveryrule `json:"txautodiscoveryrule"`
	Txcustomrule        RuleMatchTxcustomrule        `json:"txcustomrule"`
	Agenttype           string                       `json:"agenttype"`
}
type RuleMatchTxautodiscoveryrule struct {
	Autodiscoveryconfigs []interface{} `json:"autodiscoveryconfigs"`
}
type RuleMatchURI struct {
	Type         string   `json:"type"`
	Matchstrings []string `json:"matchstrings"`
}
type RuleMatchHttpmatch struct {
	Httpmethod string        `json:"httpmethod"`
	URI        RuleMatchURI  `json:"uri"`
	Parameters []interface{} `json:"parameters"`
	Headers    []interface{} `json:"headers"`
	Cookies    []interface{} `json:"cookies"`
}
type RuleMatchMatchconditions struct {
	Type      string             `json:"type"`
	Httpmatch RuleMatchHttpmatch `json:"httpmatch"`
}
type RuleMatchTxcustomrule struct {
	Type             string                     `json:"type"`
	Txentrypointtype string                     `json:"txentrypointtype"`
	Matchconditions  []RuleMatchMatchconditions `json:"matchconditions"`
	Actions          []interface{}              `json:"actions"`
	Properties       []interface{}              `json:"properties"`
}

type TxRuleUploadResult struct {
	Status string `json:"status"`
}

// TransactionRulesService provides transaction detection rules services
type TransactionRulesService service

// UploadTransactionRules - upload transaction detection rules for an application
func (s *TransactionRulesService) UploadTransactionRules(appNameOrId string, rules *MdsData) error {

	rulesXml, err := xml.Marshal(rules)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", string(rulesXml))

	url := "/controller/transactiondetection/" + appNameOrId + "/custom"

	multipartPayload := &bytes.Buffer{}
	writer := multipart.NewWriter(multipartPayload)
	part, err := writer.CreateFormFile("file", "rules.xml")
	if err != nil {
		return err
	}
	part.Write(rulesXml)

	err = writer.Close()
	if err != nil {
		return err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = writer.FormDataContentType()

	retval := TxRuleUploadResult{}

	// fmt.Printf("Upload body: %s, %v, %v\n", multipartPayload, s, headers)

	err = s.client.RestInternalHdr("POST", url, &retval, multipartPayload, headers)
	if err != nil {
		if fmt.Sprintf("%v", err) != "EOF" {
			fmt.Printf("Error uploading trx detection rules, %v\n", err)
		}
	}
	// fmt.Printf("Upload rules retval: %s\n", retval)

	return nil
}

type TxRulesResponse struct {
	RuleScopeSummaryMappings []TxRuleScopeSummaryMappings `json:"ruleScopeSummaryMappings"`
}

type TxRuleScopeSummaryMappings struct {
	Rule           TxRule                 `json:"rule"`
	ScopeSummaries []TxRuleScopeSummaries `json:"scopeSummaries"`
}
type TxMatchStrings struct {
	MatchStrings []string `json:"matchStrings"`
}
type TxHttpMatch struct {
	HttpMethod string         `json:"httpMethod"`
	Uri        TxMatchStrings `json:"uri"`
}
type TxMatchConditions struct {
	Type      string      `json:"type"`
	HttpMatch TxHttpMatch `json:"httpMatch"`
}
type TxCustomRule struct {
	Type             string              `json:"type"`
	TxEntryPointType string              `json:"txEntryPointType"`
	MatchConditions  []TxMatchConditions `json:"matchConditions"`
	Actions          []interface{}       `json:"actions"`
	Properties       []interface{}       `json:"properties"`
}
type TxMatchRule struct {
	Type         string       `json:"type"`
	TxCustomRule TxCustomRule `json:"txCustomRule"`
	AgentType    string       `json:"agentType"`
}
type TxRule struct {
	Type        string        `json:"type"`
	Summary     TxRuleSummary `json:"summary"`
	Enabled     bool          `json:"enabled"`
	Priority    int           `json:"priority"`
	Version     int           `json:"version"`
	AgentType   string        `json:"agentType"`
	TxMatchRule TxMatchRule   `json:"txMatchRule"`
}
type TxRuleScopeSummaries struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	AccountID string `json:"accountId"`
	Name      string `json:"name"`
}

type TxRuleSummary struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	AccountID   string `json:"accountId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   int64  `json:"createdOn"`
	UpdatedOn   int64  `json:"updatedOn"`
}

type ScopesResponse struct {
	Scopes []Scopes `json:"scopes"`
}
type ScopeSummary struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	AccountID   string `json:"accountId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   int64  `json:"createdOn"`
	UpdatedOn   int64  `json:"updatedOn"`
}
type ScopeIncludedTiers struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	AccountID string `json:"accountId"`
	Name      string `json:"name"`
}
type Scopes struct {
	Type          string               `json:"type"`
	Summary       ScopeSummary         `json:"summary"`
	Version       int                  `json:"version"`
	IncludedTiers []ScopeIncludedTiers `json:"includedTiers,omitempty"`
}

// GetTransactionDetectionRules - get list of transaction detection rules
func (s *TransactionRulesService) GetTransactionDetectionRules(appId string) (*TxRulesResponse, error) {

	url := fmt.Sprintf("/controller/restui/transactionConfigProto/getRules/%s", appId)

	var rules *TxRulesResponse
	err := s.client.RestInternal("GET", url, &rules, nil)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

// DeleteTransactionDetectionRule - delete transaction detection rule by id
func (s *TransactionRulesService) DeleteTransactionDetectionRule(ruleId string) error {

	url := "/controller/restui/transactionConfigProto/deleteRules"

	body := []string{ruleId}
	err := s.client.RestInternal("POST", url, nil, &body)
	if err != nil {
		return err
	}

	return nil
}

// GetTransactionDetectionRules - get list of transaction detection rules
func (s *TransactionRulesService) GetApplicationsScopes(appId int) (*ScopesResponse, error) {
	//todo
	url := fmt.Sprintf("/controller/restui/transactionConfigProto/getScopes/%d", appId)

	var scopes *ScopesResponse
	err := s.client.RestInternal("GET", url, &scopes, nil)
	if err != nil {
		return nil, err
	}

	return scopes, nil
}
