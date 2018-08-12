package adapters

import (
	"fmt"
	"github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
	"log"
)

type DictInLookupAdapter struct {
	Forwarder
	Dict map[string]string
}

func (a *DictInLookupAdapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	word, ok := a.Dict[phrase.Word]
	if !ok {
		log.Printf("missing word in Dictionary: %s", phrase.Word)
		return nil, fmt.Errorf("missing word '%s' in lookup", phrase.Word)
	}
	return a.Forwarder.Forward(ctx, &proto.Phrase{Word: word})
}

func NewDictInLookupAdapter(f Forwarder, d Dictionary) (*DictInLookupAdapter, error) {
	return &DictInLookupAdapter{
		Dict:      d,
		Forwarder: f,
	}, nil
}
