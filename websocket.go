package handlerlib

import "time"

type WebSocketModel interface {
	Read() (int, []byte, error)
	SetPingHandler(func(string) error) error
	SetCloseHandler(f func(code int, text string) error) error
	WriteControl(int, []byte, time.Time) error
	Write(int, []byte) error
	WriteJson(interface{}) error
	SetWriteDeadLine(deadline time.Duration)
	SetReadDeadLine(deadline time.Duration)
	Close()
}
