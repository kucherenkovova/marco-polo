package adapter

import (
	"github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
	"log"
)

type MarcoPoloAdapter struct {
	InDict       map[string]string
	OutDict      map[string]string
	ServerClient proto.ServerClient
}

func (a *MarcoPoloAdapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	word, ok := a.InDict[phrase.Word]
	if !ok {
		log.Printf("missing word in dict: %s", phrase.Word)
	}
	res, err := a.ServerClient.Send(ctx, &proto.Phrase{Word: word})
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return &proto.Phrase{Word: a.OutDict[res.Word]}, nil
}

func NewMarcoPoloAdapter(sc proto.ServerClient) *MarcoPoloAdapter {
	return &MarcoPoloAdapter{
		InDict: map[string]string{
			"marco": "monkey",
			"polo":  "follow",
		},
		OutDict: map[string]string{
			"monkey": "marco",
			"follow": "polo",
		},
		ServerClient: sc,
	}
}
