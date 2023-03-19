package rodent

type Rodent struct {
	ID         int
	Xid        string     `json:"Xid"`
	Properties Properties `json:"Properties"`
	Config     Config     `json:"Config"`
	Type       string     `json:"Type"`
	LastSeen   string     `json:"Last"`
	Joined     string     `json:"Joined"`
	RouterID   string     `json:"Route"`
	Tasks      []Tasks    `json:"Tasks" `
	Reg        bool
}

type Properties struct {
	CodeName  string `json:"Name,omitempty"`
	User      string `json:"User,omitempty"`
	Uid       int    `json:"Uid,omitempty"`
	Host      string `json:"Host,omitempty"`
	Os        string `json:"Os,omitempty"`
	Ip        string `json:"Ip,omitempty"`
	RDomain   string `json:"Domain,omitempty"`
	ProcName  string `json:"Procname,omitempty"`
	Pid       int    `json:"Pid,omitempty"`
	Inspected bool   `json:"View"`
}

type Config struct {
	Beacon string `json:"Beacon,omitempty"`
	Jitter string `json:"Jitter,omitempty"`
	Key    string `json:"Key,omitempty"`
	Comms  Comms  `json:"Comms,omitempty"`
	Kdn    string `json:"kdn,omitempty"`
}

type Comms struct {
	Domains     []string `json:"Domains,omitempty"`
	HostHeaders []string `json:"Hosts,omitempty"`
	Cookie      []string `json:"Cookies,omitempty"`
	CheckPaths  []string `json:"TaskPaths,omitempty"`
	RegPath     string   `json:"RegPath,omitempty"`
	UserAgent   string   `json:"UserAgent,omitempty"`
}

type Task struct {
	Xid           string `json:"Xid"`
	Async         bool   `json:"Async,omitempty"`
	Num           int    `json:"TNum"`
	Type          string `json:"Type"`
	Uri           string `json:"TUrl"`
	Data          []byte `json:"Data,omitempty"`
	Result        []byte `json:"Result,omitempty"`
	Timer         string `json:"Timer,omitempty"`
	TimeCompleted string `json:"Completed,omitempty"`
	DateIssued    string `json:"DateIsh,omitempty"`
	Done          bool   `json:"Done"`
}

// Just a struct handy for passing multiple tasks to a rodent
type Tasks struct {
	Tasks []Task `json:"RTasks"`
}
