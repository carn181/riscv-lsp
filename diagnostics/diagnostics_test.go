package diagnostics_test

import (
	"reflect"
	"riscv-lsp/diagnostics"
	"testing"
)

func TestParseAssemblerError(t *testing.T){
	teststr := "files/a.s:1: Error: symbol `.text' is already defined"
	test :=  diagnostics.ParseAssemblerError(teststr)
	expected := diagnostics.AssemblerError{"files/a.s",1,"symbol `.text' is already defined"}
	if !reflect.DeepEqual(test, expected){
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}
}
