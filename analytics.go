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
