package main

import (
	"github.com/AviParampampam/go-ffmpeg"
)

func main() {
	f := ffmpeg.New(ffmpeg.NewConfig())

	createCommand(f, "/tmp/relaying/out.m3u8").Serve(nil, nil, 8)
}

func createCommand(f *ffmpeg.FFmpeg, output string) *ffmpeg.Command {
	inputGroup := ffmpeg.NewInputGroup(
		ffmpeg.NewInput(
			"rtsp://admin:12345678@192.168.1.101:554/ch01/0",
			ffmpeg.OptionProbesize(10000000),
			ffmpeg.OptionAnalyzeduration(10000000),
		),
	)

	outputGroup := ffmpeg.NewOutputGroup(
		ffmpeg.NewOutput(
			output,
			ffmpeg.OptionVideoCodec("copy"),
			ffmpeg.OptionNoAudio(),
			ffmpeg.OptionFormat("hls"),
			ffmpeg.OptionHlsTime(2),
			ffmpeg.OptionHlsListSize(6),
			"-hls_flags delete_segments",
		),
	)

	return f.NewCommand(inputGroup, outputGroup)
}
