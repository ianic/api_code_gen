package main

import (
	"fmt"

	"github.com/ianic/api_code_gen/service/api"
	"github.com/ianic/api_code_gen/service/api/nsq"
	"github.com/minus5/svckit/log"
)

var c *api.Client

func main() {
	log.Discard()
	c = nsq.NewClient()

	add(2, 3)
	add(128, 129)

	multiply(2, 3)
	multiply(64, 3)

	multiply2(4, 3)

	cube(12)
	cube(15)

	c.Close()
}

func add(x, y int) {
	rsp, err := c.Add(api.TwoReq{X: x, Y: y})
	if err == api.ErrOverflow {
		fmt.Printf("%d + %d = owerflow\n", x, y)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d + %d = %d\n", x, y, rsp.Z)
}

func multiply(x, y int) {
	rsp, err := c.Multiply(api.TwoReq{X: x, Y: y})
	if err == api.ErrOverflow {
		fmt.Printf("%d * %d = owerflow\n", x, y)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d * %d = %d\n", x, y, rsp.Z)
}

func multiply2(x, y int) {
	rsp, err := c.Multiply2(api.TwoReq{X: x, Y: y})
	if err == api.ErrOverflow {
		fmt.Printf("%d * %d = owerflow\n", x, y)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d * %d = %d\n", x, y, *rsp)
}

func cube(x int) {
	rsp, err := c.Cube(x)
	if err == api.ErrOverflow {
		fmt.Printf("%d^2 = owerflow\n", x)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d^2 = %d\n", x, *rsp)
}
