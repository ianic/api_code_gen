package main

import (
	"context"
	"fmt"

	"github.com/ianic/api_code_gen/service/api"
	"github.com/ianic/api_code_gen/service/api/nsq"
	"github.com/minus5/svckit/log"
)

func main() {
	log.Discard()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := nsq.NewClient()
	defer c.Close()

	add := func(x, y int) {
		fmt.Printf("%d + %d = ", x, y)
		rsp, err := c.Add(ctx, api.TwoReq{X: x, Y: y})
		showError(err)
		if err == nil {
			fmt.Printf("%d\n", rsp.Z)
		}
	}
	multiply := func(x, y int) {
		fmt.Printf("%d * %d = ", x, y)
		rsp, err := c.Multiply(ctx, api.TwoReq{X: x, Y: y})
		showError(err)
		if err == nil {
			fmt.Printf("%d\n", rsp.Z)
		}
	}
	multiply2 := func(x, y int) {
		fmt.Printf("%d * %d = ", x, y)
		rsp, err := c.Multiply2(ctx, api.TwoReq{X: x, Y: y})
		showError(err)
		if err == nil {
			fmt.Printf("%d\n", *rsp)
		}
	}
	cube := func(x int) {
		fmt.Printf("%d^2 = ", x)
		rsp, err := c.Cube(ctx, x)
		showError(err)
		if err == nil {
			fmt.Printf("%d\n", *rsp)
		}
	}

	add(2, 3)
	add(128, 129)
	multiply(2, 3)
	multiply(64, 3)
	multiply2(4, 3)
	cube(12)
	cube(15)
}

func showError(err error) {
	if err == nil {
		return
	}
	if err == api.ErrOverflow {
		fmt.Printf("owerflow\n")
		return
	}
	if err == context.Canceled || err == context.DeadlineExceeded {
		fmt.Printf("%s\n", err)
		return
	}
	panic(err)
}
