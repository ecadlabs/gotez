package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ecadlabs/gotez/v2/b58/base58"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		_, data, err := base58.DecodeTZ([]byte(os.Args[i]))
		if err != nil {
			log.Fatal(err)
		}
		hex := make([]string, len(data))
		for j, x := range data {
			hex[j] = fmt.Sprintf("0x%02x", x)
		}
		fmt.Println(strings.Join(hex, ","))
	}
}
