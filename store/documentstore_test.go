package store_test

import (
	"reflect"
	"riscv-lsp/store"
	"riscv-lsp/types"
	"testing"

	"github.com/stretchr/testify/assert"
	//	"strconv"
)

type EncodingExample struct {
	Testing bool
}

func TestFindNewLineIndices(t *testing.T) {
	test := store.FindNewLineIndices("a\n\n")
	expected := []uint{0, 2, 3}
	if !reflect.DeepEqual(test, expected) {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}
	test = store.FindNewLineIndices("")
	expected = []uint{0}
	if !reflect.DeepEqual(test, expected) {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}
	test = store.FindNewLineIndices("test")
	expected = []uint{0}
	if !reflect.DeepEqual(test, expected) {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}
}

func TestPositionToOffset(t *testing.T) {
	var test, expected uint
	var err error

	// Blank String Test
	test, _ = store.PositionToOffset("", types.Position{Line: 0, Character: 0})
	expected = uint(0)
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}

	// Blank String, but the next line doesn't error out
	test, _ = store.PositionToOffset("", types.Position{Line: 1, Character: 0})
	expected = uint(0)
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}

	// Blank String, next to next line errors out
	test, err = store.PositionToOffset("", types.Position{Line: 2, Character: 0})
	assert.EqualError(t, err, "Line Number 2 out of range 0-1")

	// Column Beyond End of File
	test, err = store.PositionToOffset("", types.Position{Line: 0, Character: 4})
	assert.EqualError(t, err, "Column is beyond end of file")

	// Column Beyond End of File
	test, err = store.PositionToOffset("", types.Position{Line: 1, Character: 1})
	assert.EqualError(t, err, "Column is beyond end of file")

	// Column Beyond End of Line
	test, err = store.PositionToOffset("\n\n", types.Position{Line: 1, Character: 1})
	assert.EqualError(t, err, "Column is beyond end of line")

	// Normal
	test, err = store.PositionToOffset("a", types.Position{Line: 0, Character: 1})
	if err != nil {
		t.Fatal(err)
	}
	expected = uint(1)
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}

	test, err = store.PositionToOffset("abc\n", types.Position{Line: 1, Character: 0})
	if err != nil {
		t.Fatal(err)
	}
	expected = uint(4)
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}
	
	test, err = store.PositionToOffset("0123\n0", types.Position{Line: 1, Character: 0})
	if err != nil {
		t.Fatal(err)
	}
	expected = uint(5)
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}

	test, err = store.PositionToOffset("0123\n\n0", types.Position{Line: 1, Character: 0})
	expected = uint(5)
	if err != nil {
		t.Fatal(err)
	}
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}

	test, err = store.PositionToOffset("01234\n01234\n\n01234", types.Position{Line: 4, Character: 0})
	expected = uint(18)
	if err != nil {
		t.Fatal(err)
	}
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}

	test, err = store.PositionToOffset("01234\n01234\n\n01234", types.Position{Line: 3, Character: 2})
	expected = uint(15)
	if err != nil {
		t.Fatal(err)
	}
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}
}

func TestReplaceStrSlice(t *testing.T) {
	test := store.ReplaceStrSlice("0123456", 7, 7, "7")
	expected := "01234567"
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}

	// test = store.ReplaceStrSlice("test\n", 3, 5, "\n")
	// expected = "tes\n\n"
	// if test != expected{
	// 	t.Fatalf("Expected: %s, Got: %s",strconv.Quote(expected), strconv.Quote(test)		)
	// }

	test = store.ReplaceStrSlice("test\n", 5, 5, "\n")
	expected = "test\n\n"
	if test != expected {
		t.Fatalf("Expected: %s, Got: %s", expected, test)
	}
}

func TestApplyChange(t *testing.T) {
	test := store.ApplyChange("012345",
		types.Range{Start: types.Position{Line: 0, Character: 0},
			End: types.Position{Line: 0, Character: 3}},
		"test")
	expected := "test345"
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}
	test = store.ApplyChange("012345\n01234",
		types.Range{Start: types.Position{Line: 0, Character: 0},
			End: types.Position{Line: 1, Character: 3}},
		"test")
	expected = "test34"
	if test != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, test)
	}	
}
