package util

type Util struct {
}

func NewUtil() *Util {
	u := Util{}
	return &u
}

func (u *Util) Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
