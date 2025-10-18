package types

import (
	"sync"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestConfig_GetBech32NodeAddrPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_addr": "sent"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sent",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_addr": "sentnode"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sentnode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32NodeAddrPrefix(); got != tt.want {
				t.Errorf("GetBech32NodeAddrPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetBech32NodePubPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_pub": "sentpub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sentpub",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"node_pub": "sentnodepub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sentnodepub",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32NodePubPrefix(); got != tt.want {
				t.Errorf("GetBech32NodePubPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetBech32ProviderAddrPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_addr": "sent"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sent",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_addr": "sentprov"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sentprov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32ProviderAddrPrefix(); got != tt.want {
				t.Errorf("GetBech32ProviderAddrPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetBech32ProviderPubPrefix(t *testing.T) {
	type fields struct {
		Config   *sdk.Config
		prefixes map[string]string
		sealed   bool
		mtx      sync.Mutex
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_pub": "sentpub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sentpub",
		},
		{
			"valid prefix",
			fields{
				Config:   sdk.GetConfig(),
				prefixes: map[string]string{"provider_pub": "sentprovpub"},
				sealed:   false,
				mtx:      sync.Mutex{},
			},
			"sentprovpub",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Config:   tt.fields.Config,
				prefixes: tt.fields.prefixes,
				sealed:   tt.fields.sealed,
				mtx:      tt.fields.mtx,
			}
			if got := c.GetBech32ProviderPubPrefix(); got != tt.want {
				t.Errorf("GetBech32ProviderPubPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
