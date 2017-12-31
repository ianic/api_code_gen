// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates api for service.Service.
// It can be invoked by running:
// go generate
package main

import (
	"log"

	"github.com/ianic/api_code_gen/service"
	"github.com/ianic/api_code_gen/service/gen"
)

func main() {
	err := gen.Generate(
		&service.Service{},
		"github.com/ianic/api_code_gen/service/dto",
	)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO
// * generirati nsq paket
//   * parametar za generaciju je naziv topica
// * izdvojiti serijalizaciju iz transporta
//   transport je []bytes unutra, []bytes van
//   on na to dodaje svoj header da zna sto s time
//   onda niti ne trebaju one funkcije za registraciju tipova
//   sve se moze napraviti kada dobijemo []bytes i method name
// * izbaci null u drugom dijelu poruke
//   to je zasto sto se nil serijalizira
// * parametar za lokaciju api paketa
// * razmisliti o redizajnu nsq-a
// * ErrTransport ?
// * omoguciti da metode imaju iste parametre
