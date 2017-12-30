package dto

import "errors"

var (
	ErrOverflow  = errors.New("overflow")
	ErrTransport = errors.New("trasnsport error")
)

type AddReq struct {
	X int
	Y int
}

type AddRsp struct {
	Z int
}

type MultiplyReq struct {
	X int
	Y int
}

type MultiplyRsp struct {
	Z int
}
