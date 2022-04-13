package structures

type Messages struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class  string `json:"_class,omitempty"`
		Causes []struct {
			Class            string `json:"_class"`
			ShortDescription string `json:"shortDescription"`
		} `json:"causes,omitempty"`
		Parameters []struct {
			Class   string `json:"_class"`
			Name    string `json:"name"`
			JobName string `json:"jobName"`
			Number  string `json:"number"`
		} `json:"parameters,omitempty"`
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
		BuildsByBranchName      struct {
			OriginMaster struct {
				Class       string      `json:"_class"`
				BuildNumber int         `json:"buildNumber"`
				BuildResult interface{} `json:"buildResult"`
				Marked      struct {
					Sha1   string `json:"SHA1"`
					Branch []struct {
						Sha1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"marked"`
				Revision struct {
					Sha1   string `json:"SHA1"`
					Branch []struct {
						Sha1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"revision"`
			} `json:"origin/master"`
			RefsRemotesOriginMaster struct {
				Class       string      `json:"_class"`
				BuildNumber int         `json:"buildNumber"`
				BuildResult interface{} `json:"buildResult"`
				Marked      struct {
					Sha1   string `json:"SHA1"`
					Branch []struct {
						Sha1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"marked"`
				Revision struct {
					Sha1   string `json:"SHA1"`
					Branch []struct {
						Sha1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"revision"`
			} `json:"refs/remotes/origin/master"`
		} `json:"buildsByBranchName,omitempty"`
		LastBuiltRevision struct {
			Sha1   string `json:"SHA1"`
			Branch []struct {
				Sha1 string `json:"SHA1"`
				Name string `json:"name"`
			} `json:"branch"`
		} `json:"lastBuiltRevision,omitempty"`
		RemoteUrls []string `json:"remoteUrls,omitempty"`
		ScmName    string   `json:"scmName,omitempty"`
	} `json:"actions"`
	Artifacts []struct {
		DisplayPath  string `json:"displayPath"`
		FileName     string `json:"fileName"`
		RelativePath string `json:"relativePath"`
	} `json:"artifacts"`
	Building          bool        `json:"building"`
	Description       string      `json:"description"`
	DisplayName       string      `json:"displayName"`
	Duration          int         `json:"duration"`
	EstimatedDuration int         `json:"estimatedDuration"`
	Executor          interface{} `json:"executor"`
	FullDisplayName   string      `json:"fullDisplayName"`
	ID                string      `json:"id"`
	KeepLog           bool        `json:"keepLog"`
	Number            int         `json:"number"`
	QueueID           int         `json:"queueId"`
	Result            string      `json:"result"`
	Timestamp         int64       `json:"timestamp"`
	URL               string      `json:"url"`
	BuiltOn           string      `json:"builtOn"`
	ChangeSet         struct {
		Class string `json:"_class"`
		Items []struct {
			Class         string   `json:"_class"`
			AffectedPaths []string `json:"affectedPaths"`
			CommitID      string   `json:"commitId"`
			Timestamp     int64    `json:"timestamp"`
			Author        struct {
				AbsoluteURL string `json:"absoluteUrl"`
				FullName    string `json:"fullName"`
			} `json:"author"`
			AuthorEmail string `json:"authorEmail"`
			Comment     string `json:"comment"`
			Date        string `json:"date"`
			ID          string `json:"id"`
			Msg         string `json:"msg"`
			Paths       []struct {
				EditType string `json:"editType"`
				File     string `json:"file"`
			} `json:"paths"`
		} `json:"items"`
		Kind string `json:"kind"`
	} `json:"changeSet"`
	Culprits []struct {
		AbsoluteURL string `json:"absoluteUrl"`
		FullName    string `json:"fullName"`
	} `json:"culprits"`
}
