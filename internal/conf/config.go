package conf

import (
	"github.com/dlc-01/BackendCrypto/pkg/logger"
	"time"
)

type Config struct {
	Config          string        `env:"CONF" envDefault:""`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"600s" json:"ShutdownTimeout"`
	GRPCServer      struct {
		Address string `env:"GRPC_SERVER_ADDRESS" envDefault:":3200" json:"GRPCAddress"`
		Network string `env:"GRPC_SERVER_NETWORK" envDefault:"tcp" json:"GRPCNetwork"`
	}
	DB struct {
		DSN    string `env:"DATABASE_DSN" envDefault:"postgresql://postgres:123456@localhost:5432/postgres" json:"DB_DSN"`
		Driver string `env:"DATABASE_DRIVER" envDefault:"pgx" json:"DB_Driver"`
	}
	Logger logger.ConfLogger
}
