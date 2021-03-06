// Code generated by go generate; DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// method names constants
const (
	MethodAdd       = "Add"
	MethodCube      = "Cube"
	MethodMultiply  = "Multiply"
	MethodMultiply2 = "Multiply2"
	timeout         = 2 * time.Second
)

type transport interface {
	Call(ctx context.Context, method string, req []byte) ([]byte, string, error)
	Close()
}

type Client struct {
	t transport
}

func NewClient(t transport) *Client {
	return &Client{t: t}
}

func (c *Client) Add(ctx context.Context, req TwoReq) (*OneRsp, error) {
	rsp := new(OneRsp)
	if err := c.call(ctx, MethodAdd, req, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) Cube(ctx context.Context, req int) (*int, error) {
	rsp := new(int)
	if err := c.call(ctx, MethodCube, req, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) Multiply(ctx context.Context, req TwoReq) (*OneRsp, error) {
	rsp := new(OneRsp)
	if err := c.call(ctx, MethodMultiply, req, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) Multiply2(ctx context.Context, req TwoReq) (*int, error) {
	rsp := new(int)
	if err := c.call(ctx, MethodMultiply2, req, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) call(ctx context.Context, method string, req, rsp interface{}) error {
	reqBuf, err := Marshal(req)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	ctxT, _ := context.WithTimeout(ctx, timeout)
	rspBuf, appErr, err := c.t.Call(ctxT, method, reqBuf)
	if ctxT.Err() != nil {
		return ctxT.Err() // context.Canceled || context.DeadlineExceeded
	}
	if err != nil {
		return errors.Wrap(err, "transport failed")
	}
	if appErr != "" {
		return toAppError(appErr)
	}

	if err := Unmarshal(rspBuf, rsp); err != nil {
		return errors.Wrap(err, "unmarshal failed")
	}
	return nil
}

func (c *Client) Close() {
	c.t.Close()
}

func toAppError(txt string) error {
	if txt == "" {
		return nil
	}
	switch txt {
	case Overflow.Error():
		return Overflow
	}
	return fmt.Errorf(txt)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
