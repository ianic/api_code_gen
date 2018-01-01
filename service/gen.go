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
		TransportRetries: 128,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// TODO
//+ * generirati nsq paket
//+   * parametar za generaciju je naziv topica
//+ * izdvojiti serijalizaciju iz transporta
//+   transport je []bytes unutra, []bytes van
//+   on na to dodaje svoj header da zna sto s time
//+   onda niti ne trebaju one funkcije za registraciju tipova
//+   sve se moze napraviti kada dobijemo []bytes i method name
//+ * izbaci null u drugom dijelu poruke
//+   to je zasto sto se nil serijalizira
//+ * parametar za lokaciju api paketa, ili konvencija oko toga
//+ * razmisliti o redizajnu nsq-a
//+ * omoguciti da metode imaju iste parametre

//+ * ErrTransport ?
//   kako razlikovati poruke na koje treba ovo ili ono
// * retry ako druga strana nije ziva
//   kako razlikovati transport errors od application
//   fatal errors kada negdje pukne serijalizacija
//   na trasport errors se ispalti ponavljati, na application nema ponavljanja (opet cu dobiti isto)
//   prema tome ako je trasport error onda je ok da rpc layer sam ponovi
// Svodim na slijedece:
//   retry je zadaca rpc liba, moze mu se definirati koliko puta da ponavljanja
//   sve transport error-e svodim na jedan ErrTransport
//   aplikacija uvijek moze odluciti da na taj error napravi neku odluku, zna da je problem u komunikaciji

// omoguciti da metode imaju builtin tipove za ulazne parametre
//   vise njih koje onda zapakujem u jedan struc i posaljem
// dignuti error ako config nije dobro postavljen, ako nesto fali
// za ovo vise ne mislim da je sjajna ideja, ne mogu ocitati nazive atibuta
//   to je zapravo razlog da realizacija ne ide

//+ dto paket ukljuciti u api
//+   ne mora vise ici kao parametar

// commitaj promjene koje si napravio u svckit

// * kako se prosljedjuju informacije u generator
//+ * bolji naming Type bi mogao biti ServiceType
//+   sto je client, server, service
//+ * mogu li izolirati serijalizaciju da je ima samo u api
//+   da ovi koji su van api zovu Marshall i Unmarshall
//+ * typ u service_gen zovi methodName
//+   od toga bi mogao napraviti konstante u api
