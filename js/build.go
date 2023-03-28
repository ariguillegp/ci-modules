package js

import (
	"context"
	"fmt"
)

func (p *Pipeline) Build(ctx context.Context, baseNodeImg string) error {
	fmt.Println("Build Job")
	return nil
}
