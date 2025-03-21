package symbols_test

import (
	"riscv-lsp/symbols"
	"testing"
)

func TestInit(t *testing.T){
	err := symbols.Init()
	if err != nil{
		t.Fatal(err)
	}
}
