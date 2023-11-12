package entity

type Session struct {
	ID      string `db:"sessionId"`
	Code    string `db:"code"`
	Attempt int    `db:"attempt"`
}
