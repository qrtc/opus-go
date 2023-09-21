package opus

/*
#cgo pkg-config: opus
#include <opus.h>

int opus_encoder_ctl_set_wrapped(OpusEncoder *st, int request, int value) {
	return opus_encoder_ctl(st, request, value);
}

int opus_encoder_ctl_get_wrapped(OpusEncoder *st, int request, int *value) {
	return opus_encoder_ctl(st, request, value);
}
*/
import "C"
import "unsafe"

// Opus Encoder Config
type OpusEncoderConfig struct {
	// Sampling rate of input signal (Hz).
	// This must be one of 8000, 12000, 16000, 24000, or 48000.
	SampleRate int
	// Number of channels in input signal.
	MaxChannels int
	// Encode mode.
	Application Application
	// Enable discontinuous transmission (DTX).
	EnableDTX bool
	// Enable inband forward error correction (FEC).
	EnableInbandFEC bool
	// Disable almost all use of prediction.
	DisablePrediction bool
	// Disable variable bitrate (VBR).
	DisableVBR bool
	// Disable constrained VBR.
	DisableConstrainedVBR bool
	// Disable the use of phase inversion for intensity stereo.
	DisablePhaseInversion bool
	// Rates from 500 to 512000 bits per second are meaningful,
	// as well as the special values BitrateAuto and BitrateMax.
	Bitrate int
	// Complexity configuration, a value in the range 0-10.
	Complexity int
	// The maximum bandpass that the encoder will select automatically.
	MaxBandwidth BandwidthType
	// Encoder's bandpass to a specific value.
	Bandwidth BandwidthType
	// Expected packet loss percentage.
	PacketLossPercent int
	// The encoder's use of variable duration frames.
	FrameDuration FrameSizeType
	// The type of signal being encoded.
	SignalType SignalType
	// Force mono or stereo.
	ForceChannels int
	// The depth of signal being encoded.
	LSBDepth int
}

// Opus Encoder
type OpusEncoder struct {
	// private handler
	ph *C.OpusEncoder
	// config
	OpusEncoderConfig
}

// Encode
func (enc *OpusEncoder) Encode(in, out []byte) (int, error) {
	var inPtr, outPtr unsafe.Pointer
	if len(in) > 0 {
		inPtr = unsafe.Pointer(&in[0])
	}
	if len(out) > 0 {
		outPtr = unsafe.Pointer(&out[0])
	}
	bytesPerSample := 2
	frameSize := len(in) / bytesPerSample / enc.MaxChannels
	n := C.opus_encode(enc.ph, (*C.opus_int16)(inPtr), C.int(frameSize),
		(*C.uchar)(outPtr), C.opus_int32(len(out)))
	if n < 0 {
		return 0, genericErrors[int(n)]
	}

	return int(n), nil
}

// Close
func (enc *OpusEncoder) Close() error {
	C.opus_encoder_destroy(enc.ph)
	enc.ph = nil
	return nil
}

// Lookahead
func (enc *OpusEncoder) Lookahead() int {
	lookAhead := C.int(0)
	C.opus_encoder_ctl_get_wrapped(enc.ph, C.OPUS_GET_LOOKAHEAD_REQUEST, &lookAhead)
	return int(lookAhead)
}

// InDTX
func (enc *OpusEncoder) InDTX() bool {
	inDTX := C.int(0)
	errNo := C.opus_encoder_ctl_get_wrapped(enc.ph, C.OPUS_GET_IN_DTX_REQUEST, &inDTX)
	if errNo == C.OPUS_OK && inDTX != 0 {
		return true
	}
	return false
}

// Create Opus Encoder
func CreateOpusEncoder(config *OpusEncoderConfig) (enc *OpusEncoder, err error) {
	config = populateEncConfig(config)
	enc = &OpusEncoder{
		OpusEncoderConfig: *config,
	}
	errNo := C.int(0)
	enc.ph = C.opus_encoder_create(
		C.int(enc.SampleRate), C.int(enc.MaxChannels), C.int(enc.Application), &errNo)
	if errNo != C.OPUS_OK {
		return nil, genericErrors[int(errNo)]
	}

	defer func() {
		if errNo != C.OPUS_OK {
			enc.Close()
		}
	}()

	if enc.EnableDTX {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_DTX_REQUEST,
			C.int(1)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.EnableInbandFEC {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_INBAND_FEC_REQUEST,
			C.int(1)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.DisablePrediction {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_PREDICTION_DISABLED_REQUEST,
			C.int(1)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.DisableVBR {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_VBR_REQUEST,
			C.int(0)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.DisableConstrainedVBR {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_VBR_CONSTRAINT_REQUEST,
			C.int(0)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.DisablePhaseInversion {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST,
			C.int(1)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.Bitrate != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_BITRATE_REQUEST,
			C.int(enc.Bitrate)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.Complexity != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_COMPLEXITY_REQUEST,
			C.int(enc.Complexity)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.MaxBandwidth != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_MAX_BANDWIDTH_REQUEST,
			C.int(enc.MaxBandwidth)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.Bandwidth != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_BANDWIDTH_REQUEST,
			C.int(enc.Bandwidth)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.PacketLossPercent != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_PACKET_LOSS_PERC_REQUEST,
			C.int(enc.PacketLossPercent)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.FrameDuration != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_EXPERT_FRAME_DURATION_REQUEST,
			C.int(enc.FrameDuration)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.SignalType != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_SIGNAL_REQUEST,
			C.int(enc.SignalType)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}

	if enc.ForceChannels != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_FORCE_CHANNELS_REQUEST,
			C.int(enc.ForceChannels)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}
	if enc.LSBDepth != 0 {
		if errNo = C.opus_encoder_ctl_set_wrapped(enc.ph, C.OPUS_SET_LSB_DEPTH_REQUEST,
			C.int(enc.LSBDepth)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}

	return enc, nil
}

func populateEncConfig(c *OpusEncoderConfig) *OpusEncoderConfig {
	if c == nil {
		c = &OpusEncoderConfig{}
	}
	if c.SampleRate == 0 {
		c.SampleRate = defaultSampleRate
	}
	if c.MaxChannels == 0 {
		c.MaxChannels = defaultMaxChannels
	}
	if c.Application == 0 {
		c.Application = defaultApplication
	}

	return c
}
