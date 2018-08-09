package adapters

import (
	"errors"
	"github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
)

type Forwarder interface {
	Forward(context.Context, *proto.Phrase) (*proto.Phrase, error)
}

var MissingSCErr = errors.New("Can't establish connection to Server. ServerClient is nil.")

type Adapter struct {
	ServerClient proto.ServerClient
}

func (a *Adapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	res, err := a.ServerClient.Send(ctx, phrase)
	return res, err
}

func NewAdapter(sc proto.ServerClient) (*Adapter, error) {
	if sc == nil {
		return nil, MissingSCErr
	}
	return &Adapter{ServerClient: sc}, nil
}
