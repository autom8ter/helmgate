package config

import "helm.sh/helm/v3/pkg/repo"

type Config struct {
	Port       int64         `yaml:"port"`
	Debug      bool          `yaml:"debug"`
	RegoPolicy string        `yaml:"rego_policy"`
	RegoQuery  string        `yaml:"rego_query"`
	JwksURI    string        `yaml:"jwks_uri"`
	Repos      []*repo.Entry `yaml:"repos"`
}

func (c *Config) SetDefaults() {
	if c.Port == 0 {
		c.Port = 8820
	}
	if c.RegoPolicy == "" {
		c.RegoPolicy = `
		package helmgate.authz

		default allow = true
`
	}
	if c.RegoQuery == "" {
		c.RegoQuery = "data.helmgate.authz.allow"
	}
}
