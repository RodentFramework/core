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
