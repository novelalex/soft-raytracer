package camera_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCamera(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Camera Suite")
}
