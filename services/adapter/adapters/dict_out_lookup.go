package adapters

import (
	"github.com/kucherenkovova/marco-polo.git/proto"
	"golang.org/x/net/context"
)

type dictOutLookupAdapter struct {
	Forwarder
	outLookup Dictionary
}

func (a *dictOutLookupAdapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	res, err := a.Forwarder.Forward(ctx, phrase)
	if err != nil {
		return nil, err
	}
	return &proto.Phrase{Word: a.outLookup[res.Word]}, err
}

func NewDictOutLookupAdapter(f Forwarder, d Dictionary) (Forwarder, error) {
	return &dictOutLookupAdapter{
		outLookup: d,
		Forwarder: f,
	}, nil
}
