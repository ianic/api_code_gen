package service

import "github.com/ianic/api_code_gen/service/dto"

//go:generate go install github.com/ianic/api_code_gen/service/dto
//go:generate go run gen.go

type Service struct{}

func (s *Service) Add(req dto.AddReq) (dto.AddRsp, error) {
	z := req.X + req.Y
	if z > 128 {
		return dto.AddRsp{}, dto.ErrOverflow
	}
	return dto.AddRsp{Z: req.X + req.Y}, nil
}

func (s *Service) Multiply(req dto.MultiplyReq) (dto.MultiplyRsp, error) {
	z := req.X * req.Y
	if z > 128 {
		return dto.MultiplyRsp{}, dto.ErrOverflow
	}
	return dto.MultiplyRsp{Z: req.X + req.Y}, nil
}
