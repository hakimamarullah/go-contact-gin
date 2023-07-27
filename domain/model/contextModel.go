package model

import "context"

type ContextTimeout struct {
	CtxTimeout context.Context
	CtxCancel  context.CancelFunc
}
