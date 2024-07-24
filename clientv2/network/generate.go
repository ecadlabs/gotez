//go:build ignore

package main

import (
	"log"

	"github.com/ecadlabs/gotez/v2/clientv2/internal/generate"
)

var data = generate.TemplateData{
	Package: "network",
	Types: []*generate.TypeDef{
		{
			Method:       "GET",
			Path:         "/network/connections",
			Func:         "Connections",
			ResponseType: "ConnectionsResponse",
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
