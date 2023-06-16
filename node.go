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

import (
	"fmt"
)

// Node represents one node within one Application
type Node struct {
	AppAgentVersion     string      `json:"appAgentVersion"`
	MachineAgentVersion string      `json:"machineAgentVersion"`
	AgentType           string      `json:"agentType"`
	Type                string      `json:"type"`
	MachineName         string      `json:"machineName"`
	AppAgentPresent     bool        `json:"appAgentPresent"`
	NodeUniqueLocalID   string      `json:"nodeUniqueLocalId"`
	MachineID           int         `json:"machineId"`
	MachineOSType       string      `json:"machineOSType"`
	TierID              int         `json:"tierId"`
	TierName            string      `json:"tierName"`
	MachineAgentPresent bool        `json:"machineAgentPresent"`
	Name                string      `json:"name"`
	IPAddresses         interface{} `json:"ipAddresses"`
	ID                  int         `json:"id"`
}

// NodeService intermediates Node requests
type NodeService service

// GetNodes obtains all Nodes from an Application
func (s *NodeService) GetNodes(appIDOrName string) ([]*Node, error) {

	url := fmt.Sprintf("controller/rest/applications/%s/nodes?output=json", appIDOrName)

	var nodes []*Node
	err := s.client.Rest("GET", url, &nodes, nil)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// GetNode obtains a single Node from an Application
func (s *NodeService) GetNode(appIDOrName string, nodeNameOrID string) (*Node, error) {

	url := fmt.Sprintf("controller/rest/applications/%s/nodes/%s?output=json", appIDOrName, nodeNameOrID)

	var nodes []*Node
	err := s.client.Rest("GET", url, &nodes, nil)
	if err != nil {
		return nil, err
	}

	return nodes[0], nil
}
