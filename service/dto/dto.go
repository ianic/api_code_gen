package dto

import "errors"

var (
	ErrOverflow  = errors.New("overflow")
	ErrTransport = errors.New("trasnsport error")
)

type TwoReq struct {
	X int
	Y int
}

type OneRsp struct {
	Z int
}
