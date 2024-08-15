package env

import (
	config "user/internal/config"
	"errors"
	"fmt"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
	dataBaseUrl = "PG_URL"
)

type grpcConfig struct {
	host string
	port string
	pgUrl string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}
	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}
	pgUrl := os.Getenv(dataBaseUrl)
	fmt.Println(pgUrl)
	if len(pgUrl) == 0 {
		return nil, errors.New("db url not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
		pgUrl: pgUrl,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *grpcConfig) GetDbUrl() string {
	return cfg.pgUrl
}