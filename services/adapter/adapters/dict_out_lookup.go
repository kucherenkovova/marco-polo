package adapters

import (
	"github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
)

type DictOutLookupAdapter struct {
	Forwarder
	Dict map[string]string
}

func (a *DictOutLookupAdapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	res, err := a.Forwarder.Forward(ctx, phrase)
	if err != nil {
		return nil, err
	}
	return &proto.Phrase{Word: a.Dict[res.Word]}, err
}

func NewDictOutLookupAdapter(f Forwarder, d Dictionary) (*DictOutLookupAdapter, error) {
	return &DictOutLookupAdapter{
		Dict:      d,
		Forwarder: f,
	}, nil
}
