package entities

type Page struct {
	Value int
}

func NewPage(value int) Page {
	return Page{Value: value}
}
