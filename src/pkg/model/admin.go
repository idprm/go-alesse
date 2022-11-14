package model

type Admin struct {
	ID             uint         `gorm:"primaryKey"`
	HealthcenterID uint         `json:"-"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Phone          string       `gorm:"size:15;unique" json:"phone"`
	Password       string       `gorm:"size:45" json:"password"`
	IsActive       bool         `gorm:"type:bool" json:"is_active"`
}
