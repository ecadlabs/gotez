# teztool

An operation injection helper

## Example

```go
package example

import (
    "context"
    "fmt"
    "math/big"

    "github.com/ecadlabs/gotez/v2"
    "github.com/ecadlabs/gotez/v2/b58"
    "github.com/ecadlabs/gotez/v2/client"
    "github.com/ecadlabs/gotez/v2/crypt"
    "github.com/ecadlabs/gotez/v2/protocol/core"
    "github.com/ecadlabs/gotez/v2/protocol/latest"
    "github.com/ecadlabs/gotez/v2/teztool"
)

type logger struct{}

func (logger) Printf(format string, a ...any) {
    fmt.Printf(format, a...)
    fmt.Printf("\n")
}

func TransferToWallet(url, chainID, address, privateKey string, amount *big.Int) (core.OperationsGroup, error) {
    chain, err := b58.ParseChainID([]byte(chainID))
    if err != nil {
        return nil, err
    }
    addr, err := b58.ParsePublicKeyHash([]byte(address))
    if err != nil {
        return nil, err
    }
    c := client.Client{
        URL: url,
    }
    pk, err := b58.ParsePrivateKey([]byte(privateKey))
    if err != nil {
        return nil, err
    }
    priv, err := crypt.NewPrivateKey(pk)
    if err != nil {
        return nil, err
    }

    // initialize tezool
    tool := teztool.New(&c, chain)
    tool.DebugLogger = logger{}

    // initialize signer
    signer := teztool.NewLocalSigner(priv)

    // make a transaction
    val, err := gotez.NewBigUint(amount)
    if err != nil {
        // amount is negative
        return nil, err
    }
    tx := latest.Transaction{
        ManagerOperation: latest.ManagerOperation{
            Source: priv.Public().Hash(),
        },
        Amount:      val,
        Destination: core.ImplicitContract{PublicKeyHash: addr},
    }

    return tool.FillSignAndInjectWait(context.Background(), signer, []latest.OperationContents{&tx}, client.MetadataAlways, teztool.FillAll)
}
```
