package ffmpeg

// Config is a configuration for ffmpeg.
type Config struct {
	LogLevel string `toml:"log_level"`
}

// NewConfig is a function for creating new config.
func NewConfig() *Config {
	return &Config{
		LogLevel: "debug",
	}
}
