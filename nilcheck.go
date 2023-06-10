package nilcheck

import "reflect"

func IsNil(val any) bool {
	if val == nil {
		return true
	}

	// Check if val is a type that IsNil can be called on.
	switch reflect.TypeOf(val).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(val).IsNil()
	}

	return false
}
