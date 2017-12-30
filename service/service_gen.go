// Code generated by go generate; DO NOT EDIT.
package service

import (
	"fmt"
	"github.com/ianic/api_code_gen/service/dto"
)

func (s *Service) Serve(i interface{}) (interface{}, error) {
	switch req := i.(type) {
	case dto.AddReq:
		return s.Add(req)
	case dto.MultiplyReq:
		return s.Multiply(req)
	default:
		return nil, fmt.Errorf("unknown type %T", req)
	}
}