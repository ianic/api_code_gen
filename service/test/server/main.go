package main

import (
	"github.com/ianic/api_code_gen/service"
	"github.com/ianic/api_code_gen/service/api/nsq"

	"github.com/minus5/svckit/signal"
)

func main() {
	srv := nsq.NewServer(service.New())

	signal.WaitForInterupt()
	srv.Close()
}
