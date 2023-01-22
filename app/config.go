package app

import "time"

type Config struct {
	Service  serviceConfig
	Store    storeConfig
	Server   serverConfig
	Security securityConfig
	Auth     authConfig
}

type serviceConfig struct {
	Name string
	Port int
}

type storeConfig struct {
	Client  string
	Host    string
	Port    int
	Name    string
	Timeout time.Duration
}

type serverConfig struct {
	Timeout struct {
		Read     time.Duration
		Write    time.Duration
		Idle     time.Duration
		Shutdown time.Duration
	}
}

type securityConfig struct {
	Allowed struct {
		Origins []string
		Methods []string
		Headers []string
	}
	AllowCredentials bool
}

type authConfig struct {
	Csrf string
}
