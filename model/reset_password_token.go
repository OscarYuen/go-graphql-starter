package model

type ResetPasswordToken struct {
	ID        string
	UserID    int64  `db:"user_id"`
	CreatedAt string `db:"created_at"`
}
