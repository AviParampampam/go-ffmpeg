package ffmpeg

// FFmpeg is a structure for controlling ffmpeg processes.
type FFmpeg struct {
	config *Config
}

// New is an function for creating new FFmpeg structure.
func New(config *Config) *FFmpeg {
	return &FFmpeg{
		config: config,
	}
}
