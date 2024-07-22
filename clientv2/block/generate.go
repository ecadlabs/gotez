//go:build ignore

package main

import (
	"fmt"
	"log"

	"github.com/ecadlabs/gotez/v2/clientv2/internal/generate"
)

var data = generate.TemplateData{
	Import: []string{
		"github.com/ecadlabs/gotez/v2",
		"github.com/ecadlabs/gotez/v2/protocol",
		"github.com/ecadlabs/gotez/v2/protocol/core",
		"github.com/ecadlabs/gotez/v2/protocol/latest",
	},
	Package: "block",
	Types: []*generate.TypeDef{
		{
			RequestType:  "SimpleRequest",
			Method:       "GET",
			Path:         "/chains/{{.Chain}}/blocks/{{.Block}}/hash",
			Func:         "Hash",
			ResponseType: "gotez.BlockHash",
			AllocMode:    generate.ModeAllocate,
			PtrResult:    true,
		},
		{
			RequestType:  "SimpleRequest",
			Method:       "GET",
			Path:         "/chains/{{.Chain}}/blocks/{{.Block}}/protocols",
			Func:         "Protocols",
			ResponseType: "core.BlockProtocols",
			AllocMode:    generate.ModeAllocate,
			PtrResult:    true,
		},
		{
			RequestType:  "SimpleRequest",
			Method:       "GET",
			Path:         "/chains/{{.Chain}}/blocks/{{.Block}}/header/shell",
			Func:         "ShellHeader",
			ResponseType: "core.ShellHeader",
			AllocMode:    generate.ModeAllocate,
			PtrResult:    true,
		},
		{
			RequestType:   "BlockRequest",
			Method:        "GET",
			Path:          "/chains/{{.Chain}}/blocks/{{.Block}}/header",
			QueryParams:   map[string]string{"metadata": "Metadata"},
			Func:          "Header",
			ResponseType:  "protocol.BlockHeaderInfo",
			AllocMode:     generate.ModeConstruct,
			ConstructExpr: func(proto string) string { return fmt.Sprintf("protocol.NewBlockHeaderInfo(%s)", proto) },
		},
		{
			RequestType:   "BlockRequest",
			Method:        "GET",
			Path:          "/chains/{{.Chain}}/blocks/{{.Block}}",
			QueryParams:   map[string]string{"metadata": "Metadata"},
			Func:          "Block",
			ResponseType:  "protocol.BlockInfo",
			AllocMode:     generate.ModeConstruct,
			ConstructExpr: func(proto string) string { return fmt.Sprintf("protocol.NewBlockInfo(%s)", proto) },
		},
		{
			RequestType:  "ContractRequest",
			Method:       "GET",
			Path:         "/chains/{{.Chain}}/blocks/{{.Block}}/context/contracts/{{.ID}}/balance",
			Func:         "ContractBalance",
			ResponseType: "gotez.BigUint",
			AllocMode:    generate.ModeVar,
		},
		{
			RequestType:  "ContractRequest",
			Method:       "GET",
			Path:         "/chains/{{.Chain}}/blocks/{{.Block}}/context/contracts/{{.ID}}/balance_and_frozen_bonds",
			Func:         "ContractBalanceAndFrozenBonds",
			ResponseType: "gotez.BigUint",
			AllocMode:    generate.ModeVar,
		},
		{
			RequestType:  "ContractRequest",
			Method:       "GET",
			Path:         "/chains/{{.Chain}}/blocks/{{.Block}}/context/contracts/{{.ID}}/counter",
			Func:         "ContractCounter",
			ResponseType: "gotez.BigUint",
			AllocMode:    generate.ModeVar,
		},
		{
			RequestType:  "RunOperationRequest",
			Method:       "POST",
			Path:         "/chains/{{.Chain}}/blocks/{{.Block}}/helpers/scripts/run_operation",
			Func:         "RunOperation",
			ResponseType: "latest.OperationWithOptionalMetadata",
			AllocMode:    generate.ModeAllocate,
			PtrResult:    true,
		},
		{
			RequestType:   "ContextRequest",
			Method:        "GET",
			Path:          "/chains/{{.Chain}}/blocks/{{.Block}}/context/constants",
			Func:          "Constants",
			ResponseType:  "core.Constants",
			AllocMode:     generate.ModeConstruct,
			ConstructExpr: func(proto string) string { return fmt.Sprintf("protocol.NewConstants(%s)", proto) },
		},
	},
}

func main() {
	if err := generate.Execute(&data); err != nil {
		log.Fatal(err)
	}
}
