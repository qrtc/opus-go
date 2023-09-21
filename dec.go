package opus

/*
#cgo pkg-config: opus
#include <opus.h>


int opus_decoder_ctl_set_wrapped(OpusDecoder *st, int request, int value) {
	return opus_decoder_ctl(st, request, value);
}

int opus_decoder_ctl_get_wrapped(OpusDecoder *st, int request, int *value) {
	return opus_decoder_ctl(st, request, value);
}
*/
import "C"
import "unsafe"

// Opus Decoder Config
type OpusDecoderConfig struct {
	// Sampling rate of input signal (Hz).
	// This must be one of 8000, 12000, 16000, 24000, or 48000.
	SampleRate int
	// Number of channels in input signal.
	MaxChannels int
	// Decoder gain adjustment.
	Gain int
	// Enable inband forward error correction (FEC).
	EnableInbandFEC bool
}

// Opus Decoder
type OpusDecoder struct {
	// private handler
	ph *C.OpusDecoder
	// config
	OpusDecoderConfig
}

// Decode
func (dec *OpusDecoder) Decode(in, out []byte) (int, error) {
	var inPtr, outPtr unsafe.Pointer
	if len(in) > 0 {
		inPtr = unsafe.Pointer(&in[0])
	}
	if len(out) > 0 {
		outPtr = unsafe.Pointer(&out[0])
	}
	fecFlag := C.int(0)
	if dec.EnableInbandFEC {
		fecFlag = C.int(1)
	}
	bytesPerSample := 2
	frameSize := len(out) / bytesPerSample / dec.MaxChannels

	n := C.opus_decode(dec.ph, (*C.uchar)(inPtr), C.int(len(in)),
		(*C.opus_int16)(outPtr), C.int(frameSize), fecFlag)
	if n < 0 {
		return 0, genericErrors[int(n)]
	}

	return int(n) * bytesPerSample * dec.MaxChannels, nil
}

// Close
func (dec *OpusDecoder) Close() error {
	C.opus_decoder_destroy(dec.ph)
	dec.ph = nil
	return nil
}

// Create Opus Decoder
func CreateOpusDecoder(config *OpusDecoderConfig) (dec *OpusDecoder, err error) {
	config = populateDecConfig(config)
	dec = &OpusDecoder{
		OpusDecoderConfig: *config,
	}
	errNo := C.int(0)
	dec.ph = C.opus_decoder_create(
		C.int(dec.SampleRate), C.int(dec.MaxChannels), &errNo)
	if errNo != C.OPUS_OK {
		return nil, genericErrors[int(errNo)]
	}

	defer func() {
		if errNo != C.OPUS_OK {
			dec.Close()
		}
	}()

	if dec.Gain != 0 {
		if errNo = C.opus_decoder_ctl_set_wrapped(dec.ph, C.OPUS_SET_GAIN_REQUEST,
			C.int(dec.Gain)); errNo != C.OPUS_OK {
			return nil, genericErrors[int(errNo)]
		}
	}

	return dec, nil
}

func populateDecConfig(c *OpusDecoderConfig) *OpusDecoderConfig {
	if c == nil {
		c = &OpusDecoderConfig{}
	}
	if c.SampleRate == 0 {
		c.SampleRate = defaultSampleRate
	}
	if c.MaxChannels == 0 {
		c.MaxChannels = defaultMaxChannels
	}

	return c
}
