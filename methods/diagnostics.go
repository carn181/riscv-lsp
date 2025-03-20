package methods

import (
	"encoding/json"
	"os"
	"path/filepath"
	"riscv-lsp/diagnostics"
	"riscv-lsp/logging"
	"riscv-lsp/rpc"
	"riscv-lsp/server"
	"riscv-lsp/store"
	"riscv-lsp/types"
	"riscv-lsp/utils"
	"slices"
)

// TODO
// Refactor some of the methods here into the diagnostics module
// Improve Variable Names

func SendMessage(wr *os.File, msg any){
	resp := rpc.EncodeMessage(msg)
	wr.Write([]byte(resp))
	logging.SentMessage(string(resp))
}

func ReplicateWorkspace(workspaceBuildPath string, w server.Workspace){
	for _, actualFilePath := range w.Files {
		var shouldModify, different bool
		tmpDirFilePath := filepath.Join(workspaceBuildPath,actualFilePath[len(w.Path):])
		_, tmpDirFileExists := store.GetFile(tmpDirFilePath)
		actualFileOpen := store.IsOpened(actualFilePath)
		
		if tmpDirFileExists{
			actualFileContents, err := os.ReadFile(tmpDirFilePath)
			if err != nil {
				logging.Logger.Fatal(err)
			}
			different = !slices.Equal(store.FileContents(actualFilePath),actualFileContents)
		}

		shouldModify = different || !tmpDirFileExists
		
		if shouldModify{
			if !actualFileOpen {
				store.OpenFileFromDisk(actualFilePath)
			}
			
			utils.EnsurePathExists(tmpDirFilePath)
			if !utils.FileExists(tmpDirFilePath){
				panic("File "+tmpDirFilePath+" does not exist")
			}
			content := store.FileContents(actualFilePath)
			if len(content) != 0 {
				if content[len(content)-1] != '\n'{
					content = append(content, '\n')
				}
			}
			
			err := os.WriteFile(tmpDirFilePath, content , 0700)
			if err != nil {
				logging.Logger.Fatal(err)
			}
		}
	}
}


func RunBuildOnWorkspace(workspaceBuildPath string, w server.Workspace)[]types.Notification{
	buildFilePath := filepath.Join(w.Path,"build")
	logging.Logger.Printf("Running Build %v\n", buildFilePath)
	var notifs []types.Notification	
	if utils.FileExists(buildFilePath){
		buildCommand := store.FileContents(buildFilePath)
		commandOutput, err := utils.RunCommand(workspaceBuildPath, string(buildCommand))
		if err != nil {
			panic(err)
		}
		
		errors := diagnostics.ParseAssemblerErrors(string(commandOutput))

		// Parse all the errors the Build Command Gives
		var fileErrors = make(map[string][]types.Diagnostic)
		for _, error := range(errors){
			fileErrors[error.File] = append(fileErrors[error.File],
				diagnostics.AssemblerError2Diagnostic(w.Path,error))
		}
		// Go through every existing file in Diagnostics and if error doesn't exist for it in FileErrors, clear it out
		for filePath := range diagnostics.Diagnostics{
			_, in := fileErrors[filePath]

			if !in{
				diagnostics.Diagnostics[filePath] = &[]types.Diagnostic{}
			} else {
				diagnostic := fileErrors[filePath]
				diagnostics.Diagnostics[filePath] = &diagnostic
			}
		}

		// If diagnostics doesn't exist, create it
		if diagnostics.Diagnostics == nil{
			diagnostics.Diagnostics = make(map[string]*[]types.Diagnostic)
		}
		// Go through every file in fileErrors. If doesn't exist in Diagnostics, add it		
		for filePath, diagnostic := range fileErrors{
			_, in := diagnostics.Diagnostics[filePath]
			if !in{
				diagnostics.Diagnostics[filePath] = &diagnostic
			}
		}
		

		// Go through every diagnostic in Diagnostic and generate Notifications
		for path, diagnostics := range diagnostics.Diagnostics{
			workspaceFilePath := filepath.Join(w.Path,path)
			store.GetFileVersion(path)
			version := store.GetFileVersion(workspaceFilePath)
			params := types.PublishDiagnosticsParams{
				URI: utils.Path2Uri(workspaceFilePath),
				Version: &version,
				Diagnostics: *diagnostics,
			}

			p, _ := json.Marshal(params)
			notif := types.Notification{
				RPC: server.S.RPC,
				Method: "textDocument/publishDiagnostics",
				Params: p,
			}
			notifs = append(notifs, notif)
		}
	}
	return notifs
}

func SendWorkspacesDiagnostics(wr *os.File){
	logging.Logger.Printf("Sending Workspace Diagnostics\n")
	var notifs []types.Notification = []types.Notification{}

	for _, w := range server.S.Workspaces {
		workspaceBuildPath := filepath.Join(server.S.Tmpdir,w.Name)
		ReplicateWorkspace(workspaceBuildPath, w)
		logging.Logger.Printf("Replicated Workspaces %v\n", server.S.Workspaces)
		workspaceNotifs:=RunBuildOnWorkspace(workspaceBuildPath, w)
		notifs = append(notifs, workspaceNotifs...)
	}
	
	for _, notif := range(notifs){
		//		logging.Logger.Printf("Sending Notification: %s\n",notif)
		SendMessage(wr, notif)
	}
}
