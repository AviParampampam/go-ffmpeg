package ffmpeg

// Config is a configuration for ffmpeg.
type Config struct {
	FFmpegPath string
}

// NewConfig is a function for creating new config.
func NewConfig() *Config {
	return &Config{
		FFmpegPath: "/usr/bin/ffmpeg",
	}
}
