package sql

import (
	"context"
	"fmt"

	"github.com/idprm/go-alesse/src/common"
	"github.com/idprm/go-alesse/src/pkg/domain/model"
	"gorm.io/gorm"
)

type ChatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	r := &ChatRepository{
		DB: db,
	}
	return r
}

func (r ChatRepository) Save(ctx context.Context, chat *model.Chat) error {
	if err := r.DB.Save(chat).Find(&chat).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) FindByNumber(ctx context.Context, number string) (*model.Chat, error) {
	var chats model.Chat
	req := r.DB.Where("number = ?", number).First(&chats)

	if req.Error == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("chat %s %w", number, common.ErrNotFound)
	}

	if req.Error != nil {
		return nil, nil
	}

	return &chats, nil
}
