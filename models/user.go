package models

type User struct {
	Id uint			`json:"id" gorm:"primary_key;column:user_id"`
	Username string	`json:"username"`
	Email string	`json:"email"`
	Password string	`json:"password"`
	JoinedAt string	`json:"joinedAt" gorm:"column:created_at;<-:false"`
}
