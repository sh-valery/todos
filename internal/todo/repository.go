package todo

import "errors"

type InMemoryRepository struct {
	todos []*Todo
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		todos: []*Todo{},
	}
}

func (i *InMemoryRepository) FindByUser(UserID string) ([]*Todo, error) {
	var results []*Todo
	for _, t := range i.todos {
		if t.OwnerID == UserID {
			results = append(results, t)
		}
	}
	return results, nil
}

func (i *InMemoryRepository) Find(ID string) (*Todo, error) {
	for _, t := range i.todos {
		if t.ID == ID {
			return t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (i *InMemoryRepository) Create(todo *Todo) error {
	i.todos = append(i.todos, todo)
	return nil
}

func (i *InMemoryRepository) Update(todo *Todo) error {
	for idx, t := range i.todos {
		if t.ID == todo.ID {
			i.todos[idx] = todo
			return nil
		}
	}
	return errors.New("todo not found")
}

func (i *InMemoryRepository) Delete(ID string) error {
	for idx, t := range i.todos {
		if t.ID == ID {
			i.todos = append(i.todos[:idx], i.todos[idx+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
