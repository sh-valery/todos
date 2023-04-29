package todo

import (
	"context"
	"go.uber.org/zap"
)

type UUIDGenerator interface {
	New() string
}

type Repository interface {
	FindByUser(UserID string) ([]*Todo, error)
	Find(ID string) (*Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(ID string) error
}

type Service interface {
	GetTodos(ctx context.Context, UserID string) ([]*Todo, error)
	GetTodoByID(ctx context.Context, ID string, UserID string) (*Todo, error)
	Create(ctx context.Context, todo *Todo) error
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, ID string) error
}

func NewTodoService(repo Repository, uuid UUIDGenerator, logger *zap.Logger) Service {
	if logger == nil {
		logger = zap.NewNop()
	}

	return &ServiceImpl{
		repo:   repo,
		uuid:   uuid,
		logger: logger,
	}
}
