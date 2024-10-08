package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"nojiri1098/go-expert/3.1/common"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Warn,
	})

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"greeter": &common.GreeterPlugin{},
		},
		Cmd:    exec.Command("./ja"),
		Logger: logger,
	})
	defer client.Kill()

	rpcClienet, err := client.Client()
	if err != nil {
		exit(err)
	}

	raw, err := rpcClienet.Dispense("greeter")
	if err != nil {
		exit(err)
	}

	greeter := raw.(common.Greeter)
	resp, err := greeter.Greet()
	if err != nil {
		exit(err)
	}
	fmt.Println(resp)
}

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}
