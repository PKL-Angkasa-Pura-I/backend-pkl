package service

import (
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/adapter"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/config"
)

type svc struct {
	c    config.Config
	repo adapter.AdapterRepository
}

func NewService(repo adapter.AdapterRepository, c config.Config) adapter.AdapterService {
	return &svc{
		repo: repo,
		c:    c,
	}
}
