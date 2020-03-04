package factory

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Service interface {
	MakeBox(ctx context.Context, name string, height int32, width int32, depth int32) (int32, int32, error)
	GetBoxes(ctx context.Context, page int32, perPage int32) ([]BoxSpecification, error)
	Status(ctx context.Context, page int32, perPage int32) (string, bool, error)
	Log() log.Logger
}

type serviceImpl struct {
	log  log.Logger
	repo Repository
}

func NewService(log log.Logger, repo Repository) Service {
	return &serviceImpl{log: log, repo: repo}
}

func (s *serviceImpl) Log() log.Logger {
	return s.log
}

func (s *serviceImpl) MakeBox(ctx context.Context, name string, height int32, width int32, depth int32) (int32, int32, error) {
	resPage, resPerPage, err := s.repo.MakeBox(ctx, name, height, width, depth)
	if err != nil {
		return 0, 0, err
	}
	return resPage, resPerPage, nil
}

func (s *serviceImpl) GetBoxes(ctx context.Context, page int32, perPage int32) ([]BoxSpecification, error) {
	resBoxes, err := s.repo.GetBoxes(ctx, page, perPage)
	if err != nil {
		return nil, err
	}
	return resBoxes, nil
}

func (s *serviceImpl) Status(ctx context.Context, page int32, perPage int32) (string, bool, error) {
	resServiceName, resOk, err := s.repo.Status(ctx, page, perPage)
	if err != nil {
		return "", false, err
	}
	return resServiceName, resOk, nil
}
