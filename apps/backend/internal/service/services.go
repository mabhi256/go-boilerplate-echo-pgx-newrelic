package service

import (
	"github.com/mabhi256/go-boilerplate-echo-pgx-newrelic/internal/lib/job"
	"github.com/mabhi256/go-boilerplate-echo-pgx-newrelic/internal/repository"
	"github.com/mabhi256/go-boilerplate-echo-pgx-newrelic/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
