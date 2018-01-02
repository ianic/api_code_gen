package gen

import "text/template"

var nsqTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package nsq

import (
  "{{.ApiPkgPath}}"
	"github.com/minus5/svckit/nsq"
)

var (
	topic = "{{.NsqTopic}}"
)

func NewClient() *api.Client {
	return api.NewClient(nsq.NewRpcTransport(topic))
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
`))
