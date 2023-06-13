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
