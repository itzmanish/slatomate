package types

type ArrayString []string

func (arr ArrayString) Has(value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
