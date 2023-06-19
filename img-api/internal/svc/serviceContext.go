package svc

import (
	"auction_server/img-api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Path   string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
