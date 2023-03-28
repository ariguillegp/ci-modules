package js

import (
	"context"
	"fmt"
)

func (p *Pipeline) Test(ctx context.Context, baseNodeImg string) error {
	fmt.Println("Test Job")
	return nil
}
