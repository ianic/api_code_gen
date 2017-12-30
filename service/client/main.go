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

	call(2, 3)
	call(128, 129)

	c.Close()
}

func call(x, y int) {
	rsp, err := c.Add(dto.AddReq{X: x, Y: y})
	if err == dto.ErrOverflow {
		fmt.Printf("%d + %d = owerflow\n", x, y)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d + %d = %d\n", x, y, rsp.Z)
}
