package adapters

import (
	"errors"
	"github.com/kucherenkovova/marco-polo.git/proto"
	"golang.org/x/net/context"
)

type Dictionary map[string]string

type Forwarder interface {
	Forward(context.Context, *proto.Phrase) (*proto.Phrase, error)
}

type adapter struct {
	serverClient proto.ServerClient
}

func (a *adapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	res, err := a.serverClient.Send(ctx, phrase)
	return res, err
}

func NewAdapter(sc proto.ServerClient) (Forwarder, error) {
	if sc == nil {
		return nil, errors.New("Can't establish connection to Server. ServerClient is nil.")
	}
	return &adapter{serverClient: sc}, nil
}
