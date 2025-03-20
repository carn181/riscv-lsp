package methods

import (
	"encoding/json"

	"riscv-lsp/logging"
	"riscv-lsp/server"
	"riscv-lsp/types"
)

func HandleInitialize(incoming []byte) types.InitializeResponse {
	var req types.Request
	json.Unmarshal(incoming, &req)
	server.S.RPC = req.RPC

	var params types.InitializeParams
	json.Unmarshal(req.Params, &params)
	server.AddWorkspaces(params.WorkspaceFolders)

	logging.Logger.Printf("Opened Workspaces %v\n", server.S.Workspaces)
	
	return types.InitializeResponse{
		Response: types.Response{RPC: req.RPC, ID: &req.ID},
		Result: types.InitializeResult{
			Capabilities: types.ServerCapabilities{
				TextDocumentSync: types.TextDocumentSync{
					OpenClose: true,
					Change: types.Incremental,
				},
				Workspace: types.Workspace{
					WorkspaceFolders: types.WorkspaceFoldersServerCapabilities{
						Supported: true,
						//ChangeNotifications: "ws",
					},
				},
			},
			ServerInfo: types.ServerInfo{
				Name:    "riscv-lsp",
				Version: "0.0.1",
			},
		},
	}
}
