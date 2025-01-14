package types

type UserInfo struct {
	Id       *int   `json:"id"`
	Email    string `json:"email"     db:"email"     validate:"required,email"`
	Password string `json:"password"  db:"password"  validate:"required"`
	Role     string `json:"role"      db:"role"      validate:"required"`
	IsActive bool   `json:"is_active" db:"is_active" `
}
