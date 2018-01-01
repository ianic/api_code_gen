package api

import "errors"

// application specific errors
var (
	ErrOverflow = errors.New("overflow")
)

type TwoReq struct {
	X int
	Y int
}

type OneRsp struct {
	Z int
}
