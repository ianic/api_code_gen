package main

import (
	"fmt"

	"github.com/ianic/api_code_gen/service/api"
	"github.com/ianic/api_code_gen/service/api/nsq"
	"github.com/ianic/api_code_gen/service/dto"
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
	rsp, err := c.Add(dto.TwoReq{X: x, Y: y})
	if err == dto.ErrOverflow {
		fmt.Printf("%d + %d = owerflow\n", x, y)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d + %d = %d\n", x, y, rsp.Z)
}

func multiply(x, y int) {
	rsp, err := c.Multiply(dto.TwoReq{X: x, Y: y})
	if err == dto.ErrOverflow {
		fmt.Printf("%d * %d = owerflow\n", x, y)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d * %d = %d\n", x, y, rsp.Z)
}

func multiply2(x, y int) {
	rsp, err := c.Multiply2(dto.TwoReq{X: x, Y: y})
	if err == dto.ErrOverflow {
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
	if err == dto.ErrOverflow {
		fmt.Printf("%d^2 = owerflow\n", x)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d^2 = %d\n", x, *rsp)
}
