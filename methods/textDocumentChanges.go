package methods

import (
	"encoding/json"
	"riscv-lsp/logging"
	"riscv-lsp/store"
	"riscv-lsp/types"
)

func uri2path(uri string) string{
	return uri[7:]    // file:// is 7 chars
}

func HandleTextDocumentDidChange(incoming []byte) {
	var req types.Request
	json.Unmarshal(incoming, &req)
	var params types.DidChangeTextDocumentParams
	json.Unmarshal(req.Params,&params)
	
	path:=uri2path(params.TextDocument.Uri)
	if !store.IsOpened(path){
		logging.Logger.Fatalf("Crashing because %s is not opened already\n", path)
	} else {
		logging.Logger.Printf("Params: %v\n", params.ContentChanges[0].Range)
		store.ModifyFile(path, params.ContentChanges[0].Range, params.ContentChanges[0].Text)
	}
	
}

func HandleTextDocumentDidOpen(incoming []byte) {
	var req types.Request
	json.Unmarshal(incoming, &req)
	var params types.DidOpenDocumentParams
	json.Unmarshal(req.Params,&params)

	path:=uri2path(params.TextDocument.Uri)
	store.OpenFile(path, params.TextDocument.Text)
}

func HandleTextDocumentDidClose(incoming []byte) {
	var req types.Request
	json.Unmarshal(incoming, &req)
	var params types.DidOpenDocumentParams
	json.Unmarshal(req.Params,&params)

	path:=uri2path(params.TextDocument.Uri)
	store.CloseFile(path)
}
