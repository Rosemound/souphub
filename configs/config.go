package configs

import (
	"encoding/json"
	"os"

	"github.com/rosemound/souphub/internal/domain/models"
)


type Config struct {
	// Name is a hub instance name (uniq for one AccessToken)
	Name string `json:"name"`

	Company *models.Company `json:"company,omitempty"`

	Description string `json:"description,omitempty"`

	// Environment: dev, local, prod
	Environment string `json:"environment"`

	HttpPort string `json:"port"`

	// AccessToken for income requests
	AccessToken string `json:"accessToken"`

	Servers models.GameServers `json:"servers,omitempty"`
}

func (c *Config) GetServers() models.GameServers {
	return c.Servers
}

func (c *Config) GetName() string {
	return c.Name
}

func (c *Config) GetAccessToken() string {
	return c.AccessToken
}

func (c *Config) IsProd() bool {
	return isProd(c.Environment)
}

func (c *Config) IsDebug() bool {
	return isDebug(c.Environment)
}

func Get(buf *Config) error {
	bytes, err := os.ReadFile("souph.json")

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, buf)
}

func isDebug(v string) bool {
	return v == "" || v == "dev" || v == "local"
}

func isProd(v string) bool {
	return v == "prod"
}
