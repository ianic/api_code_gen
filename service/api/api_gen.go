// Code generated by go generate; DO NOT EDIT.
package api

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/ianic/api_code_gen/service/dto"
	"github.com/minus5/svckit/types/registry"
)

type transport interface {
	Call(string, []byte) ([]byte, error)
	Close()
}

type Client struct {
	t transport
}

func NewClient(r transport) *Client {
	return &Client{t: r}
}

func (c *Client) Add(req dto.AddReq) (*dto.AddRsp, error) {
	rsp := &dto.AddRsp{}
	if err := c.call("Add", req, rsp); err != nil {
		return nil, parseError(err)
	}
	return rsp, nil
}

func (c *Client) Multiply(req dto.MultiplyReq) (*dto.MultiplyRsp, error) {
	rsp := &dto.MultiplyRsp{}
	if err := c.call("Multiply", req, rsp); err != nil {
		return nil, parseError(err)
	}
	return rsp, nil
}

func (c *Client) call(typ string, req, rsp interface{}) error {
	reqBuf, err := json.Marshal(req)
	if err != nil {
		return err
	}
	rspBuf, err := c.t.Call(typ, reqBuf)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(rspBuf, rsp); err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() {
	c.t.Close()
}

var (
	typeRegistry = registry.New()
)

func init() {
	typeRegistry.Add([]interface{}{
		dto.AddReq{},
		dto.MultiplyReq{},
	})
}

func NameFor(i interface{}) string {
	return typeRegistry.NameFor(i)
}

func TypeFor(typ string) reflect.Type {
	return typeRegistry.TypeFor(typ)
}

func ParseError(text string) error {
	if text == "" {
		return nil
	}
	switch text {
	case dto.ErrOverflow.Error():
		return dto.ErrOverflow
	case dto.ErrTransport.Error():
		return dto.ErrTransport
	}
	return errors.New(text)
}

func parseError(err error) error {
	if err == nil {
		return nil
	}
	switch err.Error() {
	case dto.ErrOverflow.Error():
		return dto.ErrOverflow
	case dto.ErrTransport.Error():
		return dto.ErrTransport
	}
	return err
}
