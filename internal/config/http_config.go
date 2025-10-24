package config

import (
	"errors"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

// HTTPConfig holds the configuration values for the HTTP server
type HTTPConfig struct {
	Host string
	Port string
}

// NewHTTPConfig creates a new HTTPConfig instance
func NewHTTPConfig() (*HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &HTTPConfig{host, port}, nil
}

// Address returns the full network address in the format "host:port"
func (cfg *HTTPConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
