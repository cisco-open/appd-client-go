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

import "fmt"

func (a Application) String() string {
	return fmt.Sprintf(`
	ID      : %d
	Name    : %s
	BuiltIn : %t
	Active  : %t
	Running : %t
	`, a.ID, a.Name, a.BuiltIn, a.Active, a.Running)
}

func (t Tier) String() string {
	return fmt.Sprintf(`
	ID: %d
	Name: %s
	NumberOfNodes: %d
	Type: %s
	AgentType: %s
	`, t.ID, t.Name, t.NumberOfNodes, t.Type, t.AgentType)
}

func (t HealthRuleDetail) String() string {
	return fmt.Sprintf(`
	Name: %s,
	Enabled: %t                
	UseDataFromLastNMinutes: %d 
	WaitTimeAfterViolation: %d 
	ScheduleName: %s
	Affects: %s
	EvalCriterias %s
	`, *t.Name, *t.Enabled, *t.UseDataFromLastNMinutes, *t.WaitTimeAfterViolation, *t.ScheduleName, *t.Affects, *t.EvalCriterias)
}
