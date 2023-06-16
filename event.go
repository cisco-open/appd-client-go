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
	"fmt"
	"net/url"
)

// EventService provides event API
type EventService service

type Event struct {
	AppIdOrName     string            `json:"appId"`
	Severity        string            `json:"severity"`
	Summary         string            `json:"summary"`
	Comment         string            `json:"comment"`
	CustomEventType string            `json:"customEventType"`
	Tier            string            `json:"tier"`
	Node            string            `json:"node"`
	Properties      map[string]string `json:"properties"`
}

// CreateEvent - creates event
func (s *EventService) CreateEvent(event *Event) error {

	evtUrl := fmt.Sprintf("controller/rest/applications/%s/events?", //customeventtype=mycustomevent&propertynames=key1&propertynames=key2&propertyvalues=value1&propertyvalues=value2",
		url.PathEscape(event.AppIdOrName),
	)

	params := url.Values{}
	params.Add("severity", event.Severity)
	params.Add("summary", event.Summary)
	params.Add("eventtype", "CUSTOM")
	if event.Comment != "" {
		params.Add("comment", event.Comment)
	}
	if event.Tier != "" {
		params.Add("tier", event.Tier)
	}
	if event.Node != "" {
		params.Add("node", event.Node)
	}
	if event.CustomEventType != "" {
		params.Add("customeventtype", event.CustomEventType)
	}
	if event.Properties != nil {
		for key, value := range event.Properties {
			params.Add("propertynames", key)
			params.Add("propertyvalues", value)
		}
	}
	evtUrl = evtUrl + params.Encode() + "&"

	err := s.client.RestInternal("POST", evtUrl, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
