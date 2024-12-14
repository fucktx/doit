package callbacks

import "context"

type Callback interface {
	Handler(ctx context.Context, args ...any) error
}
