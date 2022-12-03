package model

type Status struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:25;unique" json:"name"`
	ValueSystem string `gorm:"size:200" json:"value_system"`
	ValueUser   string `gorm:"size:250" json:"value_user"`
	ValueNotif  string `gorm:"size:250" json:"value_notif"`
}
