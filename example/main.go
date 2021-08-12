package example

//go:generate go run ../main.go
type A struct {
	ID  string
	Foo int
	Bar string
}

type B struct {
	Foo int
	Bar string
	AID string `goentity:"A.ID"`
}

/* type AB struct {
	A
	B
}*/
