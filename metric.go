/*
MIT License

Copyright (c) 2023 David Lopes
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
// Added 2023 Cisco Systems, Inc.
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
