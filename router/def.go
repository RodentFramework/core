package router

// Router Defining Struct
type Router struct {
	ID        int
	RouterID  string      `json:"Rid"` //
	Name      string      `json:"Name"`
	Token     string      `json:"Token"`
	Online    bool        `json:"Status"`
	CommsType string      `json:"Type"`
	Commands  []RouterCmd `json:"Cmds"`
}

type RouterCmd struct {
	Name      string `json:"Name"`
	Helptext  string `json:"Help"`
	Usagetext string `json:"Usage"`
}

type Config struct {
	Reguris         []string `json:"registeruris"`
	Beaconuris      []string `json:"beaconuris"`
	Outputuris      []string `json:"outputuris"`
	Address         string   `json:"address"`
	Protocol        string   `json:"protocol"`
	Verbosity       bool     `json:"verbosity"`
	VerifyTls       bool     `json:"tls"`
	Cert            string   `json:"tlscert"`
	Certkey         string   `json:"tlscertkey"`
	CertCommonNames []string `json:"commonNames"`
	LogPath         string   `json:"logpath"`
}
