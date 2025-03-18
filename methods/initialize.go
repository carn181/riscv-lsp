package methods

import (
	"encoding/json"
	"riscv-lsp/types"
)

func HandleInitialize(incoming []byte) types.InitializeResponse {
	var req types.Request
	json.Unmarshal(incoming, &req)
	// For Future Syncing up with Client Capabilities
	_ = req

	return types.InitializeResponse{
		Response: types.Response{RPC: req.RPC, ID: &req.ID},
		Result: types.InitializeResult{
			Capabilities: types.ServerCapabilities{
				TextDocumentSync: types.TextDocumentSync{OpenClose: true, Change: types.Incremental}},
			ServerInfo: types.ServerInfo{
				Name:    "riscv-lsp",
				Version: "0.0.1"},
		},
	}
}
