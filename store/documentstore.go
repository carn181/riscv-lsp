package store

import (
	"fmt"
	//	"log"
	"riscv-lsp/logging"
	"riscv-lsp/types"
)
type File struct{
	content []byte
	version int
	open    bool
}

type DocumentStore struct{
	files         map[string]File
}

var Store DocumentStore

func Init(){
	Store.files = make(map[string]File)
}

func OpenFile(path string, content string){
	logging.Logger.Printf("Opening %s\n", path)	
	Store.files[path] = File{[]byte(content),-1,true}
	//	logging.Logger.Printf("File Buffer:\n%s\n",FileContents(path))
	//	indices := FindNewLineIndices(string(FileContents(path)))
	//	logging.Logger.Printf("Found New Line Indexes %v\n",indices)
}

func CloseFile(path string){
	f:=Store.files[path]
	f.open=false
	Store.files[path]=f
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

func ApplyChange(orig string, rng types.Range, replace string) string{
	start, err := PositionToOffset(orig, rng.Start)
	if err != nil {
		logging.Logger.Fatal(err)
	}
	end, err := PositionToOffset(orig, rng.End)
	if err != nil {
		logging.Logger.Fatal(err)
	}
	
	//	logging.Logger.Printf("Found Start and End Indexes for Range %d-%d\n",start,end)
	
	modified := ReplaceStrSlice(orig, start, end, replace)
	return modified
}
func PositionToOffset(str string, pos types.Position) (uint, error){
	// When taking slice ranges, go handles [0:] and [:0] for empty strings
	// So it's ok to have offset = len(str)
	indices := FindNewLineIndices(str)
	if pos.Line > uint(len(indices)) {
		return 0, fmt.Errorf("Line Number %d out of range 0-%d",pos.Line,len(indices))
	} else if pos.Line == uint(len(indices)) {
		if pos.Character == 0 {
			return uint(len(str)), nil
		}
		return 0, fmt.Errorf("Column is beyond end of file")
	}
	
	lineoffset := indices[pos.Line]
	content := str[lineoffset:]
	var i uint
	for  ; i < pos.Character; i++{
		if i > uint(len(content)){
			return 0, fmt.Errorf("Column is beyond end of file")
		}
		if len(content) != 0{
			if content[i] == '\n'{
				return 0, fmt.Errorf("Column is beyond end of line")
			}
		}
	}
	return lineoffset + i, nil
}

func ModifyFile(path string, rng types.Range, text string){
	logging.Logger.Printf("Modifying %s\n",path)
	orig  := string(Store.files[path].content)
	modified := ApplyChange(orig, rng, text)
	Store.files[path] = File{content: []byte(modified)}	
	//	logging.Logger.Printf("File Buffer:\n%s\n",FileContents(path))
}

func FileContents(path string) []byte{return Store.files[path].content}

func IsOpened(path string) bool{
	_, open := Store.files[path]
	return open
}
