package models

type User struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GOM setting schema
func (User) TableName() string {
	return "develop.users"
}
