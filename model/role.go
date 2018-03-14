package model

type Role struct {
	ID        string
	Name      string
	CreatedAt string `db:"created_at"`
}
