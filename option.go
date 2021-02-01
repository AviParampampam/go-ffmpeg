package ffmpeg

import (
	"strconv"
)

/*
	============== HLS =============
*/

// OptionHlsInitTime - set the initial target segment length in seconds.
// Default value is 0. Segment will be cut on the next key frame after
// this time has passed on the first m3u8 list. After the initial playlist
// is filled ffmpeg will cut segments at duration equal to hls_time.
func OptionHlsInitTime(sec int) string {
	return "-hls_init_time " + strconv.Itoa(sec)
}

// OptionHlsTime - set the target segment length in seconds.
// Default value is 2. Segment will be cut on the next key frame
// after this time has passed.
func OptionHlsTime(sec int) string {
	return "-hls_time " + strconv.Itoa(sec)
}

// OptionHlsListSize - set the maximum number of playlist entries.
// If set to 0 the list file will contain all the segments. Default
// value is 5.
func OptionHlsListSize(size int) string {
	return "-hls_list_size " + strconv.Itoa(size)
}

// OptionHlsDeleteThreshold - set the number of unreferenced segments
// to keep on disk before hls_flags delete_segments deletes them.
// Increase this to allow continue clients to download segments which
// were recently referenced in the playlist. Default value is 1,
// meaning segments older than hls_list_size+1 will be deleted.
func OptionHlsDeleteThreshold(size int) string {
	return "-hls_delete_threshold " + strconv.Itoa(size)
}

func OptionHlsFlags(flags string) string {
	return "-hls_flags " + flags
}

// OptionHlsTsOptions - set output format options using a :-separated
// list of key=value parameters. Values containing : special characters
// must be escaped.
func OptionHlsTsOptions(optionsList string) string {
	return "-hls_ts_options " + optionsList
}

// OptionHlsWrap - this is a deprecated option, you can use hls_list_size
// and hls_flags delete_segments instead it.
// This option is useful to avoid to fill the disk with many segment files,
// and limits the maximum number of segment files written to disk to wrap.
func OptionHlsWrap(wrap int) string {
	return "-hls_wrap " + strconv.Itoa(wrap)
}

// OptionHlsStartNumberSource - start the playlist sequence number
// (#EXT-X-MEDIA-SEQUENCE) according to the specified source. Unless
// hls_flags single_file is set, it also specifies source of starting
// sequence numbers of segment and subtitle filenames. In any case,
// if hls_flags append_list is set and read playlist sequence number
// is greater than the specified start sequence number, then that
// value will be used as start value.
//
// It accepts the following values:
//	generic (default)
//		Set the starting sequence numbers according to start_number option value.
//  epoch
//		The start number will be the seconds since epoch (1970-01-01 00:00:00).
//  epoch_us
//		The start number will be the microseconds since epoch (1970-01-01 00:00:00).
//  datetime
//		The start number will be based on the current date/time as YYYYmmddHHMMSS.
//		e.g. 20161231235759.
func OptionHlsStartNumberSource(value string) string {
	return "-hls_start_number_source " + value
}

// OptionStartNumber - start the playlist sequence number (#EXT-X-MEDIA-SEQUENCE)
// from the specified number when hls_start_number_source value is generic.
// (This is the default case.) Unless hls_flags single_file is set, it also
// specifies starting sequence numbers of segment and subtitle filenames. Default value is 0.
func OptionStartNumber(number int) string {
	return "-start_number " + strconv.Itoa(number)
}

// OptionHlsAllowCache - explicitly set whether the client MAY (1) or MUST NOT (0) cache
// media segments.
func OptionHlsAllowCache(value bool) string {
	if value {
		return "-hls_allow_cache 1"
	}
	return "-hls_allow_cache 0"
}

// OptionHlsBaseUrl - append baseurl to every entry in the playlist. Useful to generate
// playlists with absolute paths.
// Note that the playlist sequence number must be unique for each segment and it is not
// to be confused with the segment filename sequence number which can be cyclic, for
// example if the wrap option is specified.
func OptionHlsBaseUrl(baseurl string) string {
	return "-hls_base_url " + baseurl
}

// OptionHlsSegmentFilename - set the segment filename. Unless hls_flags single_file is
// set, filename is used as a string format with the
// segment number: -hls_segment_filename 'file%03d.ts'
// This example will produce the playlist, out.m3u8, and segment files: file000.ts,
// file001.ts, file002.ts, etc.
func OptionHlsSegmentFilename(filename string) string {
	return "-hls_segment_filename " + filename
}

// OptionUseLocaltime - same as strftime option, will be deprecated.
func OptionUseLocaltime() string {
	return "-use_localtime"
}

// OptionUseLocaltimeMkdir - same as strftime_mkdir option, will be deprecated.
func OptionUseLocaltimeMkdir(value bool) string {
	if value {
		return "-use_localtime_mkdir 1"
	}
	return "-use_localtime_mkdir 0"
}

/*
	============== RTSP =============
*/

// OptionRTSPTransport - set RTSP transport protocols.
// It accepts the following values:
// 'udp' - use UDP as lower transport protocol.
// 'tcp' - use TCP (interleaving within the RTSP control channel) as lower transport protocol.
// 'udp_multicast' - use UDP multicast as lower transport protocol.
// 'http' - use HTTP tunneling as lower transport protocol, which is useful for passing proxies.
//
// Multiple lower transport protocols may be specified, in that case they are tried one at a
// time (if the setup of one fails, the next one is tried). For the muxer, only the ‘tcp’
// and ‘udp’ options are supported.
func OptionRTSPTransport(trans string) string {
	return "-rtsp_transport " + trans
}

// OptionRTSPFlags - set RTSP flags.
// The following values are accepted:
// 'filter_src' - accept packets only from negotiated peer address and port.
// 'listen' - act as a server, listening for an incoming connection.
// 'prefer_tcp' - try TCP for RTP transport first, if TCP is available as RTSP RTP transport.
//
// Default value is 'none'
func OptionRTSPFlags(flags string) string {
	return "-rtsp_flags " + flags
}

// OptionAllowedMediaTypes - set media types to accept from the server.
// The following flags are accepted:
// 'video'
// 'audio'
// 'data'
//
// By default is accepts all media types.
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

/*
	============== SEGMENT =============
*/

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

// OptionResetTimestamp - reset timestamps at the beginning of each segment, so that each
// segment will start with near-zero timestamps.
func OptionResetTimestamp(v bool) string {
	if v {
		return "-reset_timestamps 1"
	}
	return "-reset_timestamps 0"
}

/*
	============== CODEC =============
*/

// OptionCodec - select an encoder or a decoder for all streams.
func OptionCodec(codec string) string {
	return "-c " + codec
}

// OptionVideoCodec - set the video codec.
func OptionVideoCodec(codec string) string {
	return "-vcodec " + codec
}

// OptionAudioCodec - set the audio codec.
func OptionAudioCodec(codec string) string {
	return "-acodec " + codec
}

/*
	============== GLOBAL INPUT =============
*/

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

// OptionTimeout - set maximum timeout (in seconds) to wait for incoming connections.
func OptionTimeout(timeout int) string {
	return "-timeout " + strconv.Itoa(timeout)
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

// OptionFPS - set frame rate (Hz value, fraction or abbreviation).
func OptionFPS(fps int) string {
	return "-fflags " + strconv.Itoa(fps)
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

// OptionFormat - force input or output file format.
func OptionFormat(frmt string) string {
	return "-f " + frmt
}

func OptionSize(size string) string {
	return "-s " + size
}

// OptionPreset - collection of options that will provide a certain encoding speed to compression ratio.
func OptionPreset(preset string) string {
	return "-preset " + preset
}
