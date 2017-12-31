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

// primjer da dvije metode mogu imati iste atribute
func (s *Service) Multiply(req dto.TwoReq) (*dto.OneRsp, error) {
	z := req.X * req.Y
	if z > 128 {
		return nil, dto.ErrOverflow
	}
	return &dto.OneRsp{Z: z}, nil
}

// primjer kako rezultat moze biti build-in type
func (s *Service) Multiply2(req dto.TwoReq) (*int, error) {
	z := req.X * req.Y
	if z > 128 {
		return nil, dto.ErrOverflow
	}
	return &z, nil
}

// build-in tipovi unutra i van
func (s *Service) Cube(x int) (*int, error) {
	z := x * x
	if z > 256 {
		return nil, dto.ErrOverflow
	}
	return &z, nil
}

func (s *Service) Close() {}
