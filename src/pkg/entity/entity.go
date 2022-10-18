package entity

type entity struct {
}

type Entity interface {
}

func New() Entity {
	return &entity{}
}
