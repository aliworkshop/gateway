package gateway

import (
	"context"
	"time"
)

type WebSocketHandler interface {
	Read(context.Context) (int, []byte, error)
	Write(context.Context, int, []byte) error
	WriteJson(context.Context, interface{}) error
	SetWriteDeadLine(deadline time.Duration) error
	SetReadDeadLine(deadline time.Duration) error
	Ping(context.Context) error
	Close()
}