package dtl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDtl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dtl Suite")
}
