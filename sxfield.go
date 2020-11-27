package sxfield

type SxField interface {
	Changed() bool
	SQL() string
}

func AnyChanged(sxFields []SxField) bool {
	for _, v := range sxFields {
		if v.Changed() {
			return true
		}
	}
	return false
}