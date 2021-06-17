package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	defer cancel()

	// listen os signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("linux cancel")
		cancel()
	}()

	// serve s1
	s1 := NewServer(":2222", "s1")
	group.Go(func() error {
		fmt.Println(s1.Name, "start ...")
		return s1.Serve(ctx, group)
	})

	// serve s2
	s2 := NewServer(":3333", "s2")
	group.Go(func() error {
		fmt.Println(s2.Name, "start ...")
		return s2.Serve(ctx, group)
	})

	// wait error
	if err := group.Wait(); err != nil {
		fmt.Println("error", err)
	}
}
