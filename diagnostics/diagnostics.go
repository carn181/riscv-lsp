package diagnostics

import (
	"path/filepath"
	"reflect"
	"riscv-lsp/store"
	"riscv-lsp/types"
	"riscv-lsp/utils"
	"slices"
	"strconv"
	"strings"
)

// TODO
// Parse Linker Errors as well

type AssemblerError struct {
	File    string
	Line    uint
	Message string
}

var Diagnostics map[string]*[]types.Diagnostic

func AssemblerError2Diagnostic(path string, err AssemblerError)types.Diagnostic{
	fileContents := store.FileContents(filepath.Join(path,err.File))

	lines := utils.FindNewLineIndices(string(fileContents))
	var f, n uint
	f = lines[err.Line-1]
	if err.Line >= uint(len(lines)){
		n = uint(len(fileContents))
	} else {
		n = lines[err.Line]
	}
	charsInLine := n-f
	return types.Diagnostic{
		Range: types.Range{
			Start: types.Position{
				Line: err.Line-1,
				Character: 0,
			},
			End: types.Position{
				Line: err.Line-1,
				Character: charsInLine,
			},
		},
		Severity: 1,
		Source: "gas",
		Message: err.Message,		
	}
}

func ParseAssemblerError(line string) AssemblerError{
	// Every error is format:
	// FILE:LINE_NO: Error: ERROR_MSG
	
	parts := strings.Split(line, ": Error: ")
	if(len(parts)<2){
		return AssemblerError{}
	}
	message := parts[1]
	left := parts[0]	
	
	// left = "FILE:LINE_NO"
	leftparts := strings.Split(left, ":")
	file := leftparts[0]
	line_no, _ := strconv.Atoi(leftparts[1])

	return AssemblerError{file,uint(line_no),message}
}

func ParseAssemblerErrors(messages string) []AssemblerError{
	lines := strings.Split(messages, "\n")
	lines = slices.DeleteFunc(lines, func(e string) bool{return e == ""})
	if len(lines) < 2{
		return nil
	}
	lines = lines[1:] // First Line is typically just FILE: Assembler messages:
	
	var errors []AssemblerError
	for _, line := range lines{
		error := ParseAssemblerError(line)
		if !reflect.DeepEqual(error, AssemblerError{}){
			errors = append(errors, error)
		}
	}

	return errors
}
