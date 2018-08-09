package adapter

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMarcoPolo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MarcoPolo Suite")
}
