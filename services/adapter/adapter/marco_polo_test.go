package adapter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Marco Polo adapter", func() {

	Context("MarcoPoloAdapter", func() {
		It("should fail if ServerClient is nil", func() {
			a, err := NewMarcoPoloAdapter(nil)
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(MissingSCErr))
			Expect(a).To(BeNil())
		})
	})
})
