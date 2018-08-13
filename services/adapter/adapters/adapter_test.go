package adapters

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kucherenkovova/marco-polo.git/proto"
	"github.com/kucherenkovova/marco-polo.git/services/adapter/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"time"
)

var _ = Describe("Test Marco Polo adapter", func() {

	mockCtrl := gomock.NewController(GinkgoT())
	forwarderMock := mocks.NewMockForwarder(mockCtrl)
	mockCtrl2 := gomock.NewController(GinkgoT())
	serverClientMock := mocks.NewMockServerClient(mockCtrl2)

	outDict := Dictionary{
		"monkey": "marco",
		"follow": "polo",
	}

	inDict := Dictionary{
		"marco": "monkey",
		"polo":  "follow",
	}

	It("should fail if ServerClient is nil", func() {
		a, err := NewAdapter(nil)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("Can't establish connection to Server. ServerClient is nil."))
		Expect(a).To(BeNil())
	})

	It("DictInLookupAdapter should map marco -> monkey", func() {
		a, _ := NewDictInLookupAdapter(forwarderMock, inDict)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		forwarderMock.EXPECT().Forward(ctx, gomock.Eq(&proto.Phrase{Word: "monkey"}))
		a.Forward(ctx, &proto.Phrase{Word: "marco"})
	})

	It("DictInLookupAdapter should map polo -> follow", func() {
		a, _ := NewDictInLookupAdapter(forwarderMock, inDict)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		forwarderMock.EXPECT().Forward(ctx, gomock.Eq(&proto.Phrase{Word: "follow"}))
		a.Forward(ctx, &proto.Phrase{Word: "polo"})
	})

	It("DictInLookupAdapter should return an error if no key found", func() {
		a, _ := NewDictInLookupAdapter(forwarderMock, inDict)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		_, err := a.Forward(ctx, &proto.Phrase{Word: "yolo"})
		Expect(err).ToNot(BeNil())
	})

	It("dictOutLookupAdapter should map follow -> polo", func() {
		adapter, _ := NewAdapter(serverClientMock)
		a, _ := NewDictOutLookupAdapter(adapter, outDict)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		serverClientMock.EXPECT().Send(ctx, gomock.Eq(&proto.Phrase{Word: "monkey"})).Return(&proto.Phrase{Word: "follow"}, nil)
		res, err := a.Forward(ctx, &proto.Phrase{Word: "monkey"})
		Expect(res.Word).To(Equal("polo"))
		Expect(err).To(BeNil())
	})
	It("dictOutLookupAdapter should map monkey -> marco", func() {
		adapter, _ := NewAdapter(serverClientMock)
		a, _ := NewDictOutLookupAdapter(adapter, outDict)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		serverClientMock.EXPECT().Send(ctx, gomock.Eq(&proto.Phrase{Word: "follow"})).Return(&proto.Phrase{Word: "monkey"}, nil)
		res, err := a.Forward(ctx, &proto.Phrase{Word: "follow"})
		Expect(res.Word).To(Equal("marco"))
		Expect(err).To(BeNil())
	})
	It("dictOutLookupAdapter should return error if no key found", func() {
		adapter, _ := NewAdapter(serverClientMock)
		a, _ := NewDictOutLookupAdapter(adapter, outDict)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		serverClientMock.EXPECT().Send(ctx, gomock.Any()).Return(nil, errors.New("No key found"))
		res, err := a.Forward(ctx, &proto.Phrase{Word: "yolo"})
		Expect(res).To(BeNil())
		Expect(err).NotTo(BeNil())
	})

})
