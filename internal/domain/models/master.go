package models

type MasterToken string

type Masters map[MasterToken]*Master

type Master struct {
	// User api token provided on master hub
	Name string `json:"name"`

	// Masters host
	Host string `json:"host"`

	// Masters share exp (for re-share call)
	Expiration int64 `json:"expiration"`

	// Server addrs available
	Addrs []string `json:"servers"`
}

func (m *Master) GetName() string {
	return m.Name
}

func (m *Master) GetAddrs() []string {
	return m.Addrs
}

func (m *Master) GetHost() string {
	return m.Host
}

func (m *Master) GetExpiration() int64 {
	return m.Expiration
}
