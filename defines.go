package opus

/*
#cgo pkg-config: opus
#include <opus.h>
*/
import "C"
import "errors"

var genericErrors = map[int]error{
	C.OPUS_OK:               nil,
	C.OPUS_BAD_ARG:          errors.New("one or more invalid arguments"),
	C.OPUS_BUFFER_TOO_SMALL: errors.New("not enough bytes allocated in the buffer"),
	C.OPUS_INTERNAL_ERROR:   errors.New("an internal error was detected"),
	C.OPUS_INVALID_PACKET:   errors.New("the compressed data passed is corrupted"),
	C.OPUS_UNIMPLEMENTED:    errors.New("invalid request number"),
	C.OPUS_INVALID_STATE:    errors.New("en encoder or decoder structure is invalid or already freed"),
	C.OPUS_ALLOC_FAIL:       errors.New("memory allocation has failed"),
}

const (
	// Auto/default setting
	OpusAuto = int(C.OPUS_AUTO)
	// Maximum bitrate
	BitrateMax = int(C.OPUS_BITRATE_MAX)
)

// Application
type Application int

const (
	// Optimize encoding for VoIP
	AppVoIP = Application(C.OPUS_APPLICATION_VOIP)
	// Optimize encoding for non-voice signals like music
	AppAudio = Application(C.OPUS_APPLICATION_AUDIO)
	// Optimize encoding for low latency applications
	AppRestrictedLowdelay = Application(C.OPUS_APPLICATION_RESTRICTED_LOWDELAY)
)

// Signal type
type SignalType int

const (
	// Signal being encoded is voice
	SignalVoice = SignalType(C.OPUS_SIGNAL_VOICE)
	// Signal being encoded is music
	SignalMusic = SignalType(C.OPUS_SIGNAL_MUSIC)
)

// Bandwidth type
type BandwidthType int

const (
	// 4 kHz bandpass
	BandwidthNarrowband = BandwidthType(C.OPUS_BANDWIDTH_NARROWBAND)
	// 6 kHz bandpass
	BandwidthMediumband = BandwidthType(C.OPUS_BANDWIDTH_MEDIUMBAND)
	// 8 kHz bandpass
	BandwidthWideband = BandwidthType(C.OPUS_BANDWIDTH_WIDEBAND)
	// 12 kHz bandpass
	BandwidthSuperwideband = BandwidthType(C.OPUS_BANDWIDTH_SUPERWIDEBAND)
	// 20 kHz bandpass
	BandwidthFullband = BandwidthType(C.OPUS_BANDWIDTH_FULLBAND)
)

// Frame size type
type FrameSizeType int

const (
	// Select frame size from the argument (default)
	FramesizeArg = FrameSizeType(C.OPUS_FRAMESIZE_ARG)
	// Use 2.5 ms frames
	Framesize2Dot5Ms = FrameSizeType(C.OPUS_FRAMESIZE_2_5_MS)
	// Use 5 ms frames
	Framesize5Ms = FrameSizeType(C.OPUS_FRAMESIZE_5_MS)
	// Use 10 ms frames
	Framesize10Ms = FrameSizeType(C.OPUS_FRAMESIZE_10_MS)
	// Use 20 ms frames
	Framesize20Ms = FrameSizeType(C.OPUS_FRAMESIZE_20_MS)
	// Use 40 ms frames
	Framesize40Ms = FrameSizeType(C.OPUS_FRAMESIZE_40_MS)
	// Use 60 ms frames
	Framesize60Ms = FrameSizeType(C.OPUS_FRAMESIZE_60_MS)
	// Use 80 ms frames
	Framesize80Ms = FrameSizeType(C.OPUS_FRAMESIZE_80_MS)
	// Use 100 ms frames
	Framesize100Ms = FrameSizeType(C.OPUS_FRAMESIZE_100_MS)
	// Use 120 ms frames
	Framesize120Ms = FrameSizeType(C.OPUS_FRAMESIZE_120_MS)
)

const (
	defaultSampleRate  = 48000
	defaultMaxChannels = 2
	defaultApplication = AppAudio
)
