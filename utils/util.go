package utils

import (
	"path/filepath"
	"os"
	"fmt"
	"os/exec"
	"riscv-lsp/types"
)

func Uri2Path(uri string) string{
	return uri[7:]    // file:// is 7 chars
}

func Path2Uri(path string) string{
	return "file://"+path    // file:// is 7 chars
}

func FilesInPath(path string)([]string, error){
	var files []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error{
			if err != nil {
				return err
			}
			if !info.IsDir(){
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("Couldn't Find Directory in Path")
	}
	return files, nil
}

func FileExists(path string) bool{
	_, err := os.Stat(path)

	return err == nil
}

func EnsurePathExists(path string){
	dir := filepath.Dir(path)
	if !FileExists(path) {
		os.MkdirAll(dir, 0700) // Create your file
		f, _ := os.Create(path)
		defer f.Close()
	}
}

func RunCommand(dir string, command string)([]byte, error){
	os.Chdir(dir)
	
	cmd := exec.Command("bash","-c",command)
	out, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			code := exitError.ExitCode()
			if code != 1 && code != 0{
				return nil, fmt.Errorf("Build failed with invalid exit code: %d\n",exitError.ExitCode())
			}
			out = exitError.Stderr
		}
	}
	return out, nil
}

func FindNewLineIndices(str string) []uint{
	indices := []uint{0}
	for i, r := range str {
		if r == '\n' {
			indices = append(indices, uint(i+1))
		}
	}
	return indices
}

func ReplaceStrSlice(str string, i uint, j uint, replace string) string{
	// i = index of where to start replacing from (Inclusive)
	// j = index of till where to replace till (Exclusive)
	if i > j {panic("Invalid Range: start > end")}
	if j > uint(len(str)){panic("Invalid Range: end > len(str)")}
	var start, remaining string
	
	remaining = str[j:]
	start=str[:i]
	
	return start+replace+remaining
}

func delimiters(char byte)bool{
	return char == ' ' || char == ',' || char == '\n'
}

func WordAtPos(str string, pos types.Position) string {
	// This is a very hacky function for getting Words at Positions
	// TODO
	// Use TreeSitter here to parse the line and then get the token that has the character on it
	if len(str) < 1{return ""}
	indices := FindNewLineIndices(str)
	offset := indices[pos.Line]+pos.Character
	
	start := offset
	end :=offset
	var lineStart, nextLine uint
	lineStart = indices[pos.Line]
	if(pos.Line < uint(len(indices)-1)){
		nextLine = indices[pos.Line+1]
	} else {
		nextLine = uint(len(str))
	}
	
	for start > lineStart{
		if start != uint(len(str)){		
			if delimiters(str[start]){
				start++
				break
			}
		}
		start--
	}

	for end < nextLine{
		if end != uint(len(str)){
			if delimiters(str[end]){
				break
			}
		}
		end++
	}
	if start >= uint(len(str)){
		start=uint(len(str))

	}
	if start > end{
		start--
	}	
	sym := str[start:end]
	return sym
}
