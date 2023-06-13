package appdrest

import (
	"fmt"
	"time"
)

// MetricData contains metric values for a single metric
type MetricData struct {
	MetricName   string        `json:"metricName"`
	MetricID     int           `json:"metricId"`
	MetricPath   string        `json:"metricPath"`
	Frequency    string        `json:"frequency"`
	MetricValues []MetricValue `json:"metricValues"`
}

// MetricValue is always part of an array of metrics, inside a MetricData struct
type MetricValue struct {
	Occurrences       int   `json:"occurrences"`
	Current           int   `json:"current"`
	Min               int   `json:"min"`
	Max               int   `json:"max"`
	StartTimeInMillis int64 `json:"startTimeInMillis"`
	UseRange          bool  `json:"useRange"`
	Count             int   `json:"count"`
	Sum               int   `json:"sum"`
	Value             int   `json:"value"`
	StandardDeviation int   `json:"standardDeviation"`
}

// Metric represents a Metric object that might be a folder or child
type Metric struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Consts for the technique used to obtain metric data
const (
	TimeBEFORENOW    = "BEFORE_NOW"
	TimeBEFORETIME   = "BEFORE_TIME"
	TimeAFTERTIME    = "AFTER_TIME"
	TimeBETWEENTIMES = "BETWEEN_TIMES"
)

// MetricDataService intermediates MetricData requests
type MetricDataService service

// GetMetricData obtains metrics matching a pattern
func (s *MetricDataService) GetMetricData(appIDOrName string, metricPath string, rollup bool, timeRangeType string, durationInMins int, startTime time.Time, endTime time.Time) ([]*MetricData, error) {

	url := fmt.Sprintf("controller/rest/applications/%v/metric-data?output=json", appIDOrName)
	url += fmt.Sprintf("&rollup=%t", rollup)

	url += fmt.Sprintf("&metric-path=%s", metricPath)
	url += fmt.Sprintf("&time-range-type=%s", timeRangeType)

	if timeRangeType == TimeBEFORENOW || timeRangeType == TimeBEFORETIME || timeRangeType == TimeAFTERTIME {
		url += fmt.Sprintf("&duration-in-mins=%d", durationInMins)

	}
	if timeRangeType == TimeAFTERTIME || timeRangeType == TimeBETWEENTIMES {
		url += fmt.Sprintf("&start-time=%v", startTime.UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)))
	}
	if timeRangeType == TimeBEFORETIME || timeRangeType == TimeBETWEENTIMES {
		url += fmt.Sprintf("&end-time=%v", endTime.UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)))
	}

	var metrics []*MetricData
	err := s.client.Rest("GET", url, &metrics, nil)
	if err != nil {
		return nil, fmt.Errorf("Metric API: %v -> %s", err, url)
	}

	return metrics, nil
}

// GetMetricHierarchy obtains the Metric Browser hierarchy
func (s *MetricDataService) GetMetricHierarchy(appIDOrName string, metricPath string) ([]*Metric, error) {
	url := fmt.Sprintf("controller/rest/applications/%v/metrics?output=json", appIDOrName)

	if metricPath != "" {
		url += fmt.Sprintf("&metric-path=%s", metricPath)
	}

	var metrics []*Metric
	err := s.client.Rest("GET", url, &metrics, nil)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}
