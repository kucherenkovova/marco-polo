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
		log.Printf("missing word in dict: %s", phrase.Word)
		return nil, fmt.Errorf("missing word '%s' in lookup", phrase.Word)
	}
	return a.Forwarder.Forward(ctx, &proto.Phrase{Word: word})
}

func NewDictInLookupAdapter(f Forwarder) (*DictInLookupAdapter, error) {
	return &DictInLookupAdapter{
		Dict: map[string]string{
			"marco": "monkey",
			"polo":  "follow",
		},
		Forwarder: f,
	}, nil
}
