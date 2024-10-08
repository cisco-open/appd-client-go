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
	"fmt"
	"strconv"
)

type AffectsDetail struct {
	AffectedEntityType  *string `json:"affectedEntityType"`
	AffectedEntityScope struct {
		EntityScope        *string `json:"entityScope"`
		EntityType         *string `json:"entityType"`
		AffectedEntityName *string `json:"affectedEntityName"`
	} `json:"affectedEntityScope"`
}

type Condition struct {
	Name                   *string `json:"name"`
	ShortName              *string `json:"shortName"`
	EvaluateToTrueOnNoData *bool   `json:"evaluateToTrueOnNoData"`
	EvalDetail             struct {
		EvalDetailType          *string `json:"evalDetailType"`
		MetricAggregateFunction *string `json:"metricAggregateFunction"`
		MetricPath              *string `json:"metricPath"`
		MetricEvalDetail        struct {
			CompareCondition     *string `json:"compareCondition"`
			CompareValue         *int    `json:"compareValue"`
			MetricEvalDetailType *string `json:"metricEvalDetailType"`
			BaselineCondition    *string `json:"baselineCondition"`
			BaselineName         *string `json:"baselineName"`
			BaselineUnit         *string `json:"baselineUnit"`
		} `json:"metricEvalDetail"`
	} `json:"evalDetail"`
	TriggerEnabled  *bool `json:"triggerEnabled"`
	MinimumTriggers *int  `json:"minimumTriggers"`
}
type EvalCriteriaDetail struct {
	ConditionAggregationType *string      `json:"conditionAggregationType"`
	ConditionExpression      *string      `json:"conditionExpression"`
	Conditions               []*Condition `json:"conditions"`
	EvalMatchingCriteria     struct {
		MatchType *string `json:"matchType"`
		Value     *string `json:"value"`
	} `json:"evalMatchingCriteria"`
}

type EvalCriteriasSet struct {
	CriticalCriteria *EvalCriteriaDetail `json:"criticalCriteria"`
	WarningCriteria  *EvalCriteriaDetail `json:"warningCriteria"`
}

// HealthRule describes basic info about Helth rules attached to an application
// as returned by query for all health rules
type HealthRule struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Enabled            bool   `json:"enabled"`
	AffectedEntityType string `json:"affectedEntityType"`
}

// HealthRuleDetail describes detail information about specific health rule
type HealthRuleDetail struct {
	ID                      *int              `json:"id"`
	Name                    *string           `json:"name"`
	Enabled                 *bool             `json:"enabled"`
	UseDataFromLastNMinutes *int              `json:"useDataFromLastNMinutes"`
	WaitTimeAfterViolation  *int              `json:"waitTimeAfterViolation"`
	ScheduleName            *string           `json:"scheduleName"`
	Affects                 *AffectsDetail    `json:"affects"`
	EvalCriterias           *EvalCriteriasSet `json:"evalCriterias"`
}

// DANGER ZONE
// this is an UNPUBLISHED API call - it may change in the future
type HealthRuleEvaluationResponse []struct {
	AffectedEntity          HealthRuleAffectedEntity            `json:"affectedEntity"`
	Health                  string                              `json:"health"`
	EvaluationStatus        string                              `json:"evaluationStatus"`
	AggregationScopesStates []HealthRuleAggregationScopesStates `json:"aggregationScopesStates"`
	Name                    string                              `json:"name"`
	TierName                any                                 `json:"tierName"`
}
type HealthRuleAffectedEntity struct {
	ID             int    `json:"id"`
	Version        int    `json:"version"`
	EntityType     string `json:"entityType"`
	EntityID       int    `json:"entityId"`
	PrettyToString any    `json:"prettyToString"`
}
type HealthRuleAggregationScope struct {
	ID             int    `json:"id"`
	Version        int    `json:"version"`
	EntityType     string `json:"entityType"`
	EntityID       int    `json:"entityId"`
	PrettyToString any    `json:"prettyToString"`
}
type HealthRuleState struct {
	Result              string `json:"result"`
	Severity            string `json:"severity"`
	TriggeredConditions []any  `json:"triggeredConditions"`
}
type HealthRuleAggregationScopesStates struct {
	AggregationScope    HealthRuleAggregationScope `json:"aggregationScope"`
	JmxAggregationScope any                        `json:"jmxAggregationScope"`
	State               HealthRuleState            `json:"state"`
}

// DANGER ZONE END

// HealthRuleService intermediates Health Rules requests
type HealthRuleService service

// GetHealthRules obtains all backends for an application from a controller
func (s *HealthRuleService) GetHealthRules(appID int) ([]*HealthRule, error) {

	url := "controller/alerting/rest/v1/applications/" + strconv.Itoa(appID) + "/health-rules"

	var hrs []*HealthRule
	err := s.client.Rest("GET", url, &hrs, nil)
	if err != nil {
		return nil, err
	}

	return hrs, nil
}

// GetHealthRuleDetails obtains all backends for an application from a controller
func (s *HealthRuleService) GetHealthRuleDetails(appID int, ruleID int) (*HealthRuleDetail, error) {

	url := "controller/alerting/rest/v1/applications/" + strconv.Itoa(appID) + "/health-rules/" + strconv.Itoa(ruleID)

	var hr *HealthRuleDetail
	err := s.client.Rest("GET", url, &hr, nil)
	if err != nil {
		return nil, err
	}

	return hr, nil
}

// CreateHealthRule - create health rule for an application
func (s *HealthRuleService) CreateHealthRule(appID int, hr *HealthRuleDetail) error {

	url := "controller/alerting/rest/v1/applications/" + strconv.Itoa(appID) + "/health-rules"

	err := s.client.RestInternal("POST", url, nil, hr)
	if err != nil {
		if fmt.Sprintf("%s", err) == "EOF" { // successful call returns EOF error -> empty body
			return nil
		}
		fmt.Println(err)
		return err
	}

	return nil
}

// CreateHealthRule - create health rule for an application
func (s *HealthRuleService) CreateHealthRuleStr(appID int, hr *bytes.Buffer) error {

	urlStr := "controller/alerting/rest/v1/applications/" + strconv.Itoa(appID) + "/health-rules"
	req, err := s.client.newRequestBodyBytes("POST", urlStr, hr)

	if err != nil {
		return err
	}

	err = s.client.do(req, nil, true)
	if err != nil {
		if fmt.Sprintf("%s", err) == "EOF" { // successful call returns EOF error -> empty body
			return nil
		}
		fmt.Println(err)
		return err
	}

	return nil
}

// UpdateHealthRule - create health rule for an application
// TODO - test
func (s *HealthRuleService) UpdateHealthRule(appID int, ruleID int, hr *HealthRuleDetail) error {

	url := "controller/alerting/rest/v1/applications/" + strconv.Itoa(appID) + "/health-rules/" + strconv.Itoa(ruleID)

	err := s.client.RestInternal("PUT", url, nil, hr)
	if err != nil {
		if fmt.Sprintf("%s", err) == "EOF" { // successful call returns EOF error -> empty body
			return nil
		}
		fmt.Println(err)
		return err
	}

	return nil
}

// DeleteHealthRule - create health rule for an application
// TODO - test
func (s *HealthRuleService) DeleteHealthRule(appID int, ruleID int) error {

	url := "controller/alerting/rest/v1/applications/" + strconv.Itoa(appID) + "/health-rules/" + strconv.Itoa(ruleID)

	err := s.client.RestInternal("DELETE", url, nil, nil)
	if err != nil {
		if fmt.Sprintf("%s", err) == "EOF" { // successful call returns EOF error -> empty body
			return nil
		}
		fmt.Println(err)
		return err
	}

	return nil
}

// DANGER ZONE
// this is an UNPUBLISHED API call - it may change in the future
// GET /controller/restui/healthRules/getHealthRuleCurrentEvaluationStatus/app/3503/healthRuleID/22196
func (s *HealthRuleService) GetHealthRuleEvaluationState(appID int, ruleID int) (*HealthRuleEvaluationResponse, error) {

	url := "controller/restui/healthRules/getHealthRuleCurrentEvaluationStatus/app/" + strconv.Itoa(appID) + "/healthRuleID/" + strconv.Itoa(ruleID)

	hr := HealthRuleEvaluationResponse{}
	err := s.client.Rest("GET", url, &hr, nil)
	if err != nil {
		return nil, err
	}

	return &hr, nil
}

// DANGER ZONE END
