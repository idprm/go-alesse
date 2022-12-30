package repository

import (
	"context"

	"github.com/idprm/go-alesse/src/pkg/model"
	"gorm.io/gorm"
)

type IChatRepository interface {
	FindOneByMsisdn(ctx context.Context, msisdn string) ([]model.Chat, error)
	FindOneById(ctx context.Context, id uint64) (model.Chat, error)
	Add(ctx context.Context, chat model.Chat) (model.Chat, error)
	Update(ctx context.Context, chat model.Chat, id int) (model.Chat, error)
	Delete(ctx context.Context, chat model.Chat, id int) (model.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *chatRepository {
	return &chatRepository{db: db}
}

// func RepositoryChat(db *gorm.DB) *repository {
// 	return &repository{db}
// }
