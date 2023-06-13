package appdrest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
)

// Dashboard represents a single Dashboard within AppDynamics
type Dashboard struct {
	ID                        int           `json:"id"`
	Version                   int           `json:"version"`
	Name                      string        `json:"name"`
	NameUnique                bool          `json:"nameUnique"`
	BuiltIn                   bool          `json:"builtIn"`
	CreatedBy                 string        `json:"createdBy"`
	CreatedOn                 int64         `json:"createdOn"`
	ModifiedBy                string        `json:"modifiedBy"`
	ModifiedOn                int64         `json:"modifiedOn"`
	Description               interface{}   `json:"description"`
	MissingAssociatedEntities interface{}   `json:"missingAssociatedEntities"`
	Widgets                   []Widget      `json:"widgets"`
	SecurityToken             interface{}   `json:"securityToken"`
	SharingRevoked            bool          `json:"sharingRevoked"`
	WarRoom                   bool          `json:"warRoom"`
	Template                  bool          `json:"template"`
	TemplateEntityType        string        `json:"templateEntityType"`
	MinutesBeforeAnchorTime   int           `json:"minutesBeforeAnchorTime"`
	StartTime                 int           `json:"startTime"`
	EndTime                   int           `json:"endTime"`
	RefreshInterval           int           `json:"refreshInterval"`
	BackgroundColor           int           `json:"backgroundColor"`
	Color                     int           `json:"color"`
	Height                    int           `json:"height"`
	Width                     int           `json:"width"`
	Disabled                  bool          `json:"disabled"`
	CanvasType                string        `json:"canvasType"`
	LayoutType                string        `json:"layoutType"`
	Properties                []interface{} `json:"properties"`
}

// Widget - One Dashboard contains multiple Widgets
type Widget struct {
	Type                        string        `json:"type"`
	ID                          int           `json:"id"`
	Version                     int           `json:"version"`
	GUID                        string        `json:"guid"`
	Title                       string        `json:"title"`
	DashboardID                 int           `json:"dashboardId"`
	WidgetsMetricMatchCriterias interface{}   `json:"widgetsMetricMatchCriterias"`
	Height                      int           `json:"height"`
	Width                       int           `json:"width"`
	MinHeight                   int           `json:"minHeight"`
	MinWidth                    int           `json:"minWidth"`
	X                           int           `json:"x"`
	Y                           int           `json:"y"`
	Label                       string        `json:"label"`
	Description                 string        `json:"description"`
	DrillDownURL                string        `json:"drillDownUrl"`
	UseMetricBrowserAsDrillDown bool          `json:"useMetricBrowserAsDrillDown"`
	DrillDownActionType         interface{}   `json:"drillDownActionType"`
	BackgroundColor             int           `json:"backgroundColor"`
	Color                       int           `json:"color"`
	FontSize                    int           `json:"fontSize"`
	UseAutomaticFontSize        bool          `json:"useAutomaticFontSize"`
	BorderEnabled               bool          `json:"borderEnabled"`
	BorderThickness             int           `json:"borderThickness"`
	BorderColor                 int           `json:"borderColor"`
	BackgroundAlpha             float64       `json:"backgroundAlpha"`
	ShowValues                  bool          `json:"showValues"`
	BackgroundColors            interface{}   `json:"backgroundColors"`
	CompactMode                 bool          `json:"compactMode"`
	ShowTimeRange               bool          `json:"showTimeRange"`
	RenderIn3D                  bool          `json:"renderIn3D"`
	ShowLegend                  bool          `json:"showLegend"`
	LegendPosition              interface{}   `json:"legendPosition"`
	LegendColumnCount           interface{}   `json:"legendColumnCount"`
	StartTime                   interface{}   `json:"startTime"`
	EndTime                     interface{}   `json:"endTime"`
	MinutesBeforeAnchorTime     int           `json:"minutesBeforeAnchorTime"`
	IsGlobal                    bool          `json:"isGlobal"`
	Properties                  []interface{} `json:"properties"`
	MissingEntities             interface{}   `json:"missingEntities"`
	AdqlQueries                 []string      `json:"adqlQueries"`
	AnalyticsType               string        `json:"analyticsType"`
	SearchMode                  string        `json:"searchMode"`
	IsStackingEnabled           bool          `json:"isStackingEnabled"`
	LegendsLayout               string        `json:"legendsLayout"`
	MaxAllowedYAxisFields       int           `json:"maxAllowedYAxisFields"`
	MaxAllowedXAxisFields       int           `json:"maxAllowedXAxisFields"`
	Min                         interface{}   `json:"min"`
	Max                         interface{}   `json:"max"`
	MinType                     interface{}   `json:"minType"`
	MaxType                     interface{}   `json:"maxType"`
	ShowMinExtremes             bool          `json:"showMinExtremes"`
	ShowMaxExtremes             bool          `json:"showMaxExtremes"`
	IntervalType                interface{}   `json:"intervalType"`
	Interval                    interface{}   `json:"interval"`
	DisplayPercentileMarkers    bool          `json:"displayPercentileMarkers"`
	PercentileValue1            interface{}   `json:"percentileValue1"`
	PercentileValue2            interface{}   `json:"percentileValue2"`
	PercentileValue3            interface{}   `json:"percentileValue3"`
	PercentileValue4            interface{}   `json:"percentileValue4"`
	Resolution                  interface{}   `json:"resolution"`
	DataFetchSize               interface{}   `json:"dataFetchSize"`
	PercentileLine              interface{}   `json:"percentileLine"`
	TimeRangeInterval           interface{}   `json:"timeRangeInterval"`
	PollingInterval             interface{}   `json:"pollingInterval"`
	Unit                        int           `json:"unit"`
	IsRawQuery                  bool          `json:"isRawQuery"`
	ViewState                   interface{}   `json:"viewState"`
	GridState                   interface{}   `json:"gridState"`
}

// DataSeriesTemplate - described data series for timeseries widgets
type DataSeriesTemplate struct {
	SeriesType                  *string     `json:"seriesType"`
	MetricType                  *string     `json:"metricType"`
	ShowRawMetricName           *bool       `json:"showRawMetricName"`
	ColorPalette                interface{} `json:"colorPalette"`
	Name                        *string     `json:"name"`
	MetricMatchCriteriaTemplate struct {
		EntityMatchCriteria      *string `json:"entityMatchCriteria"`
		MetricExpressionTemplate struct {
			MetricExpressionType *string `json:"metricExpressionType"`
			FunctionType         *string `json:"functionType"`
			DisplayName          *string `json:"displayName"`
			InputMetricText      *bool   `json:"inputMetricText"`
			InputMetricPath      *string `json:"inputMetricPath"`
			MetricPath           *string `json:"metricPath"`
			ScopeEntity          struct {
				ApplicationName   *string `json:"applicationName"`
				EntityType        *string `json:"entityType"`
				EntityName        *string `json:"entityName"`
				ScopingEntityType *string `json:"scopingEntityType"`
				ScopingEntityName *string `json:"scopingEntityName"`
				Subtype           *string `json:"subtype"`
			} `json:"scopeEntity"`
		} `json:"metricExpressionTemplate"`
		RollupMetricData              *bool   `json:"rollupMetricData"`
		ExpressionString              *string `json:"expressionString"`
		UseActiveBaseline             *bool   `json:"useActiveBaseline"`
		SortResultsAscending          *bool   `json:"sortResultsAscending"`
		MaxResults                    *int    `json:"maxResults"`
		EvaluationScopeType           *string `json:"evaluationScopeType"`
		BaselineName                  *string `json:"baselineName"`
		ApplicationName               *string `json:"applicationName"`
		MetricDisplayNameStyle        *string `json:"metricDisplayNameStyle"`
		MetricDisplayNameCustomFormat *string `json:"metricDisplayNameCustomFormat"`
	} `json:"metricMatchCriteriaTemplate"`
	AxisPosition *string `json:"axisPosition"`
}

// WidgetExported - single widget as in dashboard export
type WidgetExported struct {
	WidgetType                  *string               `json:"widgetType"`
	Title                       *string               `json:"title"`
	Height                      *int                  `json:"height"`
	Width                       *int                  `json:"width"`
	MinHeight                   *int                  `json:"minHeight"`
	MinWidth                    *int                  `json:"minWidth"`
	X                           *int                  `json:"x"`
	Y                           *int                  `json:"y"`
	Label                       *string               `json:"label"`
	Description                 *string               `json:"description"`
	DrillDownURL                *string               `json:"drillDownUrl"`
	UseMetricBrowserAsDrillDown *bool                 `json:"useMetricBrowserAsDrillDown"`
	DrillDownActionType         *string               `json:"drillDownActionType"`
	BackgroundColor             *int                  `json:"backgroundColor"`
	BackgroundColors            *string               `json:"backgroundColors"`
	BackgroundColorsStr         *string               `json:"backgroundColorsStr"`
	Color                       *int                  `json:"color"`
	FontSize                    *int                  `json:"fontSize"`
	UseAutomaticFontSize        *bool                 `json:"useAutomaticFontSize"`
	BorderEnabled               *bool                 `json:"borderEnabled"`
	BorderThickness             *int                  `json:"borderThickness"`
	BorderColor                 *int                  `json:"borderColor"`
	BackgroundAlpha             *float64              `json:"backgroundAlpha"`
	ShowValues                  *bool                 `json:"showValues"`
	FormatNumber                *string               `json:"formatNumber"`
	NumDecimals                 *int                  `json:"numDecimals"`
	RemoveZeros                 *string               `json:"removeZeros"`
	CompactMode                 *bool                 `json:"compactMode"`
	ShowTimeRange               *bool                 `json:"showTimeRange"`
	RenderIn3D                  *bool                 `json:"renderIn3D"`
	ShowLegend                  *bool                 `json:"showLegend"`
	LegendPosition              *string               `json:"legendPosition"`
	LegendColumnCount           *int                  `json:"legendColumnCount"`
	StartTime                   *string               `json:"startTime"`
	EndTime                     *string               `json:"endTime"`
	MinutesBeforeAnchorTime     *int                  `json:"minutesBeforeAnchorTime"`
	IsGlobal                    *bool                 `json:"isGlobal"`
	PropertiesMap               interface{}           `json:"propertiesMap"`
	DataSeriesTemplates         []*DataSeriesTemplate `json:"dataSeriesTemplates"`
	SourceURL                   *string               `json:"sourceURL"`
	Sandbox                     *bool                 `json:"sandbox"`
	VerticalAxisLabel           *string               `json:"verticalAxisLabel"`
	HideHorizontalAxis          *string               `json:"hideHorizontalAxis"`
	HorizontalAxisLabel         *string               `json:"horizontalAxisLabel"`
	AxisType                    *string               `json:"axisType"`
	StackMode                   *bool                 `json:"stackMode"`
	MultipleYAxis               *bool                 `json:"multipleYAxis"`
	CustomVerticalAxisMin       *int                  `json:"customVerticalAxisMin"`
	CustomVerticalAxisMax       *int                  `json:"customVerticalAxisMax"`
	ShowEvents                  *bool                 `json:"showEvents"`
	InterpolateDataGaps         *bool                 `json:"interpolateDataGaps"`
	ShowAllTooltips             *bool                 `json:"showAllTooltips"`
	StaticThresholdList         interface{}           `json:"staticThresholdList"`
	EventFilterTemplate         *string               `json:"eventFilterTemplate"`
	Text                        *string               `json:"text"`
	TextAlign                   *string               `json:"textAlign"`
	Margin                      *int                  `json:"margin"`
	ImageURL                    *string               `json:"imageURL"`
}

// AssociatedEntityTemplate - associated entity type in exported dashboard
type AssociatedEntityTemplate struct {
	ApplicationName   *string `json:"applicationName"`
	EntityType        *string `json:"entityType"`
	EntityName        *string `json:"entityName"`
	ScopingEntityType *string `json:"scopingEntityType"`
	ScopingEntityName *string `json:"scopingEntityName"`
	Subtype           *string `json:"subtype"`
}

// DashboardExport represents a single Dashboard within AppDynamics
// as exported by Export function in json format
type DashboardExport struct {
	Name                      *string                     `json:"name"`
	Description               interface{}                 `json:"description"`
	DashboardFormatVersion    *string                     `json:"dashboardFormatVersion"`
	SchemaVersion             *string                     `json:"schemaVersion"`
	WidgetTemplates           []WidgetExported            `json:"widgetTemplates"`
	WarRoom                   *bool                       `json:"warRoom"`
	Template                  *bool                       `json:"template"`
	TemplateEntityType        *string                     `json:"templateEntityType"`
	AssociatedEntityTemplates []*AssociatedEntityTemplate `json:"associatedEntityTemplates"`
	MinutesBeforeAnchorTime   *int                        `json:"minutesBeforeAnchorTime"`
	StartTime                 *int                        `json:"startTime"`
	EndTime                   *int                        `json:"endTime"`
	RefreshInterval           *int                        `json:"refreshInterval"`
	BackgroundColor           *int                        `json:"backgroundColor"`
	Color                     *int                        `json:"color"`
	Height                    *int                        `json:"height"`
	Width                     *int                        `json:"width"`
	CanvasType                *string                     `json:"canvasType"`
	LayoutType                *string                     `json:"layoutType"`
	Properties                []interface{}               `json:"properties"`
}

type DashboardUploadResponse struct {
	Success              bool          `json:"success"`
	Errors               []interface{} `json:"errors"`
	Warnings             []interface{} `json:"warnings"`
	Dashboard            Dashboard     `json:"dashboard"`
	CreatedDashboardName string        `json:"createdDashboardName"`
}

// DashboardService intermediates Dashboard requests
type DashboardService service

// GetDashboards obtains all dashboards from a controller
func (s *DashboardService) GetDashboards() ([]*Dashboard, error) {

	url := "/controller/restui/dashboards/getAllDashboardsByType/false"
	var dashboards []*Dashboard
	err := s.client.RestInternal("GET", url, &dashboards, nil)
	if err != nil {
		return nil, err
	}

	return dashboards, nil
}

// GetDashboard obtains a single dashboard from a controller
func (s *DashboardService) GetDashboard(ID int) (*Dashboard, error) {

	url := fmt.Sprintf("/controller/restui/dashboards/dashboardIfUpdated/%d/-1", ID)

	var dashboard *Dashboard
	err := s.client.RestInternal("GET", url, &dashboard, nil)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}

// TODO: Implement UpdateWidget
// UpdateWidget updates a Widget of a Dashboard ID
// func (s *DashboardService) UpdateWidget(dashboardID int, widget *Widget) {
// /controller/restui/dashboards/updateWidget
// }

// GetDashboardListForTier - get list of dashboards for a given application tier
func (s *DashboardService) GetDashboardListForTier(tierID int) ([]*Dashboard, error) {

	url := fmt.Sprintf("/controller/restui/templates/getAllDashboardTemplatesByTier/%d?isTierDashboard=true", tierID)

	var dashboards []*Dashboard
	err := s.client.RestInternal("GET", url, &dashboards, nil)
	if err != nil {
		return nil, err
	}

	return dashboards, nil
}

// DeleteDashboard - delete a dashboards by Id
func (s *DashboardService) DeleteDashboard(tierID int) error {

	url := "/controller/restui/dashboards/deleteDashboards"

	body := []int{tierID}
	err := s.client.RestInternal("POST", url, nil, &body)
	if err != nil {
		return err
	}

	return nil
}

// GetDashboardExport - get dashboard in export/import format
func (s *DashboardService) GetDashboardExport(dashboardID int) (*DashboardExport, error) {

	url := fmt.Sprintf("/controller/CustomDashboardImportExportServlet?dashboardId=%d", dashboardID)

	var dashboard *DashboardExport
	err := s.client.RestInternal("GET", url, &dashboard, nil)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}

// UploadDashboardExport - upload a new dashboard in export/import format
func (s *DashboardService) UploadDashboardExport(dashboard *DashboardExport) (*DashboardUploadResponse, error) {

	dashboardJSON, err := json.Marshal(dashboard)

	url := "/controller/CustomDashboardImportExportServlet"

	multipartPayload := &bytes.Buffer{}
	writer := multipart.NewWriter(multipartPayload)
	part, err := writer.CreateFormFile("file", "dashboard.json")
	if err != nil {
		return nil, err
	}
	part.Write(dashboardJSON)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = writer.FormDataContentType()

	retval := DashboardUploadResponse{}

	// fmt.Printf("Upload body: %s\n", multipartPayload)

	err = s.client.RestInternalHdr("POST", url, &retval, multipartPayload, headers)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("Upload retval: %s\n", retval)

	return &retval, nil
}
