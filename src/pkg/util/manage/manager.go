package manage

import "github.com/idprm/go-alesse/src/pkg/util/localconfig"

// Manager handle business logic related to payment gateway
type Manager struct {
	config localconfig.Config
	secret localconfig.Secret
}

func NewManager(
	config localconfig.Config,
	secret localconfig.Secret,
) *Manager {
	return &Manager{
		config: config,
		secret: secret,
	}
}
