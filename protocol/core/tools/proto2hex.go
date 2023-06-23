package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ecadlabs/gotez/v2/b58"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		hash, err := b58.ParseProtocolHash([]byte(os.Args[i]))
		if err != nil {
			log.Fatal(err)
		}
		hex := make([]string, len(hash))
		for j, x := range hash {
			hex[j] = fmt.Sprintf("0x%02x", x)
		}
		fmt.Println(strings.Join(hex, ","))
	}
}
