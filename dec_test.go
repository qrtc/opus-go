package opus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("decoder test", func() {
	var (
		decoder *OpusDecoder
		err     error
	)

	BeforeEach(func() {
		decoder = nil
		err = nil
	})

	It("decoder create and close", func() {
		decoder, err = CreateOpusDecoder(&OpusDecoderConfig{
			SampleRate:  48000,
			MaxChannels: 2,
		})
		Expect(err).To(BeNil())
		Expect(decoder).NotTo(BeNil())

		decoder.Close()
		Expect(decoder.ph).To(BeNil())
	})

	It("decoder decode", func() {
		decoder, err = CreateOpusDecoder(&OpusDecoderConfig{
			SampleRate:  48000,
			MaxChannels: 1,
		})

		in := []byte{0xEC, 0xFF, 0xFE}
		out := make([]byte, 4096)
		n, err := decoder.Decode(in, out)
		Expect(err).To(BeNil())
		Expect(n).NotTo(Equal(0))

		decoder.Close()
	})
})
