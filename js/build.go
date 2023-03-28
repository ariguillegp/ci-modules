package js

import (
	"context"
	"fmt"
)

func (p *Pipeline) Build(ctx context.Context, baseNodeImg string) error {
	src := p.Client.Host().Directory(".") // get the project source directory

	npm := p.Client.Container().From(baseNodeImg). // initialize new container from npm image
							WithMountedDirectory("/src", src).WithWorkdir("/src"). // mount source directory to /src
							WithExec([]string{"npm", "install"}).                  // execute npm install
							WithExec([]string{"npm", "run", "test"})               // execute npm test command

	// get test output
	test, err := npm.Stdout(ctx)
	if err != nil {
		return err
	}
	// print output to console
	fmt.Println(test)

	// execute build command and get build output
	build, err := npm.WithExec([]string{"npm", "run", "build"}).Stdout(ctx)
	if err != nil {
		return err
	}
	// print output to console
	fmt.Println(build)

	// all good
	return nil
}
