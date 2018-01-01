package gen

import (
	"fmt"
	"go/importer"
	"html/template"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strings"
)

type config struct {
	ServiceType reflect.Type
	NsqTopic    string
	NsqTtl      int
	apiPkgDir   string
	nsqPkgDir   string
	apiPkgPath  string
}

// Config returns defalut config
func Config() config {
	return config{
		NsqTopic:  "",
		NsqTtl:    60,
		apiPkgDir: "api",
		nsqPkgDir: "api/nsq",
	}
}

// data collects atributes for template execution
type data struct {
	Package    string
	Struct     string
	Methods    []method
	Errors     []string
	NsqTopic   string
	NsqTtl     int
	ApiPkgPath string
}

type method struct {
	Name      string
	In        string
	InWithPkg string
	Out       string
}

// Generator code generator
type Generator struct {
	c    config
	data data
}

// Generate run generator with  config
func Generate(c config) error {
	c.apiPkgPath = c.ServiceType.PkgPath() + "/" + c.apiPkgDir

	g := Generator{c: c}

	ms, err := g.findMethods()
	if err != nil {
		return err
	}
	es, err := g.findErrors()
	if err != nil {
		return err
	}

	pkg, stc := g.packageStruct()
	g.data = data{
		Package:    pkg,
		Struct:     stc,
		Methods:    ms,
		Errors:     es,
		NsqTopic:   c.NsqTopic,
		NsqTtl:     c.NsqTtl,
		ApiPkgPath: c.apiPkgPath,
	}

	if err := g.execTemplate(apiTemplate, c.apiPkgDir+"/api_gen.go"); err != nil {
		return err
	}

	if err := g.execTemplate(nsqTemplate, c.nsqPkgDir+"/nsq_gen.go"); err != nil {
		return err
	}

	fn := fmt.Sprintf("%s_gen.go", strings.ToLower(stc))
	if err := g.execTemplate(serviceTemplate, fn); err != nil {
		return err
	}

	return nil
}

func (g *Generator) execTemplate(t *template.Template, fn string) error {
	os.MkdirAll(path.Dir(fn), os.ModePerm)
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	t.Execute(f, g.data)
	f.Close()
	err = exec.Command("go", "fmt", fn).Run()
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) findMethods() ([]method, error) {
	v := reflect.New(g.c.ServiceType)
	var ms []method
	for i := 0; i < v.NumMethod(); i++ {
		tm := v.Type().Method(i)

		if tm.Name == "Serve" {
			fmt.Printf("skipping generated method %s\n", tm.Name)
			continue
		}
		m := v.Method(i)

		if m.Type().NumIn() != 1 &&
			m.Type().NumOut() != 2 {
			fmt.Printf("skipping method %s, unsupported signature\n", tm.Name)
			continue
		}
		if m.Type().Out(1).String() != "error" {
			fmt.Printf("skipping method %s, unsupported signature\n", tm.Name)
			continue
		}

		in := m.Type().In(0).String()
		out := m.Type().Out(0).String()

		if isPointer(in) {
			fmt.Printf("skipping method %s, input arg must be passed by value\n", tm.Name)
			continue
		}
		if !isPointer(out) {
			fmt.Printf("skipping method %s, output arg must be passed by reference\n", tm.Name)
			continue
		}

		ms = append(ms, method{
			Name:      tm.Name,
			InWithPkg: m.Type().In(0).String(),
			In:        removePackagePrefix(m.Type().In(0).String()),
			Out:       removePackagePrefix(removePointerPrefix(out)),
		})
	}
	return ms, nil
}

func isPointer(typ string) bool {
	return strings.HasPrefix(typ, "*")
}

func removePointerPrefix(typ string) string {
	if isPointer(typ) {
		return typ[1:]
	}
	return typ
}

func removePackagePrefix(typ string) string {
	p := strings.Split(typ, ".")
	return p[len(p)-1]
}

func (g *Generator) findErrors() ([]string, error) {
	var es []string
	pkg, err := importer.Default().Import(g.c.apiPkgPath)
	if err != nil {
		return nil, err
	}
	for _, n := range pkg.Scope().Names() {
		if strings.HasPrefix(n, "Err") {
			//es = append(es, fmt.Sprintf("%s.%s", pkg.Name(), n))
			es = append(es, n)
		}
	}
	return es, nil
}

func (g *Generator) packageStruct() (string, string) {
	typ := g.c.ServiceType.String()
	typ = removePointerPrefix(typ)
	p := strings.Split(typ, ".")
	return p[0], p[1]
}
