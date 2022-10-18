package server

import "gorm.io/gorm"

type Server struct {
}

func migrate(db *gorm.DB) {
	db.AutoMigrate()
}

func NewServer()
