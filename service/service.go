package service

import "github.com/ianic/api_code_gen/service/dto"

//go:generate go install github.com/ianic/api_code_gen/service/dto
//go:generate go run gen.go

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Add(req dto.TwoReq) (*dto.OneRsp, error) {
	z := req.X + req.Y
	if z > 128 {
		return nil, dto.ErrOverflow
	}
	return &dto.OneRsp{Z: z}, nil
}

func (s *Service) Multiply(req dto.TwoReq) (*dto.OneRsp, error) {
	z := req.X * req.Y
	if z > 128 {
		return nil, dto.ErrOverflow
	}
	return &dto.OneRsp{Z: z}, nil
}
