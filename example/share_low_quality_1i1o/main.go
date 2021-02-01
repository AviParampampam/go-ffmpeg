package main

import (
	"os"

	"github.com/AviParampampam/go-ffmpeg"
)

func main() {
	f := ffmpeg.New(ffmpeg.NewConfig())

	// go createCommand(f, "rtsp://admin:12345678@192.168.1.101:554/ch01/0", "/tmp/relaying_low/out1.mpd").Serve(os.Stderr, os.Stdout, 8)
	// go createCommand(f, "rtsp://admin:12345678@192.168.1.15:554/ch01.264?dev=1", "/tmp/relaying_low/out2.mpd").Serve(os.Stderr, os.Stdout, 8)
	cmd := createCommand(f, "rtsp://admin:12345678@192.168.1.101:554/ch01/0", "./relaying/high/out.m3u8")
	go cmd.Serve(os.Stderr, os.Stdout, 8)

	cmd = createCommand(f, "rtsp://admin:12345678@192.168.1.101:554/ch01/1", "./relaying/low/out.m3u8")
	cmd.Serve(os.Stderr, os.Stdout, 8)
}

func createCommand(f *ffmpeg.FFmpeg, input, output string) *ffmpeg.Command {
	inputGroup := ffmpeg.NewInputGroup(
		ffmpeg.NewInput(
			input,
			ffmpeg.OptionHideBanner(),
			ffmpeg.OptionProbesize(10000000),
			ffmpeg.OptionAnalyzeduration(10000000),
			ffmpeg.OptionRTSPTransport("tcp"),
		),
	)

	outputGroup := ffmpeg.NewOutputGroup(
		ffmpeg.NewOutput(
			output,
			// "-b:v 512K",
			// "-bufsize 1M",
			// "-maxrate 768K",

			ffmpeg.OptionPreset("ultrafast"),
			// "-preset veryslow",
			// "-tune zerolatency",
			// "-strict -2",
			// "-threads 1",

			// "-tune psnr",
			"-y",
			// "-qscale 0",

			// "-g 30",
			ffmpeg.OptionVideoCodec("copy"),
			// "-q:v 5",
			ffmpeg.OptionNoAudio(),
			// "-crf 42",
			// "-ss 15",
			// ffmpeg.OptionSize("1920x1080"),
			// ffmpeg.OptionSize("640x360"),
			// ffmpeg.OptionSize("360x240"),
			// ffmpeg.OptionFPS(2),
			// ffmpeg.OptionSize("100x50"),
			// "-profile:v baseline",
			// "-level:v 4.0",

			ffmpeg.OptionFormat("hls"),
			ffmpeg.OptionHlsTime(1),
			ffmpeg.OptionHlsListSize(4),
			ffmpeg.OptionHlsWrap(4),
			ffmpeg.OptionHlsFlags("delete_segments"),
			// "-hls_allow_cache 0",

			// ffmpeg.OptionFormat("dash"),
		),
	)

	return f.NewCommand(inputGroup, outputGroup)
}
