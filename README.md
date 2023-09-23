[![PkgGoDev](https://pkg.go.dev/badge/github.com/qrtc/opus-go)](https://pkg.go.dev/github.com/qrtc/opus-go)

# opus-go

Go bindings for [opus](https://github.com/xiph/opus). Modern audio compression for the internet.

## Why opus-go

The purpose of opus-go is easing the adoption of opus codec library. Using Go, with just a few lines of code you can implement an application that encode/decode data easy.

##  Is this a new implementation of opus?

No! We are just exposing the great work done by the research organization of [Xiph](https://xiph.org/) as a golang library. All the functionality and implementation still resides in the official opus project.

# Features supported

- Decode Opus to PCM
- Encode PCM to Opus

# Usage

## Decode AAC frame to PCM

```go
package main

import (
	"fmt"

	opus "github.com/qrtc/opus-go"
)

func main() {
	decoder, err := opus.CreateOpusDecoder(&opus.OpusDecoderConfig{
		SampleRate:  48000,
		MaxChannels: 2,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		decoder.Close()
	}()

	inBuf := []byte{
		// Opus frame
	}
	outBuf := make([]byte, 4096)

	n, err := decoder.Decode(inBuf, outBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outBuf[0:n])
}
```

## Enode PCM to AAC

```go
package main

import (
	"fmt"

	opus "github.com/qrtc/opus-go"
)

func main() {
	encoder, err := opus.CreateOpusEncoder(&opus.OpusEncoderConfig{
		SampleRate:  48000,
		MaxChannels: 2,
		Application: opus.AppVoIP,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		encoder.Close()
	}()

	inBuf := []byte{
		// PCM bytes
	}
	outBuf := make([]byte, 4096)

	n, err := encoder.Encode(inBuf, outBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outBuf[0:n])
}
```

# Dependencies

* opus
