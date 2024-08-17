package common

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Greeter interface {
	Greet() (string, error)
}

type GreeterRPC struct {
	client *rpc.Client
}

func (g *GreeterRPC) Greet() (string, error) {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		return "", err
	}
	return resp, nil
}

type GreeterPRCServer struct {
	Impl Greeter
}

func (s *GreeterPRCServer) Greet(args interface{}, resp *string) error {
	var err error
	*resp, err = s.Impl.Greet()
	return err
}

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  0,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

type GreeterPlugin struct {
	Impl Greeter
}

func (_ GreeterPlugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}

func (p GreeterPlugin) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return &GreeterPRCServer{Impl: p.Impl}, nil
}
