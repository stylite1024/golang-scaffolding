package config

type AppConfig struct {
	*ApplicationConfig `mapstructure:"application"`
	*LogConfig         `mapstructure:"log"`
}

type ApplicationConfig struct {
	Mode    string `mapstructure:"mode"`
}

type LogConfig struct {
	Mode          string `mapstructure:"mode"`
	InfoFilename  string `mapstructure:"info_filename"`
	ErrorFilename string `mapstructure:"error_filename"`
	MaxSize       int    `mapstructure:"max_size"`
	MaxAge        int    `mapstructure:"max_age"`
	MaxBackups    int    `mapstructure:"max_backups"`
	Compress      bool   `mapstructure:"compress"`
}
