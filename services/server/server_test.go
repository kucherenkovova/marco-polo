package main

import (
	"context"

	"github.com/kucherenkovova/marco-polo.git/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	dict := map[string]string{
		"monkey": "follow",
		"follow": "monkey",
	}
	mps := &MarcoPoloServer{dict}
	It("should return follow for monkey", func() {
		phrase := &proto.Phrase{Word: "monkey"}
		result, err := mps.Send(context.Background(), phrase)
		Expect(err).To(BeNil())
		Expect(result.GetWord()).To(Equal("follow"))
	})
	It("should return monkey for follow", func() {
		phrase := &proto.Phrase{Word: "follow"}
		result, err := mps.Send(context.Background(), phrase)
		Expect(err).To(BeNil())
		Expect(result.GetWord()).To(Equal("monkey"))
	})
	It("should return empty string for random word", func() {
		phrase := &proto.Phrase{Word: "random word"}
		result, err := mps.Send(context.Background(), phrase)
		Expect(err).To(BeNil())
		Expect(result.GetWord()).To(Equal(""))
	})
})
