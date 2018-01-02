// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates api for service.Service.
// It can be invoked by running:
// go generate
package main

import (
	"log"
	"reflect"

	"github.com/ianic/api_code_gen/service"
	"github.com/ianic/api_code_gen/service/gen"
)

func main() {
	err := gen.Generate(gen.Config{
		ServiceType:      reflect.TypeOf(service.Service{}),
		NsqTopic:         "nsq_rr.req",
		TransportTimeout: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
}
