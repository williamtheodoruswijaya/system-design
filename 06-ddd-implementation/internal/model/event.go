package model

type Event interface {
	GetId() string // berguna untuk jadi key pada saat producer mengirim pesan
}
