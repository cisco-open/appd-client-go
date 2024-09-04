/*
MIT License

Copyright (c) 2023 David Lopes
Copyright (c) 2024 Cisco Systems, Inc. and its affiliates

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

// allApplicationTypes is a wrapper on the json response of GetApplicationAllTypes
type applicationAllTypes struct {
	Applications []*Application `json:"apmApplications"`
}

// Application represents a single Business Application within AppDynamics
// Note that the REST version only has ID, Name and Description
type Application struct {
	ID                    int           `json:"id"`
	Version               int           `json:"version"`
	Name                  string        `json:"name"`
	NameUnique            bool          `json:"nameUnique"`
	BuiltIn               bool          `json:"builtIn"`
	CreatedBy             string        `json:"createdBy"`
	CreatedOn             int64         `json:"createdOn"`
	ModifiedBy            string        `json:"modifiedBy"`
	ModifiedOn            int64         `json:"modifiedOn"`
	Description           string        `json:"description"`
	Template              bool          `json:"template"`
	Active                bool          `json:"active"`
	Running               bool          `json:"running"`
	RunningSince          interface{}   `json:"runningSince"`
	DeployWorkflowID      int           `json:"deployWorkflowId"`
	UndeployWorkflowID    int           `json:"undeployWorkflowId"`
	Visualization         interface{}   `json:"visualization"`
	EnvironmentProperties []interface{} `json:"environmentProperties"`
	EumAppName            string        `json:"eumAppName"`
	ApplicationTypeInfo   struct {
		ApplicationTypes   []string `json:"applicationTypes"`
		EumEnabled         bool     `json:"eumEnabled"`
		EumWebEnabled      bool     `json:"eumWebEnabled"`
		EumMobileEnabled   bool     `json:"eumMobileEnabled"`
		EumIotEnabled      bool     `json:"eumIotEnabled"`
		HasEumWebEntities  bool     `json:"hasEumWebEntities"`
		HasMobileApps      bool     `json:"hasMobileApps"`
		HasTiers           bool     `json:"hasTiers"`
		NumberOfMobileApps int      `json:"numberOfMobileApps"`
	} `json:"applicationTypeInfo"`
}

// DANGER ZONE
// following types are used for UNPUBLISHED api call and it may change in the future
type AllInternalApplications struct {
	ApmApplications              []GenericApplication `json:"apmApplications"`
	EumWebApplications           []GenericApplication `json:"eumWebApplications"`
	DbMonApplication             GenericApplication   `json:"dbMonApplication"`
	OverageMonitoringApplication GenericApplication   `json:"overageMonitoringApplication"`
	SimApplication               GenericApplication   `json:"simApplication"`
	AnalyticsApplication         GenericApplication   `json:"analyticsApplication"`
	MobileAppContainers          []GenericApplication `json:"mobileAppContainers"`
	IotApplications              []GenericApplication `json:"iotApplications"`
	CloudMonitoringApplication   GenericApplication   `json:"cloudMonitoringApplication"`
	APIMonitoringApplications    []GenericApplication `json:"apiMonitoringApplications"`
	CoreWebVitalsApplication     GenericApplication   `json:"coreWebVitalsApplication"`
}
type GenericApplication struct {
	ID                    int                     `json:"id"`
	Version               int                     `json:"version"`
	Name                  string                  `json:"name"`
	NameUnique            bool                    `json:"nameUnique"`
	BuiltIn               bool                    `json:"builtIn"`
	CreatedBy             string                  `json:"createdBy"`
	CreatedOn             int64                   `json:"createdOn"`
	ModifiedBy            string                  `json:"modifiedBy"`
	ModifiedOn            int64                   `json:"modifiedOn"`
	Description           string                  `json:"description"`
	Template              bool                    `json:"template"`
	Active                bool                    `json:"active"`
	Running               bool                    `json:"running"`
	RunningSince          any                     `json:"runningSince"`
	DeployWorkflowID      int                     `json:"deployWorkflowId"`
	UndeployWorkflowID    int                     `json:"undeployWorkflowId"`
	Visualization         any                     `json:"visualization"`
	EnvironmentProperties []EnvironmentProperties `json:"environmentProperties"`
	EumAppName            string                  `json:"eumAppName"`
	AccountGUID           string                  `json:"accountGuid"`
	ApplicationTypeInfo   ApplicationTypeInfo     `json:"applicationTypeInfo"`
}
type ApplicationTypeInfo struct {
	ApplicationTypes        []string `json:"applicationTypes"`
	EumEnabled              bool     `json:"eumEnabled"`
	EumWebEnabled           bool     `json:"eumWebEnabled"`
	EumMobileEnabled        bool     `json:"eumMobileEnabled"`
	EumIotEnabled           bool     `json:"eumIotEnabled"`
	EumAPIMonitoringEnabled bool     `json:"eumApiMonitoringEnabled"`
	HasEumWebEntities       bool     `json:"hasEumWebEntities"`
	HasMobileApps           bool     `json:"hasMobileApps"`
	HasTiers                bool     `json:"hasTiers"`
	NumberOfMobileApps      int      `json:"numberOfMobileApps"`
}
type EnvironmentProperties struct {
	ID      int    `json:"id"`
	Version int    `json:"version"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

// DANGER ZONE END

// ApplicationService intermediates Application requests
type ApplicationService service

// GetApplications obtains all applications from a controller
func (s *ApplicationService) GetApplications() ([]*Application, error) {

	url := "controller/rest/applications?output=json"

	var apps []*Application
	err := s.client.Rest("GET", url, &apps, nil)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

// GetApplication gets an Application by Name or ID
func (s *ApplicationService) GetApplication(appNameOrID string) (*Application, error) {

	url := fmt.Sprintf("controller/rest/applications/%s?output=json", appNameOrID)

	var apps []*Application
	err := s.client.Rest("GET", url, &apps, nil)
	if err != nil {
		return nil, err
	}

	return apps[0], nil
}

// GetApplicationsAllTypes is a RESTUI call.
// It might break in future versions of AppDynamics
func (s *ApplicationService) GetApplicationsAllTypes() ([]*Application, error) {

	url := fmt.Sprintf("controller/restui/applicationManagerUiBean/getApplicationsAllTypes")

	var apps applicationAllTypes
	err := s.client.RestInternal("GET", url, &apps, nil)
	if err != nil {
		return nil, err
	}

	return apps.Applications, nil

}

// ExportApplicationConfig will export an Application to the io.Writer specified
func (s *ApplicationService) ExportApplicationConfig(appID int) ([]byte, error) {
	url := fmt.Sprintf("controller/ConfigObjectImportExportServlet?applicationId=%d", appID)

	body, err := s.client.DoRawRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// DANGER ZONE
// this is an UNPUBLISHED API call - it may change in the future
func (s *ApplicationService) GetAllInternalApplications() (*AllInternalApplications, error) {

	url := "/controller/restui/applicationManagerUiBean/getApplicationsAllTypes"

	apps := AllInternalApplications{}
	err := s.client.Rest("GET", url, &apps, nil)
	if err != nil {
		return nil, err
	}

	return &apps, nil
}

// DANGER ZONE END
