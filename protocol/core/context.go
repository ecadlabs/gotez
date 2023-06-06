package core

import tz "github.com/ecadlabs/gotez/v2"

type DelegatesList struct {
	Delegates []tz.PublicKeyHash `tz:"dyn"`
}
