//go:build ignore

package main

import (
	"log"

	"github.com/ecadlabs/gotez/v2/clientv2/internal/generate"
)

var data = generate.TemplateData{
	Package: "mempool",
	Import: []string{
		"github.com/ecadlabs/gotez/v2",
	},
	Types: []*generate.TypeDef{
		{
			RequestType:  "gotez.ChainID",
			Method:       "GET",
			Path:         "/chains/{{.}}/mempool/pending_operations",
			Func:         "PendingOperations",
			ResponseType: "PendingOperationsResponse",
			AllocMode:    generate.ModeAllocate,
			PtrResult:    true,
		},
		{
			Stream:       true,
			RequestType:  "gotez.ChainID",
			Path:         "/chains/{{.}}/mempool/monitor_operations",
			Func:         "Monitor",
			ResponseType: "MonitorResponse",
		},
	},
}

func main() {
	if err := generate.Execute(&data); err != nil {
		log.Fatal(err)
	}
}
