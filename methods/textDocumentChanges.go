package methods

import (
	"encoding/json"
	"riscv-lsp/logging"
	"riscv-lsp/store"
	"riscv-lsp/types"
	"riscv-lsp/utils"
)

func HandleTextDocumentDidChange(incoming []byte) {
	var req types.Request
	json.Unmarshal(incoming, &req)
	var params types.DidChangeTextDocumentParams
	json.Unmarshal(req.Params,&params)
	
	path:=utils.Uri2Path(params.TextDocument.Uri)
	if !store.IsOpened(path){
		logging.Logger.Fatalf("Crashing because %s is not opened already\n", path)
	} else {
		logging.Logger.Printf("Params: %v\n", params.ContentChanges[0].Range)
		store.ModifyFile(path, params.ContentChanges[0].Range, params.TextDocument.Version, params.ContentChanges[0].Text)
	}
	
}

func HandleTextDocumentDidOpen(incoming []byte) {
	var req types.Request
	json.Unmarshal(incoming, &req)
	var params types.DidOpenDocumentParams
	json.Unmarshal(req.Params,&params)

	path:=utils.Uri2Path(params.TextDocument.Uri)
	store.OpenFile(path,params.TextDocument.Version, params.TextDocument.Text)
}

func HandleTextDocumentDidClose(incoming []byte) {
	var req types.Request
	json.Unmarshal(incoming, &req)
	var params types.DidOpenDocumentParams
	json.Unmarshal(req.Params,&params)

	path:=utils.Uri2Path(params.TextDocument.Uri)
	store.CloseFile(path)
}
