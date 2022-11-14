package model

type SuperAdmin struct {
	ID       uint   `gorm:"primaryKey"`
	Phone    string `gorm:"size:15;unique" json:"phone"`
	Password string `gorm:"size:45" json:"password"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
}
