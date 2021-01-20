package main

import (
	"os"

	"github.com/AviParampampam/Metis/pkg/ffmpeg"
)

func main() {
	config := ffmpeg.NewConfig()
	f := ffmpeg.New(config)

	createCommand(f).Start(os.Stderr, os.Stdout)
}

func createCommand(f *ffmpeg.FFmpeg) *ffmpeg.Command {
	input := ffmpeg.NewInput(
		"rtsp://admin:12345678@192.168.1.15:554/ch01.264?dev=1",
		ffmpeg.OptionHideBanner(),
		ffmpeg.OptionFormatFlags("nobuffer"),
		ffmpeg.OptionProbesize(10000000),
		ffmpeg.OptionAnalyzeduration(10000000),
	)

	output1 := ffmpeg.NewOutput(
		"records/cam1/high/%Y-%m-%dT%H:%M:%S.mp4",
		ffmpeg.OptionFPS(25),
		ffmpeg.OptionVideoCodec("copy"),
		ffmpeg.OptionFormat("segment"),
		ffmpeg.OptionSegmentFormat("mp4"),
		ffmpeg.OptionSegmentTime(16),
		ffmpeg.OptionStrftime(true),
	)
	output2 := ffmpeg.NewOutput(
		"records/cam1/low/%Y-%m-%dT%H:%M:%S.mp4",
		ffmpeg.OptionFPS(15),
		ffmpeg.OptionVideoCodec("copy"),
		ffmpeg.OptionFormat("segment"),
		ffmpeg.OptionSegmentFormat("mp4"),
		ffmpeg.OptionSegmentTime(24),
		ffmpeg.OptionStrftime(true),
	)

	ig := ffmpeg.NewInputGroup(input)
	og := ffmpeg.NewOutputGroup(output1, output2)

	return f.NewCommand(ig, og)
}