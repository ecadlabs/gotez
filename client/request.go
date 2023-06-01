package client

//go:generate go run generate.go

type MetadataMode int

const (
	MetadataDefault MetadataMode = iota
	MetadataAlways
	MetadataNever
)

func (m MetadataMode) String() string {
	switch m {
	case MetadataAlways:
		return "always"
	case MetadataNever:
		return "never"
	default:
		return "default"
	}
}

type BlockRequest struct {
	Chain    string
	Block    string
	Metadata MetadataMode
}
