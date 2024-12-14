package callbacks

import "context"

type handler interface {
	Step(ctx context.Context, args ...any) error
}

//type CallbackFunc func(ctx context.Context, args ...any) error
//
//func (f CallbackFunc) Callback(ctx context.Context, args ...any) error {
//	return f(ctx, args...)
//}

type Callback interface {
	Handler(ctx context.Context, args ...any) error
}
