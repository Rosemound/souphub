package models

type GameServerAddr string

type GameServers map[GameServerAddr]*GameServer

// hub is only received addrs of servers
type GameServer struct {
	Name     string `json:"name"`
	Category string `json:"category,omitempty"`
}
