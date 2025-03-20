package methods

import (
	"riscv-lsp/server"
	"riscv-lsp/types"
)

func WorkspaceFoldersRequest() types.Request {
	req := types.Request{
		RPC: "2.0",
		ID: int(server.S.Currid)+1,
		Method: "workspace/workspaceFolders",
		Params: nil,
	}
	server.S.Currid+=1
	return req
}
