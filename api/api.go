package api

type Nodes []*Node

type Node struct {
	Name       string    `json:"name"`
	DataCenter string    `json:"dataCenter"`
	IP         string    `json:"ip"`
	Port       int       `json:"port,omitempty"`
	Version    int       `json:"version"`
	State      StateType `json:"state"`
	SnapShot   string    `json:"snapShot,omitempty"`
	Services   `json:"services,omitempty"`
}

type Services []*Service

type Service struct {
	Name  string    `json:"name"`
	IP    string    `json:"ip,omitempty"`
	Port  int       `json:"port,omitempty"`
	State StateType `json:"state,omitempty"`
	Nodes `json:"nodes,omitempty"`
}