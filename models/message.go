package models

type Person[T any] struct {
	From string
	To   string
	Msg  T
}

var (
	person Person[int]
)

func Test[T int](a T) (b T) {
	p := &person
	p.Msg = 12
	return
}
