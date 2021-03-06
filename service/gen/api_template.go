package gen

import "text/template"

var apiTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package api

import (
  "context"
	"encoding/json"
  "fmt"
  "time"

  "github.com/pkg/errors"
)

// method names constants
const (
{{- range .Methods }}
  Method{{.Name}} = "{{.Name}}"
{{- end}}
  timeout = {{.Timeout}} * time.Second
)

type transport interface {
	Call(ctx context.Context, method string, req []byte) ([]byte, string, error)
  Close()
}

type Client struct {
  t transport
}

func NewClient(t transport) *Client {
  return &Client{t: t}
}

{{- range .Methods }}

func (c *Client) {{.Name}}(ctx context.Context, req {{ .In }}) (*{{ .Out }}, error) {
  rsp := new({{ .Out }})
  if err := c.call(ctx, Method{{.Name}},req, rsp); err != nil {
    return nil, err
  }
  return rsp, nil
}
{{- end }}

func (c *Client) call(ctx context.Context, method string, req, rsp interface{}) error {
	reqBuf, err := Marshal(req)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	ctxT, _ := context.WithTimeout(ctx, timeout)
	rspBuf, appErr, err := c.t.Call(ctxT, method, reqBuf)
	if ctxT.Err() != nil {
		return ctxT.Err() // context.Canceled || context.DeadlineExceeded
	}
	if err != nil {
		return errors.Wrap(err, "transport failed")
	}
	if appErr != "" {
		return toAppError(appErr)
	}

	if err := Unmarshal(rspBuf, rsp); err != nil {
		return errors.Wrap(err, "unmarshal failed")
	}
	return nil
}

func (c *Client) Close() {
	c.t.Close()
}

func toAppError(txt string) error {
	if txt == "" {
		return nil
	}
	switch txt {
{{- range .Errors }}
  case {{ . }}.Error():
    return {{ . }}
{{- end }}
  }
	return fmt.Errorf(txt)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
`))
