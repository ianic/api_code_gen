package dto

import "errors"

type AddReq struct {
	X int
	Y int
}

type AddRsp struct {
	Z int
}

var (
	ErrOverflow = errors.New("overflow")
)
