package model

type Status struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:50;unique" json:"name"`
	ValueSystem string `gorm:"size:300" json:"value_system"`
	ValueUser   string `gorm:"size:300" json:"value_user"`
	ValueNotif  string `gorm:"size:300" json:"value_notif"`
}
