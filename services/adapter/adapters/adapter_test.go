package adapters

import (
	"github.com/golang/mock/gomock"
	"github.com/kucherenkovova/marco-polo/proto"
	"github.com/kucherenkovova/marco-polo/services/adapter/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"time"
)

var _ = Describe("Test Marco Polo adapter", func() {

	mockCtrl := gomock.NewController(GinkgoT())
	forwarderMock := mocks.NewMockForwarder(mockCtrl)

	It("should fail if ServerClient is nil", func() {
		a, err := NewAdapter(nil)
		Expect(err).ToNot(BeNil())
		Expect(err).To(Equal(MissingSCErr))
		Expect(a).To(BeNil())
	})

	It("DictInLookupAdapter should map marco -> monkey", func() {
		a, _ := NewDictInLookupAdapter(forwarderMock)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		forwarderMock.EXPECT().Forward(ctx, gomock.Eq(&proto.Phrase{Word: "monkey"}))
		a.Forward(ctx, &proto.Phrase{Word: "marco"})
	})

	It("DictInLookupAdapter should map polo -> follow", func() {
		a, _ := NewDictInLookupAdapter(forwarderMock)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		forwarderMock.EXPECT().Forward(ctx, gomock.Eq(&proto.Phrase{Word: "follow"}))
		a.Forward(ctx, &proto.Phrase{Word: "polo"})
	})

	It("DictInLookupAdapter should return an error if no key found", func() {
		a, _ := NewDictInLookupAdapter(forwarderMock)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		_, err := a.Forward(ctx, &proto.Phrase{Word: "yolo"})
		Expect(err).ToNot(BeNil())
	})
})
