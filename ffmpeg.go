package ffmpeg

type OptionIO interface {
	String() string
}

type report struct {
	File     string
	LogLevel int
}

type FFmpeg struct {
	BinPath string
	Report  report
	workers map[string]*Worker
}

func New() *FFmpeg {
	return &FFmpeg{
		BinPath: "ffmpeg",
		Report:  report{},
		workers: make(map[string]*Worker),
	}
}
