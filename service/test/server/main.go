package main

import (
	"context"

	"github.com/ianic/api_code_gen/service"
	"github.com/ianic/api_code_gen/service/api/nsq"

	"github.com/minus5/svckit/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	srv := nsq.Server(ctx, service.New())
	defer srv.Close()
	defer cancel()

	signal.WaitForInterupt()
}
