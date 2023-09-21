package opus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("encoder test", func() {
	var (
		encoder *OpusEncoder
		err     error
	)

	BeforeEach(func() {
		encoder = nil
		err = nil
	})

	It("encoder create and close", func() {
		encoder, err = CreateOpusEncoder(&OpusEncoderConfig{
			SampleRate:  48000,
			MaxChannels: 2,
			Application: AppAudio,
		})
		Expect(err).To(BeNil())
		Expect(encoder).NotTo(BeNil())

		encoder.Close()
		Expect(encoder.ph).To(BeNil())
	})

	It("encoder encode", func() {
		encoder, err = CreateOpusEncoder(&OpusEncoderConfig{
			SampleRate:    48000,
			MaxChannels:   2,
			Application:   AppAudio,
			FrameDuration: Framesize5Ms,
		})

		frameSize := encoder.SampleRate * encoder.MaxChannels * 5 * (16 / 2) / 1000

		in := make([]byte, frameSize)
		out := make([]byte, frameSize)
		n, err := encoder.Encode(in, out)
		Expect(err).To(BeNil())
		Expect(n).NotTo(Equal(0))

		encoder.Close()
	})
})
