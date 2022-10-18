package datastore

import (
	"context"

	"github.com/idprm/go-alesse/src/pkg/domain/model"
)

type UserRepository interface {
}

type ChatRepository interface {
	FindByNumber(ctx context.Context, number string) (*model.Chat, error)
	Save(ctx context.Context, chat *model.Chat) error
}
