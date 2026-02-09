package nmath_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNmath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nmath Suite")
}
