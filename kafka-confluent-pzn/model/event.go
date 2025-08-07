package model

type Event interface {
	GetId() string // ini bakal berperan sebagai key untuk message yang dikirim ke topic
}