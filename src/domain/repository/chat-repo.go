package repository

import "github.com/idprm/go-alesse/src/domain/model"

type ChatRepository interface {
	Get(id int) (*model.Chat, error)
	GetAll() ([]model.Chat, error)
	Save(*model.Chat) error
	Remove(id int) error
	Update(*model.Chat) error
}
