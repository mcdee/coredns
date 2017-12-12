package ens

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/ethereal/ens"

	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("ens", caddy.Plugin{
		ServerType: "dns",
		Action:     setupENS,
	})
}

func setupENS(c *caddy.Controller) error {

	connection, err := ensParse(c)
	if err != nil {
		return plugin.Error("ens", err)
	}

	client, err := ethclient.Dial(connection)
	if err != nil {
		return plugin.Error("ens", err)
	}

	// Obtain the registry contract
	registryContract, err := ens.RegistryContract(client)
	if err != nil {
		return plugin.Error("ens", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return ENS{Next: next,
			Client:   client,
			Registry: registryContract}
	})

	return nil
}

func ensParse(c *caddy.Controller) (string, error) {
	var connection string
	c.Next()
	for c.NextBlock() {
		switch c.Val() {
		case "connection":
			args := c.RemainingArgs()
			if len(args) == 0 {
				return "", c.Errf("invalid connection; no value")
			}

			if len(args) > 1 {
				return "", c.Errf("invalid connection; multiple values")
			}
			connection = args[0]
		default:
			return "", c.Errf("unknown value %v", c.Val())
		}
	}
	if connection == "" {
		return "", c.Errf("no connection")
	}
	return connection, nil
}
