package methods

import (
	"encoding/json"
	"os"
	"riscv-lsp/logging"
	"riscv-lsp/store"
	"riscv-lsp/symbols"
	"riscv-lsp/types"
	"riscv-lsp/utils"
	"strings"
)


func HandleHover(wr *os.File, incoming []byte) {
	var req types.Request
	var params types.TextDocumentPositionParams
	json.Unmarshal(incoming, &req)
	json.Unmarshal(req.Params, &params)

	logging.Logger.Printf("Got params: %v\n", params)
	sym := strings.ToLower(utils.WordAtPos(string(store.FileContents(utils.Uri2Path(params.TextDocument.Uri))), params.Position))
	desc := symbols.Store[symbols.Sym(strings.ToLower(sym))]
	
	logging.Logger.Printf("Got Description for %s as %s\n",sym,desc)
	hovResponseParams := types.Hover{Contents: types.MarkedString(desc)}
	hovParamsRaw, _ := json.Marshal(hovResponseParams)
	hovResponse := types.Response{
		RPC: req.RPC,
		ID: &req.ID,
		Result: hovParamsRaw,
	}
	SendMessage(wr, hovResponse)
}
