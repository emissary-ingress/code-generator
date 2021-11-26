package testtypes

type MyEnum int

type StructA struct {
	Foo int `json:"a"`
	Bar string
}

type StructB struct {
	Foo int `json:"b"`
	Bar string
}

type BoolOrString struct {
	Bool   *bool
	String *string
}
