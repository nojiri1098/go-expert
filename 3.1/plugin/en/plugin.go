package main

import (
	"github.com/hashicorp/go-plugin"

	"nojiri1098/go-expert/3.1/common"
)

type Greeter struct{}

func (g Greeter) Greet() (string, error) {
	return "hello!", nil
}

func main() {
	var greeter Greeter
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"greeter": &common.GreeterPlugin{Impl: greeter},
		},
	})
}
