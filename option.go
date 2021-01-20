package ffmpeg

import (
	"strconv"
)

/* -+-+-+-+- INPUT -+-+-+-+- */

// OptionAnalyzeduration - specify how many microseconds are analyzed
// to probe the input.
func OptionAnalyzeduration(microseconds int) string {
	return "-analyzeduration " + strconv.Itoa(microseconds)
}

// OptionProbesize - set probing size in bytes, i.e. the size of the
// data to analyze to get stream information.
func OptionProbesize(bytes int) string {
	return "-probesize " + strconv.Itoa(bytes)
}

// OptionFormatFlags - set format flags.
func OptionFormatFlags(flags string) string {
	return "-fflags " + flags
}

// OptionLogLevel - set logging level and flags used by the library.
func OptionLogLevel(lvl string) string {
	return "-loglevel " + lvl
}

// OptionReport - dump full command line and log output to a file named
// program-YYYYMMDD-HHMMSS.log in the current directory.
func OptionReport() string {
	return "-report"
}

// OptionHideBanner - suppress printing banner.
func OptionHideBanner() string {
	return "-hide_banner"
}

// OptionCPUFlags - allows setting and clearing cpu flags.
func OptionCPUFlags(flags string) string {
	return "-cpuflags " + flags
}

// OptionFormat - force input or output file format.
func OptionFormat(frmt string) string {
	return "-f " + frmt
}

// OptionOverwrite - overwrite output files without asking.
func OptionOverwrite() string {
	return "-y"
}

// OptionNoOverwrite - do not overwrite output files, and exit immediately
// if a specified output file already exists.
func OptionNoOverwrite() string {
	return "-n"
}

// OptionStreamLoop - set number of times input stream shall be looped.
// Loop 0 means no loop, loop -1 means infinite loop.
func OptionStreamLoop(number int) string {
	return "-stream_loop " + strconv.Itoa(number)
}

/* -+-+-+-+- SEGMENT -+-+-+-+- */

// OptionSegmentFormat - override the inner container format, by default it
// is guessed by the filename extension.
func OptionSegmentFormat(frmt string) string {
	return "-segment_format " + frmt
}

// OptionSegmentFormatOptions - set output format options using a :-separated list of
// key=value parameters. Values containing the : special character must be escaped.
func OptionSegmentFormatOptions(opts string) string {
	return "-segment_format_options " + opts
}

// OptionSegmentList - generate also a listfile named name. If not specified
// no listfile is generated.
func OptionSegmentList(name string) string {
	return "-segment_list " + name
}

// OptionSegmentListFlags - set flags affecting the segment list generation.
func OptionSegmentListFlags(flags string) string {
	return "-segment_list_flags " + flags
}

// OptionSegmentListSize - update the list file so that it contains at most size segments.
// If 0 the list file will contain all the segments. Default value is 0.
func OptionSegmentListSize(size int) string {
	return "-segment_list_size " + strconv.Itoa(size)
}

// OptionSegmentListEntryPrefix - prepend prefix to each entry. Useful to generate
// absolute paths. By default no prefix is applied.
func OptionSegmentListEntryPrefix(prefix string) string {
	return "-segment_list_entry_prefix " + prefix
}

// OptionSegmentListType - select the listing format.
func OptionSegmentListType(tp string) string {
	return "-segment_list_type " + tp
}

// OptionSegmentTime - set segment duration to time, the value must be a duration specification.
func OptionSegmentTime(time int) string {
	return "-segment_time " + strconv.Itoa(time)
}

// OptionSegmentAtClockTime - If set to "true" split at regular clock time intervals starting
// from 00:00 o’clock. The time value specified in segment_time is used for setting
// the length of the splitting interval.
func OptionSegmentAtClockTime(v bool) string {
	if v {
		return "-segment_atclocktime 1"
	}
	return "-segment_atclocktime 0"
}

// OptionSegmentClockTimeOffset - delay the segment splitting times with the specified
// duration when using segment_atclocktime.
func OptionSegmentClockTimeOffset(dur int) string {
	return "-segment_clocktime_offset " + strconv.Itoa(dur)
}

// OptionSegmentClockTimeWrapDuration - force the segmenter to only start a new segment
// if a packet reaches the muxer within the specified duration after the segmenting clock time.
func OptionSegmentClockTimeWrapDuration(dur int) string {
	return "-segment_clocktime_wrap_duration " + strconv.Itoa(dur)
}

// OptionSegmentDelta - specify the accuracy time when selecting the start time for a segment,
// expressed as a duration specification. Default value is "0".
func OptionSegmentDelta(delta int) string {
	return "-segment_time_delta " + strconv.Itoa(delta)
}

// OptionSegmentTimes - specify a list of split points. times contains a list of comma
// separated duration specifications, in increasing order. See also the segment_time option.
func OptionSegmentTimes(times string) string {
	return "-segment_times " + times
}

// OptionSegmentFrames - specify a list of split video frame numbers. frames contains a
// list of comma separated integer numbers, in increasing order.
func OptionSegmentFrames(frames string) string {
	return "-segment_frames " + frames
}

// OptionSegmentWrap - wrap around segment index once it reaches limit.
func OptionSegmentWrap(limit int) string {
	return "-segment_wrap " + strconv.Itoa(limit)
}

// OptionSegmentStartNumber - set the sequence number of the first segment.
func OptionSegmentStartNumber(number int) string {
	return "-segment_start_number " + strconv.Itoa(number)
}

/* -+-+-+-+- RTSP -+-+-+-+- */

// OptionTimeout - set maximum timeout (in seconds) to wait for incoming connections.
func OptionTimeout(timeout int) string {
	return "-timeout " + strconv.Itoa(timeout)
}

// OptionRTSPTransport - set RTSP transport protocols.
func OptionRTSPTransport(trans string) string {
	return "-rtsp_transport " + trans
}

// OptionRTSPFlags - set RTSP flags.
func OptionRTSPFlags(flags string) string {
	return "-rtsp_flags " + flags
}

// OptionAllowedMediaTypes - set media types to accept from the server.
func OptionAllowedMediaTypes(types string) string {
	return "-allowed_media_types " + types
}

// OptionMinPort - set minimum local UDP port. Default value is 5000.
func OptionMinPort(port int) string {
	return "-min_port " + strconv.Itoa(port)
}

// OptionMaxPort - set maximum local UDP port. Default value is 65000.
func OptionMaxPort(port int) string {
	return "-max_port " + strconv.Itoa(port)
}

// OptionReorderQueueSize - set number of packets to buffer for handling of reordered packets.
func OptionReorderQueueSize(number int) string {
	return "-reorder_queue_size " + strconv.Itoa(number)
}

// OptionSTimeout - set socket TCP I/O timeout in microseconds.
func OptionSTimeout(timeout int) string {
	return "-stimeout " + strconv.Itoa(timeout)
}

// OptionUserAgent - override User-Agent header. If not specified, it
// defaults to the libavformat identifier string.
func OptionUserAgent(ua string) string {
	return "user-agent " + ua
}

/* -+-+-+-+- OUTPUT -+-+-+-+- */

// OptionCodec - select an encoder or a decoder for all streams.
func OptionCodec(codec string) string {
	return "-c " + codec
}

// OptionFPS - set frame rate (Hz value, fraction or abbreviation).
func OptionFPS(fps int) string {
	return "-fflags " + strconv.Itoa(fps)
}

// OptionVideoCodec - set the video codec.
func OptionVideoCodec(codec string) string {
	return "-vcodec " + codec
}

// OptionAudioCodec - set the audio codec.
func OptionAudioCodec(codec string) string {
	return "-acodec " + codec
}

// OptionNoAudio - as an input option, blocks all audio streams of a file from being filtered
// or being automatically selected or mapped for any output.
func OptionNoAudio() string {
	return "-an"
}

// OptionStrftime - use the strftime function to define the name of the new segments to write.
func OptionStrftime(v bool) string {
	if v {
		return "-strftime 1"
	}
	return "-strftime 0"
}

// OptionResetTimestamp - reset timestamps at the beginning of each segment, so that each
// segment will start with near-zero timestamps.
func OptionResetTimestamp(v bool) string {
	if v {
		return "-reset_timestamps 1"
	}
	return "-reset_timestamps 0"
}
