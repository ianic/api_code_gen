// Code generated by go generate; DO NOT EDIT.
package nsq

import (
	"time"

	"github.com/ianic/api_code_gen/service/api"
	"github.com/minus5/svckit/nsq"
)

var (
	topic = "nsq_rr.req"
	ttl   = 2 * time.Second
)

func NewClient() *api.Client {
	return api.NewClient(
		nsq.NewRpcTransport(
			topic,
			ttl,
			&nsq.ErrorsMapping{
				ErrStopped: api.ErrTransport,
				ErrTimeout: api.ErrTransportTimeout,
				ErrFatal:   api.ErrTransport,
			},
		))
}

type Closer interface {
	Close()
}

type server interface {
	Serve(typ string, req []byte) ([]byte, error)
}

func NewServer(srv server) Closer {
	return nsq.RpcServe(topic, srv.Serve)
}
