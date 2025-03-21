package utils_test

import (
	"os"
	"riscv-lsp/types"
	"riscv-lsp/utils"
	"testing"
	"fmt"
)

func TestUri2Path(t *testing.T){
	teststr := "file:///a.S"
	test := utils.Uri2Path(teststr)
	expected := "/a.S"
	if test != expected{
		t.Fatalf("Expected: %s, Got %s\n",expected, test)
	}
}

func TestFileExists(t *testing.T){
	path := "a.S"
	defer os.RemoveAll(path)
	os.WriteFile(path, []byte(""), 0700)
	if !utils.FileExists(path){
		t.Fatalf("Path %s exists, but FileExists doesn't recognize that", path)
	}
	path = "randomS"
	if utils.FileExists(path){
		t.Fatalf("Path %s doesn't exist, but FileExists doesn't recognize that", path)
	}	
}

func TestWordAtPos(t *testing.T){
	teststr := "worda,wordb wordc"
	word := utils.WordAtPos(teststr, types.Position{Line:0,Character:7})
	fmt.Print(word)
}
