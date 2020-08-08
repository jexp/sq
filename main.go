// Package main contains sq's main function.
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/neilotoole/sq/cli"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)

		<-stop
		cancelFn()
	}()

	err := cli.Execute(ctx, os.Stdin, os.Stdout, os.Stderr, os.Args)
	if err != nil {
		cancelFn()
		os.Exit(1)
	}
}
