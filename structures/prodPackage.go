package structures

type ProdPackage struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class      string `json:"_class,omitempty"`
		Parameters []struct {
			Class   string `json:"_class"`
			Name    string `json:"name"`
			JobName string `json:"jobName"`
			Number  string `json:"number"`
		} `json:"parameters,omitempty"`
		Causes []struct {
			Class            string `json:"_class"`
			ShortDescription string `json:"shortDescription"`
			UserName         string `json:"userName"`
		} `json:"causes,omitempty"`
		OneClickDeployPossible  bool    `json:"oneClickDeployPossible,omitempty"`
		OneClickDeployReady     bool    `json:"oneClickDeployReady,omitempty"`
		OneClickDeployValid     bool    `json:"oneClickDeployValid,omitempty"`
		BlockedDurationMillis   int     `json:"blockedDurationMillis,omitempty"`
		BlockedTimeMillis       int     `json:"blockedTimeMillis,omitempty"`
		BuildableDurationMillis int     `json:"buildableDurationMillis,omitempty"`
		BuildableTimeMillis     int     `json:"buildableTimeMillis,omitempty"`
		BuildingDurationMillis  int     `json:"buildingDurationMillis,omitempty"`
		ExecutingTimeMillis     int     `json:"executingTimeMillis,omitempty"`
		ExecutorUtilization     float64 `json:"executorUtilization,omitempty"`
		SubTaskCount            int     `json:"subTaskCount,omitempty"`
		WaitingDurationMillis   int     `json:"waitingDurationMillis,omitempty"`
		WaitingTimeMillis       int     `json:"waitingTimeMillis,omitempty"`
	} `json:"actions"`
	Artifacts         []interface{} `json:"artifacts"`
	Building          bool          `json:"building"`
	Description       interface{}   `json:"description"`
	DisplayName       string        `json:"displayName"`
	Duration          int           `json:"duration"`
	EstimatedDuration int           `json:"estimatedDuration"`
	Executor          interface{}   `json:"executor"`
	FullDisplayName   string        `json:"fullDisplayName"`
	ID                string        `json:"id"`
	KeepLog           bool          `json:"keepLog"`
	Number            int           `json:"number"`
	QueueID           int           `json:"queueId"`
	Result            string        `json:"result"`
	Timestamp         int64         `json:"timestamp"`
	URL               string        `json:"url"`
	BuiltOn           string        `json:"builtOn"`
	ChangeSet         struct {
		Class string        `json:"_class"`
		Items []interface{} `json:"items"`
		Kind  interface{}   `json:"kind"`
	} `json:"changeSet"`
	Culprits []interface{} `json:"culprits"`
	Target   struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"target"`
}
