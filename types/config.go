package types

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Config struct {
	*sdk.Config

	prefixes map[string]string
	sealed   bool
	mtx      sync.Mutex
}

var (
	config = func() *Config {
		config := Config{
			Config: sdk.GetConfig(),
			prefixes: map[string]string{
				"provider_addr": Bech32PrefixProvAddr,
				"node_addr":     Bech32PrefixNodeAddr,
				"provider_pub":  Bech32PrefixProvPub,
				"node_pub":      Bech32PrefixNodePub,
			},
		}

		config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
		config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
		config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)

		return &config
	}()
)

func GetConfig() *Config {
	return config
}

func (c *Config) assert() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.sealed {
		panic("config is sealed")
	}
}

func (c *Config) Seal() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.Config.Seal()
	c.sealed = true
}

func (c *Config) SetBech32PrefixForProvider(addr, pub string) {
	c.assert()

	config.prefixes["provider_addr"] = addr
	config.prefixes["provider_pub"] = pub
}

func (c *Config) SetBech32PrefixForNode(addr, pub string) {
	c.assert()

	config.prefixes["node_addr"] = addr
	config.prefixes["node_pub"] = pub
}

func (c *Config) GetBech32ProviderAddrPrefix() string {
	return c.prefixes["provider_addr"]
}

func (c *Config) GetBech32ProviderPubPrefix() string {
	return c.prefixes["provider_pub"]
}

func (c *Config) GetBech32NodeAddrPrefix() string {
	return c.prefixes["node_addr"]
}

func (c *Config) GetBech32NodePubPrefix() string {
	return c.prefixes["node_pub"]
}
