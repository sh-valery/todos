package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/sh-valery/todos/internal/todo"
	"math/big"
	"time"

	"github.com/sh-valery/todos/internal/transport/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	t := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	tt := &todo.Todo{
		ID:        t.ID,
		Text:      t.Text,
		Done:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		OwnerID:   "",
	}
	r.srv.Create(ctx, tt)
	return t, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.srv.GetTodos(ctx, "")
	if err != nil {
		return nil, err
	}
	var todoQL []*model.Todo
	for _, v := range todos {
		todoQL = append(todoQL, &model.Todo{
			ID:   v.ID,
			Text: v.Text,
			Done: v.Done,
			User: &model.User{
				ID:   v.OwnerID,
				Name: "",
			},
		})
	}
	return todoQL, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
