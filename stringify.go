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
