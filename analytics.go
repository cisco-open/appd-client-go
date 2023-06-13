package appdrest

// AnalyticsWidget is a widget inside an Analytics Search
type AnalyticsWidget struct {
	ID         int         `json:"id"`
	Version    int         `json:"version"`
	Name       string      `json:"name"`
	NameUnique bool        `json:"nameUnique"`
	ViewState  interface{} `json:"viewState"`
	Properties struct {
		MinSizeY        string `json:"minSizeY"`
		Col             string `json:"col"`
		SizeX           string `json:"sizeX"`
		BackgroundColor string `json:"backgroundColor"`
		MinSizeX        string `json:"minSizeX"`
		Color           string `json:"color"`
		LegendsLayout   string `json:"legendsLayout"`
		FontSize        string `json:"fontSize"`
		Row             string `json:"row"`
		Type            string `json:"type"`
		Title           string `json:"title"`
		SizeY           string `json:"sizeY"`
	} `json:"properties"`
	TimeRangeSpecifier struct {
		Type              string      `json:"type"`
		DurationInMinutes int         `json:"durationInMinutes"`
		StartTime         interface{} `json:"startTime"`
		EndTime           interface{} `json:"endTime"`
		TimeRange         struct {
			StartTime int64 `json:"startTime"`
			EndTime   int64 `json:"endTime"`
		} `json:"timeRange"`
		TimeRangeAdjusted bool `json:"timeRangeAdjusted"`
	} `json:"timeRangeSpecifier"`
	AdqlQueries []string `json:"adqlQueries"`
}

// AnalyticsSearch represents a Saved Analytics Search, as of 4.4.3 this can only be accessed through the RestUI
type AnalyticsSearch struct {
	ID                 int               `json:"id"`
	Version            int               `json:"version"`
	Name               string            `json:"name"`
	NameUnique         bool              `json:"nameUnique"`
	BuiltIn            bool              `json:"builtIn"`
	CreatedBy          string            `json:"createdBy"`
	CreatedOn          int64             `json:"createdOn"`
	ModifiedBy         string            `json:"modifiedBy"`
	ModifiedOn         int64             `json:"modifiedOn"`
	SearchName         string            `json:"searchName"`
	SearchDescription  interface{}       `json:"searchDescription"`
	SearchType         string            `json:"searchType"`
	SearchMode         string            `json:"searchMode"`
	ViewMode           string            `json:"viewMode"`
	Visualization      string            `json:"visualization"`
	SelectedFields     []string          `json:"selectedFields"`
	TimeRangeSpecifier interface{}       `json:"timeRangeSpecifier"`
	AdqlQueries        []string          `json:"adqlQueries"`
	Widgets            []AnalyticsWidget `json:"widgets"`
	GridState          interface{}       `json:"gridState"`
}

// AnalyticsService intermediates Analytics Queries
type AnalyticsService service

// GetAnalyticsSearches obtains all Analytics Serches saved
func (s *AnalyticsService) GetAnalyticsSearches() ([]*AnalyticsSearch, error) {

	url := "controller/restui/analyticsSavedSearches/getAllAnalyticsSavedSearches"

	var analyticsSearches []*AnalyticsSearch
	err := s.client.RestInternal("GET", url, &analyticsSearches, nil)
	if err != nil {
		return nil, err
	}

	return analyticsSearches, nil
}
