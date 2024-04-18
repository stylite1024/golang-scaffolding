package config

type AppConfig struct {
	*ApplicationConfig `mapstructure:"application"`
	*LogConfig         `mapstructure:"log"`
}

type ApplicationConfig struct {
	Mode string `mapstructure:"mode"`
}

type LogConfig struct {
	LogLevel    string `mapstructure:"log_level"`
	LogFilename string `mapstructure:"log_filename"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxBackups  int    `mapstructure:"max_backups"`
	Compress    bool   `mapstructure:"compress"`
}
