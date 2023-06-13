package appdrest

import (
	"fmt"
	"strings"
	"time"
)

// ErrorDetail details the errors that are present on the Snapshot
type ErrorDetail struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Snapshot represents one Snapshot within one Application
type Snapshot struct {
	FirstInChain                   bool          `json:"firstInChain"`
	TransactionProperties          []interface{} `json:"transactionProperties"`
	ErrorIDs                       []interface{} `json:"errorIDs"`
	CriticalThreshold              string        `json:"criticalThreshold"`
	HTTPHeaders                    []interface{} `json:"httpHeaders"`
	ServerStartTime                int64         `json:"serverStartTime"`
	SessionKeys                    []interface{} `json:"sessionKeys"`
	EndToEndUserExperience         string        `json:"endToEndUserExperience"`
	TimeTakenInMilliSecs           int           `json:"timeTakenInMilliSecs"`
	CallChain                      string        `json:"callChain"`
	LocalStartTime                 int64         `json:"localStartTime"`
	ID                             int           `json:"id"`
	LogMessages                    []interface{} `json:"logMessages"`
	RequestGUID                    string        `json:"requestGUID"`
	HTTPSessionID                  string        `json:"httpSessionID"`
	HasDeepDiveData                bool          `json:"hasDeepDiveData"`
	WarningThreshold               string        `json:"warningThreshold"`
	DeepDivePolicy                 string        `json:"deepDivePolicy"`
	SnapshotExitCalls              []interface{} `json:"snapshotExitCalls"`
	ExitCallsDataTruncated         bool          `json:"exitCallsDataTruncated"`
	ResponseHeaders                []interface{} `json:"responseHeaders"`
	ApplicationComponentNodeID     int           `json:"applicationComponentNodeId"`
	ApplicationID                  int           `json:"applicationId"`
	ExitCallsDataTruncationMessage string        `json:"exitCallsDataTruncationMessage"`
	HTTPParameters                 []interface{} `json:"httpParameters"`
	ErrorDetails                   []ErrorDetail `json:"errorDetails"`
	ApplicationComponentID         int           `json:"applicationComponentId"`
	DelayedDeepDiveOffSet          int           `json:"delayedDeepDiveOffSet"`
	UserExperience                 string        `json:"userExperience"`
	URL                            string        `json:"URL"`
	BusinessTransactionID          int           `json:"businessTransactionId"`
	Cookies                        []interface{} `json:"cookies"`
	Archived                       bool          `json:"archived"`
	DiagnosticSessionGUID          string        `json:"diagnosticSessionGUID"`
	StallDump                      string        `json:"stallDump"`
	BusinessData                   []interface{} `json:"businessData"`
	SnapshotExitSequence           string        `json:"snapshotExitSequence"`
	TransactionEvents              []interface{} `json:"transactionEvents"`
	ErrorOccured                   bool          `json:"errorOccured"`
	UnresolvedCallInCallChain      bool          `json:"unresolvedCallInCallChain"`
	Summary                        string        `json:"summary"`
	DelayedDeepDive                bool          `json:"delayedDeepDive"`
	EndToEndLatency                int           `json:"endToEndLatency"`
	ErrorSummary                   string        `json:"errorSummary"`
	LocalID                        int           `json:"localID"`
	ThreadName                     string        `json:"threadName"`
	ThreadID                       string        `json:"threadID"`
	Async                          bool          `json:"async"`
	DotnetProperty                 []interface{} `json:"dotnetProperty"`
	CPUTimeTakenInMilliSecs        int           `json:"cpuTimeTakenInMilliSecs"`
	StackTraces                    []interface{} `json:"stackTraces"`
}

// Consts for the deepDivePolicy argument
const (
	DeepDivePolicySLAFAILURE                     = "SLA_FAILURE"
	DeepDivePolicyTIMESAMPLING                   = "TIME_SAMPLING"
	DeepDivePolicyERRORSAMPLING                  = "ERROR_SAMPLING"
	DeepDivePolicyOCCURRENCESAMPLING             = "OCCURRENCE_SAMPLING "
	DeepDivePolicyONDEMAND                       = "ON_DEMAND"
	DeepDivePolicyHOTSPOT                        = "HOTSPOT"
	DeepDivePolicyHOTSPOTLEARN                   = "HOTSPOT_LEARN"
	DeepDivePolicyAPPLICATIONSTARTUP             = "APPLICATION_STARTUP"
	DeepDivePolicySLOWDIAGNOSTICSESSION          = "SLOW_DIAGNOSTIC_SESSION"
	DeepDivePolicyERRORDIAGNOSTICSESSION         = "ERROR_DIAGNOSTIC_SESSION"
	DeepDivePolicyPOLICYFAILUREDIAGNOSTICSESSION = "POLICY_FAILURE_DIAGNOSTIC_SESSION"
	DeepDivePolicyDIAGNOSTICSESSION              = "DIAGNOSTIC_SESSION"
	DeepDivePolicyINFLIGHTSLOWSESSION            = "INFLIGHT_SLOW_SESSION"
)

// Consts to the userExperience argument
const (
	UserExperienceNORMAL   = "NORMAL"
	UserExperienceSLOW     = "SLOW"
	UserExperienceVERYSLOW = "VERYSLOW"
	UserExperienceSTALL    = "STALL"
	UserExperienceERROR    = "ERROR"
)

// SnapshotService intermediates Snapshot requests
type SnapshotService service

// SnapshotFilters can be used to obtain a list of snapshots
type SnapshotFilters struct {
	Guids                       []string // Array of comma-separated guids for the transaction snapshots. If not specified, retrieves all snapshots in the specified time range
	Archived                    bool     // True to retrieve archived snapshots. Default is false.
	DeepDivePolicy              []string // Array of comma-separated snapshot policy filters to apply.
	ApplicationComponentIds     []int    // Array of comma-separated tier IDs to filters. Default is all the tiers in the application
	ApplicationComponentNodeIds []int    // Array of comma-separated node ID filters. Default is all the nodes in the application
	BusinessTransactionIds      []int    // Array of comma-separated business transaction ID filters. Default is all the business transactions in the application.
	UserExperience              []string // Array of comma-separated user experiences filters
	FirstInChain                bool     // If true, retrieve only the first request from the chain. Default is false.
	NeedProps                   bool     // If true, the values of the following snapshot properties are included in the output. These values correspond to the values of the data-collector-type parameter. If false, the default, these values are empty in the output.
	NeedExitCalls               bool     // If true, exit calls are included in the result. Default is false.
	ExecutionTimeInMilis        int      // If set, retrieves only data for requests with execution times greater than this value.
	SessionID                   string   // If set, retrieves data only for this session id.
	UserPrincipalID             string   // If set, retrieves data only for this user login.
	ErrorIDs                    []int    // Array of comma-separated error codes to filter by. Default is to retrieve all error codes.
	StartingRequestID           string   // If set, retrieves data only for this range of request IDs.
	EndingRequestID             string   // If set, retrieves data only for this range of request IDs.
	ErrorOccurred               bool     // If true, retrieves only error requests. Default is false.
	DiagnosticSnapshot          bool     // If true, retrieves only diagnostic snapshots. Default is false.
	BadRequest                  bool     // If true, retrieves only slow and error requests. Default is false.
	DiagnosticSessionGUID       []string // Array of comma-separated diagnostic session guids to filters.
	DataCollectorName           string   // Used with data-collector-value to filter snapshot collection based on the value of a data collector.
	DataCollectorValue          string   // Used with data-collector-name to filter snapshot collection based on the value of a data collector.
	DataCollectorType           string   // Used with data-collector-name and data-collector-value to filter snapshot collection based on the value of a data collector. Some of the values contain spaces. All are case-sensitive and where indicated the spaces are required.
	MaximumResults              int      // A number, if specified, this number of maximum results will be returned. If not specified, default 600 results can be returned at most.
}

// GetSnapshots can be called by passing an AppID and a SnapshotFilters struct
func (s *SnapshotService) GetSnapshots(appID int, timeRangeType string, durationInMins int, startTime time.Time, endTime time.Time, filters *SnapshotFilters) ([]*Snapshot, error) {
	return s.getSnapshotsParams(appID, timeRangeType,
		durationInMins,
		startTime,
		endTime,
		filters.Guids,
		filters.Archived,
		filters.DeepDivePolicy,
		filters.ApplicationComponentIds,
		filters.ApplicationComponentNodeIds,
		filters.BusinessTransactionIds,
		filters.UserExperience,
		filters.FirstInChain,
		filters.NeedProps,
		filters.NeedExitCalls,
		filters.ExecutionTimeInMilis,
		filters.SessionID,
		filters.UserPrincipalID,
		filters.ErrorIDs,
		filters.StartingRequestID,
		filters.EndingRequestID,
		filters.ErrorOccurred,
		filters.DiagnosticSnapshot,
		filters.BadRequest,
		filters.DiagnosticSessionGUID,
		filters.DataCollectorName,
		filters.DataCollectorValue,
		filters.DataCollectorType,
		filters.MaximumResults)
}

// GetSnapshotsParams obtains all Snapshots for a timerange
func (s *SnapshotService) getSnapshotsParams(appID int, // Provide either the application name or application id.
	timeRangeType string, // Consts TimeBEFORENOW, TimeBEFORETIME, TimeAFTERTIME, TimeBETWEENTIMES
	durationInMins int, // Duration (in minutes) to return the data.
	startTime time.Time, // Start time (in milliseconds) from which the data is returned.
	endTime time.Time, // End time (in milliseconds) until which the data is returned.
	guids []string, // Array of comma-separated guids for the transaction snapshots. If not specified, retrieves all snapshots in the specified time range
	archived bool, // True to retrieve archived snapshots. Default is false.
	deepDivePolicy []string, // Array of comma-separated snapshot policy filters to apply.
	applicationComponentIds []int, // Array of comma-separated tier IDs to filters. Default is all the tiers in the application
	applicationComponentNodeIds []int, // Array of comma-separated node ID filters. Default is all the nodes in the application
	businessTransactionIds []int, // Array of comma-separated business transaction ID filters. Default is all the business transactions in the application.
	userExperience []string, // Array of comma-separated user experiences filters
	firstInChain bool, // If true, retrieve only the first request from the chain. Default is false.
	needProps bool, // If true, the values of the following snapshot properties are included in the output. These values correspond to the values of the data-collector-type parameter. If false, the default, these values are empty in the output.
	needExitCalls bool, // If true, exit calls are included in the result. Default is false.
	executionTimeInMilis int, // If set, retrieves only data for requests with execution times greater than this value.
	sessionID string, // If set, retrieves data only for this session id.
	userPrincipalID string, // If set, retrieves data only for this user login.
	errorIDs []int, // Array of comma-separated error codes to filter by. Default is to retrieve all error codes.
	startingRequestID string, // If set, retrieves data only for this range of request IDs.
	endingRequestID string, // If set, retrieves data only for this range of request IDs.
	errorOccurred bool, // If true, retrieves only error requests. Default is false.
	diagnosticSnapshot bool, // If true, retrieves only diagnostic snapshots. Default is false.
	badRequest bool, // If true, retrieves only slow and error requests. Default is false.
	diagnosticSessionGUID []string, // Array of comma-separated diagnostic session guids to filters.
	dataCollectorName string, // Used with data-collector-value to filter snapshot collection based on the value of a data collector.
	dataCollectorValue string, // Used with data-collector-name to filter snapshot collection based on the value of a data collector.
	dataCollectorType string, // Used with data-collector-name and data-collector-value to filter snapshot collection based on the value of a data collector. Some of the values contain spaces. All are case-sensitive and where indicated the spaces are required.
	maximumResults int, // A number, if specified, this number of maximum results will be returned. If not specified, default 600 results can be returned at most.
) ([]*Snapshot, error) {

	url := fmt.Sprintf("controller/rest/applications/%d/request-snapshots?output=json", appID)

	url += fmt.Sprintf("&time-range-type=%s", timeRangeType)

	if timeRangeType == TimeBEFORENOW || timeRangeType == TimeBEFORETIME || timeRangeType == TimeAFTERTIME {
		url += fmt.Sprintf("&duration-in-mins=%d", durationInMins)

	}
	if timeRangeType == TimeAFTERTIME || timeRangeType == TimeBETWEENTIMES {
		url += fmt.Sprintf("&start-time=%v", startTime.Unix()*1000)
	}
	if timeRangeType == TimeBEFORETIME || timeRangeType == TimeBETWEENTIMES {
		url += fmt.Sprintf("&end-time=%v", endTime.Unix()*1000)
	}

	if len(guids) > 0 {
		url += fmt.Sprintf("&guids=%v", strings.Join(guids, ","))
	}

	url += fmt.Sprintf("&archived=%v", archived)

	if len(deepDivePolicy) > 0 {
		url += fmt.Sprintf("&deep-dive-policy=%v", strings.Join(deepDivePolicy, ","))
	}

	if len(applicationComponentIds) > 0 {
		url += fmt.Sprintf("&application-component-ids=%v", arrayToString(applicationComponentIds, ","))
	}

	if len(applicationComponentNodeIds) > 0 {
		url += fmt.Sprintf("&application-component-node-ids=%v", arrayToString(applicationComponentNodeIds, ","))
	}

	if len(businessTransactionIds) > 0 {
		url += fmt.Sprintf("&business-transaction-ids=%v", arrayToString(businessTransactionIds, ","))
	}

	if len(userExperience) > 0 {
		url += fmt.Sprintf("&user-experience=%v", strings.Join(userExperience, ","))
	}

	url += fmt.Sprintf("&first-in-chain=%v", firstInChain)
	url += fmt.Sprintf("&need-props=%v", needProps)
	url += fmt.Sprintf("&need-exit-calls=%v", needExitCalls)

	if executionTimeInMilis > 0 {
		url += fmt.Sprintf("&execution-time-in-milis=%v", executionTimeInMilis)
	}

	if sessionID != "" {
		url += fmt.Sprintf("&session-id=%v", sessionID)
	}

	if userPrincipalID != "" {
		url += fmt.Sprintf("&user-principal-id=%v", userPrincipalID)
	}

	if len(errorIDs) > 0 {
		url += fmt.Sprintf("&error-ids=%v", arrayToString(errorIDs, ","))
	}

	if startingRequestID != "" {
		url += fmt.Sprintf("&starting-request-id=%v", startingRequestID)
	}

	if endingRequestID != "" {
		url += fmt.Sprintf("&ending-request-id=%v", endingRequestID)
	}

	url += fmt.Sprintf("&error-occurred=%v", errorOccurred)
	url += fmt.Sprintf("&diagnostic-snapshot=%v", diagnosticSnapshot)
	url += fmt.Sprintf("&bad-request=%v", badRequest)

	if len(diagnosticSessionGUID) > 0 {
		url += fmt.Sprintf("&diagnostic-session-guid=%v", strings.Join(diagnosticSessionGUID, ","))
	}

	if dataCollectorName != "" {
		url += fmt.Sprintf("&data-collector-name=%v", dataCollectorName)
	}
	if dataCollectorValue != "" {
		url += fmt.Sprintf("&data-collector-value=%v", dataCollectorValue)
	}
	if dataCollectorType != "" {
		url += fmt.Sprintf("&data-collector-type=%v", dataCollectorType)
	}

	if maximumResults > 0 {
		url += fmt.Sprintf("&maximum-results=%v", maximumResults)
	}

	var snapshots []*Snapshot
	err := s.client.Rest("GET", url, &snapshots, nil)
	if err != nil {
		return nil, err
	}

	return snapshots, nil
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
