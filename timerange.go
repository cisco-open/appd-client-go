package appdrest

import (
	"fmt"
)

// TimeRange represents a TimeRange on the controller
type TimeRange struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	Name        string `json:"name"`
	NameUnique  bool   `json:"nameUnique"`
	BuiltIn     bool   `json:"builtIn"`
	CreatedBy   string `json:"createdBy"`
	CreatedOn   int64  `json:"createdOn"`
	ModifiedBy  string `json:"modifiedBy"`
	ModifiedOn  int64  `json:"modifiedOn"`
	Description string `json:"description"`
	TimeRange   struct {
		Type              string `json:"type"`
		DurationInMinutes int    `json:"durationInMinutes"`
		StartTime         int64  `json:"startTime"`
		EndTime           int64  `json:"endTime"`
		TimeRange         struct {
			StartTime int64 `json:"startTime"`
			EndTime   int64 `json:"endTime"`
		} `json:"timeRange"`
		TimeRangeAdjusted bool `json:"timeRangeAdjusted"`
	} `json:"timeRange"`
	Shared     bool `json:"shared"`
	Modifiable bool `json:"modifiable"`
}

// TimeRangeService intermediates TimeRange operations
type TimeRangeService service

// GetTimeRanges will return an array with all time ranges on the controller
func (s *TimeRangeService) GetTimeRanges() ([]*TimeRange, error) {

	url := "controller/restui/user/getAllCustomTimeRanges"

	var timeRanges []*TimeRange
	err := s.client.RestInternal("GET", url, &timeRanges, nil)
	if err != nil {
		return nil, err
	}

	return timeRanges, nil
}

// GetTimeRangeByName is a helper function that gets all time ranges
// But only returns the one that matches the name
func (s *TimeRangeService) GetTimeRangeByName(name string) (*TimeRange, error) {
	timeRanges, err := s.GetTimeRanges()
	if err != nil {
		return nil, err
	}

	for _, timeRange := range timeRanges {
		if timeRange.Name == name {
			return timeRange, nil
		}
	}

	err = fmt.Errorf("Could not find Time Range with name: %s ", name)
	return nil, err

}

// UpdateTimeRange will update an existing Time Range
func (s *TimeRangeService) UpdateTimeRange(tr TimeRange) (*TimeRange, error) {

	url := "controller/restui/user/updateCustomRange"

	var returnTr *TimeRange
	err := s.client.RestInternal("POST", url, &returnTr, &tr)
	if err != nil {
		return nil, err
	}

	return returnTr, nil
}
