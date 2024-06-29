package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/black-pepper-team/community-indexer/internal/service/core"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	coreCtxKey
	mockAPICtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxCore(valud *core.Core) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, coreCtxKey, valud)
	}
}

func Core(r *http.Request) *core.Core {
	return r.Context().Value(coreCtxKey).(*core.Core)
}

func CtxMockAPI(value bool) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mockAPICtxKey, value)
	}
}

func MockAPI(r *http.Request) bool {
	return r.Context().Value(mockAPICtxKey).(bool)
}
