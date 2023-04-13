package encoding

import (
	"fmt"
	"reflect"
	"strings"
)

type ErrorPath []PathSeg

func (e ErrorPath) String() string {
	var out strings.Builder
	for _, x := range e {
		out.WriteString(x.String())
	}
	return out.String()
}

type PathSeg interface {
	SegSelector()
	fmt.Stringer
}

type IndexSelector int

func (IndexSelector) SegSelector() {}
func (i IndexSelector) String() string {
	return fmt.Sprintf("[%d]", i)
}

type FieldSelector reflect.StructField

func (FieldSelector) SegSelector() {}
func (f *FieldSelector) String() string {
	return fmt.Sprintf(".%s", f.Name)
}

type TypeSelector struct {
	reflect.Type
}

func (TypeSelector) SegSelector() {}
func (t TypeSelector) String() string {
	return fmt.Sprintf("<%v>", t.Type)
}

type Error struct {
	Path ErrorPath
	Err  error
}

func (e *Error) Error() string {
	return fmt.Sprintf("gotez: %v", e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func wrapError(err error, path ErrorPath) error {
	if err != nil {
		return &Error{Path: path, Err: err}
	}
	return nil
}
