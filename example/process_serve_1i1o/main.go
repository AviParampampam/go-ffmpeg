package main

import (
	"github.com/AviParampampam/go-ffmpeg"
)

func main() {
	f := ffmpeg.New(ffmpeg.NewConfig())

	go createCommand(f, "rtsp://admin:12345678@192.168.1.15:554/ch01.264?dev=1", "records/cam1/%Y-%m-%dT%H:%M:%S.mp4").Serve(nil, nil, 8)
	go createCommand(f, "rtsp://admin:12345678@192.168.1.70:554/h264Preview_01_main", "records/cam2/%Y-%m-%dT%H:%M:%S.mp4").Serve(nil, nil, 8)
	createCommand(f, "rtsp://admin:12345678@192.168.1.101:554/ch01/0", "records/cam3/%Y-%m-%dT%H:%M:%S.mp4").Serve(nil, nil, 8)

	// createCommand(f, "rtsp://admin:12345678@192.168.1.101:554/ch01/0", "records/cam3/%Y-%m-%dT%H:%M:%S.mp4").Serve(os.Stderr, os.Stdout)
}

func createCommand(f *ffmpeg.FFmpeg, input, output string) *ffmpeg.Command {
	inputGroup := ffmpeg.NewInputGroup(
		ffmpeg.NewInput(
			input,
			ffmpeg.OptionProbesize(10000000),
			ffmpeg.OptionAnalyzeduration(10000000),
		),
	)

	outputGroup := ffmpeg.NewOutputGroup(
		ffmpeg.NewOutput(
			output,
			ffmpeg.OptionVideoCodec("copy"),
			ffmpeg.OptionFormat("segment"),
			ffmpeg.OptionSegmentFormat("mp4"),
			ffmpeg.OptionSegmentTime(16),
			ffmpeg.OptionStrftime(true),
			ffmpeg.OptionResetTimestamp(true),
		),
	)

	return f.NewCommand(inputGroup, outputGroup)
}
