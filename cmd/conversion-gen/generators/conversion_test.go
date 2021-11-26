package generators

import (
	"reflect"
	"testing"

	"k8s.io/gengo/parser"
	"k8s.io/gengo/types"

	"k8s.io/code-generator/cmd/conversion-gen/generators/internal/testtypes"
	"k8s.io/code-generator/cmd/conversion-gen/generators/internal/testtypes/example1"
	"k8s.io/code-generator/cmd/conversion-gen/generators/internal/testtypes/example2"
)

func reflectIsDirectlyAssignable(outType, inType reflect.Type) bool {
	// This might produce false-positives if they're both interfaces.
	return (inType.Kind() == outType.Kind()) && inType.AssignableTo(outType)
}

func reflectIsDirectlyConvertible(outType, inType reflect.Type) bool {
	return (inType.Kind() == outType.Kind()) && inType.ConvertibleTo(outType)
}

func TestDirect(t *testing.T) {
	builder := parser.New()
	if err := builder.AddDir("./internal/testtypes"); err != nil {
		t.Fatal(err)
	}
	if err := builder.AddDir("./internal/testtypes/example1"); err != nil {
		t.Fatal(err)
	}
	if err := builder.AddDir("./internal/testtypes/example2"); err != nil {
		t.Fatal(err)
	}
	testPackages, err := builder.FindTypes()
	if err != nil {
		t.Fatal(err)
	}

	typeMapReflect := map[string]reflect.Type{
		"bool":    reflect.TypeOf(false),
		"int":     reflect.TypeOf(0),
		"string":  reflect.TypeOf(""),
		"*bool":   reflect.PtrTo(reflect.TypeOf(false)),
		"*string": reflect.PtrTo(reflect.TypeOf("")),

		"./internal/testtypes.BoolOrString": reflect.TypeOf(testtypes.BoolOrString{}),
		"./internal/testtypes.MyEnum":       reflect.TypeOf(testtypes.MyEnum(0)),
		"./internal/testtypes.StructA":      reflect.TypeOf(testtypes.StructA{}),
		"./internal/testtypes.StructB":      reflect.TypeOf(testtypes.StructB{}),

		"./internal/testtypes/example1.MyString":   reflect.TypeOf(example1.MyString("")),
		"./internal/testtypes/example1.MyStruct":   reflect.TypeOf(example1.MyStruct{}),
		"*./internal/testtypes/example1.MyStruct":  reflect.PtrTo(reflect.TypeOf(example1.MyStruct{})),
		"[]./internal/testtypes/example1.MyStruct": reflect.SliceOf(reflect.TypeOf(example1.MyStruct{})),

		"./internal/testtypes/example2.MyString":   reflect.TypeOf(example2.MyString("")),
		"./internal/testtypes/example2.MyStruct":   reflect.TypeOf(example2.MyStruct{}),
		"*./internal/testtypes/example2.MyStruct":  reflect.PtrTo(reflect.TypeOf(example2.MyStruct{})),
		"[]./internal/testtypes/example2.MyStruct": reflect.SliceOf(reflect.TypeOf(example2.MyStruct{})),
	}

	typeMapGengo := map[string]*types.Type{}
	for _, pkgdata := range testPackages {
		for _, typedata := range pkgdata.Types {
			name := typedata.Name.String()
			if _, ok := typeMapReflect[name]; !ok {
				t.Errorf("this needs updated: no typeMapReflect entry for %q", name)
			}
			typeMapGengo[name] = typedata
		}
	}
	if t.Failed() {
		return
	}

	for typename1 := range typeMapReflect {
		for typename2 := range typeMapReflect {
			expAssign := reflectIsDirectlyAssignable(typeMapReflect[typename1], typeMapReflect[typename2])
			t.Logf("exp: isDirctlyAssignable(%s, %s) => %v", typename1, typename2, expAssign)
			actAssign := isDirectlyAssignable(typeMapGengo[typename1], typeMapGengo[typename2])
			if actAssign != expAssign {
				t.Errorf("expected isDirectlyAssignable(%s, %s) to return %v, but got %v",
					typename1, typename2, expAssign, actAssign)
			}

			expConvert := reflectIsDirectlyConvertible(typeMapReflect[typename1], typeMapReflect[typename2])
			t.Logf("exp: isDirctlyConvertible(%s, %s) => %v", typename1, typename2, expConvert)
			actConvert := isDirectlyConvertible(typeMapGengo[typename1], typeMapGengo[typename2])
			if actConvert != expConvert {
				t.Errorf("expected isDirectlyConvertible(%s, %s) to return %v, but got %v",
					typename1, typename2, expConvert, actConvert)
			}
		}
	}
}
