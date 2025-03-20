package utils_test

import (
	"os"
	"riscv-lsp/utils"
	"testing"
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
