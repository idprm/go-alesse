package model

type Medicine struct {
	ID             uint64       `gorm:"primaryKey" json:"id"`
	HealthcenterID uint         `json:"healthcenter_id"`
	Healthcenter   Healthcenter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category       string       `gorm:"size:20" json:"category"`
	Code           string       `gorm:"size:20" json:"code"`
	Type           string       `gorm:"size:25" json:"type"`
	Merk           string       `gorm:"size:200" json:"merk"`
	Name           string       `gorm:"size:200;not null" json:"name"`
	Unit           string       `gorm:"size:25" json:"unit"`
	IsActive       bool         `gorm:"type:bool;default:true" json:"is_active"`
}
