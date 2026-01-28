package models

type UserModel struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:150;unique;not null"`
}

func (UserModel) TableName() string {
	return "users"
}
