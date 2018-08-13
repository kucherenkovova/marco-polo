package adapters

import (
	"fmt"
	"github.com/kucherenkovova/marco-polo.git/proto"
	"golang.org/x/net/context"
	"log"
)

type dictInLookupAdapter struct {
	Forwarder
	inLookup map[string]string
}

func (a *dictInLookupAdapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	word, ok := a.inLookup[phrase.Word]
	if !ok {
		log.Printf("missing word in Dictionary: %s", phrase.Word)
		return nil, fmt.Errorf("missing word '%s' in lookup", phrase.Word)
	}
	return a.Forwarder.Forward(ctx, &proto.Phrase{Word: word})
}

func NewDictInLookupAdapter(f Forwarder, d Dictionary) (Forwarder, error) {
	return &dictInLookupAdapter{
		inLookup:  d,
		Forwarder: f,
	}, nil
}
