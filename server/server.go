package server

import (
	"riscv-lsp/types"
	"riscv-lsp/logging"
	"riscv-lsp/utils"
	"os"
)

type Server struct{
	Currid uint
	Workspaces []Workspace
	RPC    string
	// Directory where we do builds
	Tmpdir string 
}

type Workspace struct{
	Path string
	Files []string
	Name string
}

var S Server

func Init(){
	S.Currid = 0
	var err error
	S.Tmpdir, err = os.MkdirTemp("/tmp/", "riscv-build-")
	if err != nil{
		panic(err)
	}
	logging.Logger.Printf("Initialized Server\n")
}

func Close(){
	os.RemoveAll(S.Tmpdir)	
}

func WorkspaceFolderToWorkspace(wf types.WorkspaceFolder) Workspace{
	path := utils.Uri2Path(wf.Uri)
	files, err := utils.FilesInPath(path)
	if err != nil {
		logging.Logger.Fatal(err)
	}
	return Workspace{
		Path: path,
		Files: files,
		Name: wf.Name,
	}
}

func WorkspaceFoldersToWorkspaces(wfs []types.WorkspaceFolder) []Workspace{
	var ws []Workspace
	for _, wf := range wfs {
		ws = append(ws, WorkspaceFolderToWorkspace(wf))
	}
	return ws
}

func AddWorkspaces(wfs []types.WorkspaceFolder){
	S.Workspaces = append(S.Workspaces,WorkspaceFoldersToWorkspaces(wfs)...)
}
