package model

type Courier struct {
	ID             uint         `gorm:"primaryKey"`
	HealthcenterID uint         `json:"-"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name           string       `gorm:"size:100;not null" json:"name"`
	Photo          string       `gorm:"size:150;not null" json:"photo"`
	Phone          string       `gorm:"size:15" json:"phone"`
	IsActive       bool         `gorm:"type:bool" json:"is_active"`
}
