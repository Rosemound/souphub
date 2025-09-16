package models


// hub is only received addrs of servers
type GameServer struct {
	Name string `json:"name,omitempty"`
	Category string `json:"category"`
}