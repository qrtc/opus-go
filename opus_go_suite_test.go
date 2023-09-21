package opus

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOpusGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OpusGo Suite")
}
