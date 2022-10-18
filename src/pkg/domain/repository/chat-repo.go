package repository

import "github.com/idprm/go-alesse/src/pkg/domain/model"

type ChatRepository interface {
	Get(id int) (model.Chat, error)
	GetAll() ([]model.Chat, error)
	Add(model.Chat) error
	Remove(id int) error
	Update(model.Chat) error
}
