package core

import tz "github.com/ecadlabs/gotez"

type DelegatesList struct {
	Delegates []tz.PublicKeyHash `tz:"dyn"`
}
