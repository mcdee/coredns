package ens

import (
	"fmt"
	"time"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/ens/registrycontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/net/context"

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
	registryContract, err := RegistryContract(client)
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

func RegistryContractAddress(client *ethclient.Client) (address common.Address, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return
	}

	// Instantiate the registry contract
	if chainID.Cmp(params.MainnetChainConfig.ChainID) == 0 {
		address = common.HexToAddress("314159265dd8dbb310642f98f50c066173c1259b")
	} else if chainID.Cmp(params.TestnetChainConfig.ChainID) == 0 {
		address = common.HexToAddress("112234455c3a32fd11230c42e7bccd4a84e02010")
	} else if chainID.Cmp(params.RinkebyChainConfig.ChainID) == 0 {
		address = common.HexToAddress("e7410170f87102DF0055eB195163A03B7F2Bff4A")
	} else {
		err = fmt.Errorf("No contract for network ID %v", chainID)
	}
	return
}

// RegistryContract obtains the registry contract for a chain
func RegistryContract(client *ethclient.Client) (registry *registrycontract.RegistryContract, err error) {
	var address common.Address
	address, err = RegistryContractAddress(client)
	if err != nil {
		return
	}

	// Instantiate the registry contract
	registry, err = registrycontract.NewRegistryContract(address, client)

	return
}
