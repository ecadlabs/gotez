//go:build ignore

package main

import (
	"log"

	"github.com/ecadlabs/gotez/v2/clientv2/internal/generate"
)

var data = generate.TemplateData{
	Package: "monitor",
	Types: []*generate.TypeDef{
		{
			Stream:       true,
			RequestType:  "HeadsRequest",
			Path:         "/monitor/heads/{{.Chain}}",
			QueryParams:  map[string]string{"protocol": "Protocol", "next_protocol": "NextProtocol"},
			Func:         "Heads",
			ResponseType: "Head",
		},
	},
}

func main() {
	if err := generate.Execute(&data); err != nil {
		log.Fatal(err)
	}
}
