package nsq

import (
	"time"

	"github.com/ianic/api_code_gen/service/api"
	"github.com/ianic/api_code_gen/service/dto"
	"github.com/minus5/svckit/nsq"
)

var topic = "nsq_rr.req"

func NewClient() *api.Client {
	rr := nsq.NewRrClient(topic,
		api.NameFor,
		nil,
		5*time.Second,
		&nsq.ErrorsMapping{
			Parser:     api.ParseError,
			ErrStopped: dto.ErrTransport,
			ErrTimeout: dto.ErrTransport,
			ErrFatal:   dto.ErrTransport,
		})
	c := api.NewClient(rr)
	return c
}

type Closer interface {
	Close()
}

type Server interface {
	Serve(req interface{}) (interface{}, error)
}

func NewServer(srv Server) Closer {
	rr := nsq.NewRrServer(topic, srv, api.TypeFor, nil)
	return rr
}
