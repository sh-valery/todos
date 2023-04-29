package todo

import (
	"context"
	"go.uber.org/zap"
)

type ServiceImpl struct {
	repo Repository
	uuid UUIDGenerator
	// infra layer
	logger *zap.Logger
}

func (s ServiceImpl) GetTodos(ctx context.Context, UserID string) ([]*Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) GetTodoByID(ctx context.Context, ID string, UserID string) (*Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) Create(ctx context.Context, todo *Todo) error {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) Update(ctx context.Context, todo *Todo) error {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) Delete(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}
