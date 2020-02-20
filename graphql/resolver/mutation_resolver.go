package resolver

import (
	"context"
)

func (r *TodoMutationResolver) Test(ctx context.Context) *string {
	x := "test"
	return &x
}

