package client

//go:generate go run generate.go

type BlockRequest struct {
	Chain string
	Block string
}
