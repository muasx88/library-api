package config

var Config config

type config struct {
	App appConfig `mapstructure:"app"`
	DB  dBConfig  `mapstructure:"db"`
}

type appConfig struct {
	Env        string           `mapstructure:"env"`
	Version    string           `mapstructure:"version"`
	Name       string           `mapstructure:"name"`
	Port       string           `mapstructure:"port"`
	Encryption encryptionConfig `mapstructure:"encryption"`
}

type encryptionConfig struct {
	Salt      uint8  `mapstructure:"salt"`
	JWTSecret string `mapstructure:"jwt_secret"`
}

type dBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`

	ConnectionPool dBConnectionPoolConfig `mapstructure:"connection_pool"`
}

type dBConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `mapstructure:"max_idle_connection"`
	MaxOpenConnetcion     uint8 `mapstructure:"max_open_connection"`
	MaxLifetimeConnection uint8 `mapstructure:"max_lifetime_connection"`
	MaxIdletimeConnection uint8 `mapstructure:"max_idletime_connection"`
}
