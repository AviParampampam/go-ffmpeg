package ffmpeg

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

// FFmpeg is a structure to controlling ffmpeg processes.
type FFmpeg struct {
	// Path to binary files.
	FFmpegBin  string
	FFprobeBin string

	Version   string
	UserAgent string
	config    *Config

	// Logger.
	LogFile string `toml:"log_file"`
}

// New is an function to creating new FFmpeg structure.
func New(config *Config) *FFmpeg {
	return &FFmpeg{
		UserAgent: "Signum-Video Software 0.1",
		config:    config,
	}
}

func (f *FFmpeg) Init() (*FFmpeg, error) {
	// Gettings ffmpeg version.
	out, err := exec.Command("ffmpeg", "-version").Output()
	if err != nil {
		return nil, err
	}
	if reResult := regexp.MustCompile(`ffmpeg version \S*`).FindString(string(out)); len(reResult) >= 2 {
		f.Version = strings.Split(reResult, " ")[2]
	} else {
		return nil, errors.New("failed to get the FFmpeg version")
	}

	// Gettings ffmpeg bin.
	out, err = exec.Command("which", "ffmpeg").Output()
	if err != nil {
		return nil, err
	}
	if result := string(out); len(result) > 0 {
		f.FFmpegBin = result
	} else {
		return nil, errors.New("failed to get the FFmpeg binary")
	}

	// Gettings ffprobe bin.
	out, err = exec.Command("which", "ffprobe").Output()
	if err != nil {
		return nil, err
	}
	if result := string(out); len(result) > 0 {
		f.FFprobeBin = result
	} else {
		return nil, errors.New("failed to get the FFprobe binary")
	}

	return f, nil
}
