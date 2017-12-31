package nsq

import (
	"time"

	"github.com/ianic/api_code_gen/service/api"
	"github.com/minus5/svckit/nsq"
)

var (
	topic = "nsq_rr.req"
	ttl   = 5 * time.Second
)

func NewClient() *api.Client {
	return api.NewClient(nsq.NewRpcTransport(topic, ttl))
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
