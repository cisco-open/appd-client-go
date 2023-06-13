package appdrest

import (
	"fmt"
	"time"
)

// Account is needed to access all operations here
type Account struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

// LicenseModules is small wrapper to a list of LicenseModule
type licenseModules struct {
	LicenseModules []*LicenseModule `json:"modules"`
	Links          []Link           `json:"links"`
}

// LicenseModule has an agent type and links to the properties and usage
type LicenseModule struct {
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

// Link has a name and a location
type Link struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

// usages is a wrapper for a list of Usage
type usages struct {
	Usages []*Usage `json:"usages"`
}

// Usage has the usage details for a license type
type Usage struct {
	ID                  string    `json:"id"`
	MaxUnitsUsed        float64   `json:"maxUnitsUsed"`
	MinUnitsUsed        float64   `json:"minUnitsUsed"`
	AvgUnitsUsed        float64   `json:"avgUnitsUsed"`
	TotalUnitsUsed      float64   `json:"totalUnitsUsed"`
	SampleCount         float64   `json:"sampleCount"`
	AvgUnitsAllowed     float64   `json:"avgUnitsAllowed"`
	AvgUnitsProvisioned float64   `json:"avgUnitsProvisioned"`
	AccountID           float64   `json:"accountId"`
	AgentType           string    `json:"agentType"`
	CreatedOn           int64     `json:"createdOn"`
	CreatedOnIsoDate    time.Time `json:"createdOnIsoDate"`
}

type properties struct {
	Properties []*Property `json:"properties"`
}

// Property is a simple license property
type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AccountService intermediates Account requests
type AccountService service

// GetMyAccount obtains an Account object
func (s *AccountService) GetMyAccount() (*Account, error) {

	url := "controller/api/accounts/myaccount"
	var account *Account
	err := s.client.Rest("GET", url, &account, nil)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// GetLicenseModules obtains all license modules and links
func (s *AccountService) GetLicenseModules(accID string) ([]*LicenseModule, error) {

	url := fmt.Sprintf("controller/api/accounts/%s/licensemodules?output=json", accID)

	var licenseModules *licenseModules
	err := s.client.Rest("GET", url, &licenseModules, nil)
	if err != nil {
		return nil, err
	}

	return licenseModules.LicenseModules, nil
}

// GetLicenseProperties obtains all properties for one agent type
func (s *AccountService) GetLicenseProperties(accID string, agentType string) ([]*Property, error) {

	url := fmt.Sprintf("controller/api/accounts/%s/licensemodules/%s/properties?output=json", accID, agentType)

	var licenseProperties *properties
	err := s.client.Rest("GET", url, &licenseProperties, nil)
	if err != nil {
		return nil, err
	}

	return licenseProperties.Properties, nil
}

// GetLicenseUsages obtains usage data for one agent type
func (s *AccountService) GetLicenseUsages(accID string, agentType string) ([]*Usage, error) {

	url := fmt.Sprintf("controller/api/accounts/%s/licensemodules/%s/usages?output=json", accID, agentType)

	var licenseUsages *usages
	err := s.client.Rest("GET", url, &licenseUsages, nil)
	if err != nil {
		return nil, err
	}

	return licenseUsages.Usages, nil
}
