//go:build ignore

package main

import (
	"log"

	"github.com/ecadlabs/gotez/v2/clientv2/internal/generate"
)

var data = generate.TemplateData{
	Import: []string{
		"github.com/ecadlabs/gotez/v2",
	},
	Package: "utils",
	Types: []*generate.TypeDef{
		{
			RequestType:  "InjectOperationRequest",
			Method:       "POST",
			Path:         "/injection/operation",
			QueryParams:  map[string]string{"chain": "Chain", "async": "Async"},
			Func:         "InjectOperation",
			ResponseType: "gotez.OperationHash",
			AllocMode:    generate.ModeAllocate,
			PtrResult:    true,
		},
		{
			RequestType:  "gotez.ChainID",
			Method:       "GET",
			Path:         "/chains/{{.}}/is_bootstrapped",
			Func:         "IsBootstrapped",
			ResponseType: "BootstrappedResponse",
			AllocMode:    generate.ModeAllocate,
			PtrResult:    true,
		},
	},
}

func main() {
	if err := generate.Execute(&data); err != nil {
		log.Fatal(err)
	}
}
