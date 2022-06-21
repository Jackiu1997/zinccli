package zinccli

import "zinccli/api"

type (
	ZincConfig api.ZincConfig

	ZincClient struct {
		*api.ZincApi
		Config ZincConfig
	}
)

func NewZincClient(cfg ZincConfig) *ZincClient {
	api.SetRequestConfig(&api.ZincConfig{
		Address:  cfg.Address,
		User:     cfg.User,
		Password: cfg.Password,
	})

	return &ZincClient{
		Config: cfg,
	}
}
