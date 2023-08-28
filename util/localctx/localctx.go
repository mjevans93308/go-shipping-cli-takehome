package localctx

import (
	"context"

	"go.uber.org/zap"
)

type Localctx struct {
	Lctx   *context.Context
	Logger *zap.SugaredLogger
}

func NewLocalCtx(ctx *context.Context, slog *zap.SugaredLogger) *Localctx {
	return &Localctx{
		Lctx:   ctx,
		Logger: slog,
	}
}
