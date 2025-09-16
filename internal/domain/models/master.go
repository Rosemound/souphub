package models

type Master struct {
	// User api token provided on master hub
	Name string `json:"name"`

	// Server addrs available
	Addrs []string `json:"servers"`
}

func (m *Master) GetName() string {
	return m.Name
}

func (m *Master) GetAddrs() []string {
	return m.Addrs
}
