// Code generated by go generate; DO NOT EDIT.
package service

import (
	"fmt"

	"github.com/ianic/api_code_gen/service/api"
)

func (s *Service) Serve(method string, buf []byte) ([]byte, error) {
	switch method {
	case api.MethodAdd:
		var req api.TwoReq
		if err := api.Unmarshal(buf, &req); err != nil {
			return nil, err
		}
		rsp, err := s.Add(req)
		if err != nil {
			return nil, err
		}
		return api.Marshal(rsp)
	case api.MethodCube:
		var req int
		if err := api.Unmarshal(buf, &req); err != nil {
			return nil, err
		}
		rsp, err := s.Cube(req)
		if err != nil {
			return nil, err
		}
		return api.Marshal(rsp)
	case api.MethodMultiply:
		var req api.TwoReq
		if err := api.Unmarshal(buf, &req); err != nil {
			return nil, err
		}
		rsp, err := s.Multiply(req)
		if err != nil {
			return nil, err
		}
		return api.Marshal(rsp)
	case api.MethodMultiply2:
		var req api.TwoReq
		if err := api.Unmarshal(buf, &req); err != nil {
			return nil, err
		}
		rsp, err := s.Multiply2(req)
		if err != nil {
			return nil, err
		}
		return api.Marshal(rsp)
	default:
		return nil, fmt.Errorf("unknown method %s", method)
	}
}
