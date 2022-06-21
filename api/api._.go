package api

type ZincConfig struct {
	Address  string
	User     string
	Password string
}

type ZincApi struct {
	Address  string
	User     string
	Password string
}

// NewZincApi returns a model for zinc connect.
func NewZincApi(cfg ZincConfig) *ZincApi {
	return &ZincApi{
		Address:  cfg.Address,
		User:     cfg.User,
		Password: cfg.Password,
	}
}
