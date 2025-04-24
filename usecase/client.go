package usecase

import (
	"context"
	"fullstacktest/entity"
	"fullstacktest/repository"
	"strings"
)

type ClientUsecaseItf interface {
	InsertClient(ctx context.Context, data *entity.MyClient) (*int, error)
	SelectAllClient(ctx context.Context) ([]entity.MyClient, error)
	UpdateClient(ctx context.Context, data *entity.MyClient) error
	DeleteClient(ctx context.Context, slug string) error
}

type ClientUsecaseImpl struct {
	cr repository.ClientRepoItf
}

func NewClientUsecase(cr repository.ClientRepoItf) ClientUsecaseImpl {
	return ClientUsecaseImpl{
		cr: cr,
	}
}

func (uc ClientUsecaseImpl) InsertClient(ctx context.Context, data *entity.MyClient) (*int, error) {
	data.Slug = strings.ToLower(data.Slug)

	id, err := uc.cr.InsertClient(ctx, data)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (uc ClientUsecaseImpl) SelectAllClient(ctx context.Context) ([]entity.MyClient, error) {
	clients, err := uc.cr.SelectAllClient(ctx)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (uc ClientUsecaseImpl) UpdateClient(ctx context.Context, data *entity.MyClient) error {
	data.Slug = strings.ToLower(data.Slug)

	err := uc.cr.UpdateClient(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (uc ClientUsecaseImpl) DeleteClient(ctx context.Context, slug string) error {
	slug = strings.ToLower(slug)

	err := uc.cr.DeleteClient(ctx, slug)
	if err != nil {
		return err
	}

	return nil
}
