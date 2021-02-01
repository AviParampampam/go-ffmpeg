package main

import (
	"os"

	ffmpeg "github.com/AviParampampam/go-ffmpeg"
)

func main() {
	f := ffmpeg.New(ffmpeg.NewConfig())

	createCommand(f).Run(os.Stderr, os.Stdout)
}

func createCommand(f *ffmpeg.FFmpeg) *ffmpeg.Command {
	input := ffmpeg.NewInput(
		"rtsp://admin:12345678@84.39.252.195:1116/0",
		ffmpeg.OptionRTSPTransport("tcp"),
		ffmpeg.OptionHideBanner(),
		ffmpeg.OptionFormatFlags("nobuffer"),
		ffmpeg.OptionProbesize(20000000),
		ffmpeg.OptionAnalyzeduration(20000000),
	)

	// output1 := ffmpeg.NewOutput(
	// 	"records/cam1/high/%Y-%m-%dT%H:%M:%S.mp4",
	// 	ffmpeg.OptionFPS(25),
	// 	ffmpeg.OptionVideoCodec("copy"),
	// 	ffmpeg.OptionFormat("segment"),
	// 	ffmpeg.OptionSegmentFormat("mp4"),
	// 	ffmpeg.OptionSegmentTime(16),
	// 	ffmpeg.OptionStrftime(true),
	// )
	output2 := ffmpeg.NewOutput(
		"/mnt/4tb2/records/%Y-%m-%dT%H:%M:%S.mp4",
		ffmpeg.OptionSTimeout(60000000),
		ffmpeg.OptionFPS(20),
		ffmpeg.OptionVideoCodec("copy"),
		ffmpeg.OptionFormat("segment"),
		ffmpeg.OptionSegmentFormat("mp4"),
		ffmpeg.OptionSegmentTime(60),
		ffmpeg.OptionStrftime(true),
		ffmpeg.OptionResetTimestamp(true),
	)

	// Grouping is necessary
	ig := ffmpeg.NewInputGroup(input)
	og := ffmpeg.NewOutputGroup(output2)

	return f.NewCommand(ig, og)
}
