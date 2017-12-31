// Code generated by go generate; DO NOT EDIT.
package service

import (
	"encoding/json"
	"fmt"
	"github.com/ianic/api_code_gen/service/dto"
)

func (s *Service) Serve(typ string, buf []byte) ([]byte, error) {
	switch typ {
	case "Add":
		var req dto.TwoReq
		if err := json.Unmarshal(buf, &req); err != nil {
			return nil, err
		}
		rsp, err := s.Add(req)
		if err != nil {
			return nil, err
		}
		return json.Marshal(rsp)
	case "Multiply":
		var req dto.TwoReq
		if err := json.Unmarshal(buf, &req); err != nil {
			return nil, err
		}
		rsp, err := s.Multiply(req)
		if err != nil {
			return nil, err
		}
		return json.Marshal(rsp)
	default:
		return nil, fmt.Errorf("unknown type %s", typ)
	}
}
