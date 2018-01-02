package service

import (
	"context"

	"github.com/ianic/api_code_gen/service/api"
)

//go:generate find . -type f -name "*_gen.go" -exec rm -f {} ;
//go:generate go install github.com/ianic/api_code_gen/service/api
//go:generate go run gen.go

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Add(ctx context.Context, req api.TwoReq) (*api.OneRsp, error) {
	z := req.X + req.Y
	if z > 128 {
		return nil, api.ErrOverflow
	}
	return &api.OneRsp{Z: z}, nil
}

// primjer da dvije metode mogu imati iste atribute
func (s *Service) Multiply(ctx context.Context, req api.TwoReq) (*api.OneRsp, error) {
	z := req.X * req.Y
	if z > 128 {
		return nil, api.ErrOverflow
	}
	return &api.OneRsp{Z: z}, nil
}

// primjer kako rezultat moze biti build-in type
func (s *Service) Multiply2(ctx context.Context, req api.TwoReq) (*int, error) {
	z := req.X * req.Y
	if z > 128 {
		return nil, api.ErrOverflow
	}
	return &z, nil
}

// build-in tipovi unutra i van
func (s *Service) Cube(ctx context.Context, x int) (*int, error) {
	z := x * x
	if z > 256 {
		return nil, api.ErrOverflow
	}
	return &z, nil
}

func (s *Service) Close() {}
