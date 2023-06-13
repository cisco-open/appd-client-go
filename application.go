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
