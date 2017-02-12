package server

import (
	"time"

	"github.com/10gen/mongo-go-driver/conn"
)

func newConfig(opts ...Option) *config {
	cfg := &config{
		dialer:            conn.Dial,
		heartbeatInterval: time.Duration(10) * time.Second,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// Option configures a server.
type Option func(*config)

type config struct {
	connOpts          []conn.Option
	dialer            conn.Dialer
	heartbeatInterval time.Duration
}

// ConnectionOptions configures server's connections.
func ConnectionOptions(opts ...conn.Option) Option {
	return func(c *config) {
		c.connOpts = opts
	}
}

// HeartbeatInterval configures a server's heartbeat interval.
func HeartbeatInterval(interval time.Duration) Option {
	return func(c *config) {
		c.heartbeatInterval = interval
	}
}

// ConnectionDialerOpt configures a server's connection dialer.
func ConnectionDialerOpt(dialer conn.Dialer) Option {
	return func(c *config) {
		c.dialer = dialer
	}
}
