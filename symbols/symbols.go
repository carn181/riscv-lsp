package symbols

import (
	//	"encoding/json"
	//	"os"
	//	"path/filepath"
	//	"riscv-lsp/types"
	//	"riscv-lsp/logging"
	//	"riscv-lsp/utils"
)


type SymbolDescription string
type Sym string

type SymbolsStore map[Sym]SymbolDescription

var Store SymbolsStore
// CHANGE THIS BEFORE YOU BUILD PLEASE
//var symbolsDir string = "riscv-docs/"

func Init() error{
	//	files, err := utils.FilesInPath(symbolsDir)
	//	if err != nic {
	//		panic(err)
	//	}
	Store = make(SymbolsStore)

	// regsFile, err := filepath.Abs(filepath.Join(symbolsDir+"regs.json"))
	// if err != nil{return err}	
	// opsFile, err := filepath.Abs(filepath.Join(symbolsDir+"opcodes.json"))
	// if err != nil{return err}

	// if logging.Logger != nil {
	// 	logging.Logger.Printf("Getting Symbols From %s and %s\n", regsFile, opsFile)
	// }
	
	// regsContent, err := os.ReadFile(regsFile)
	// if err != nil{return(err)}
	// opsContent, err := os.ReadFile(opsFile)
	// if err != nil{return(err)}
	
	// var regs types.RegistersList
	// var ops types.OpcodesList

	// json.Unmarshal(regsContent, &regs)
	// json.Unmarshal(opsContent, &ops)
	
	//	fmt.Println(len(regs.Registers))
	//	fmt.Println(len(ops.Opcodes))

	for _, reg := range regs{
		Store[Sym(reg.Name)] = SymbolDescription(reg.Description)
		Store[Sym(reg.Register)] = SymbolDescription(reg.Description)		
	}

	for _, op := range ops{
		Store[Sym(op.Name)] = SymbolDescription(op.Description)
	}


	return nil
}
