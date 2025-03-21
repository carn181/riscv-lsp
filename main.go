package main

import (
	"bufio"
	"os"
	"riscv-lsp/logging"
	"riscv-lsp/methods"
	"riscv-lsp/rpc"
	"riscv-lsp/server"
	"riscv-lsp/store"
	"riscv-lsp/symbols"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	logging.Init()
	store.Init()
	writer := os.Stdout
	symbols.Init()
	scanner.Split(rpc.Split)
	
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents , err := rpc.DecodeMessage(msg)

		if err != nil {
			logging.Logger.Fatalf("Got Error: %s\n",err)
			continue
		}

		logging.ReceivedMessage(string(msg))

		handleMethod(writer, method, contents)
	}

}


func handleMethod(wr *os.File, method string, contents []byte){
	switch method {
	case "initialize":
		// Send Initialize Response
		SendMessage(wr, methods.HandleInitialize(contents))
		
	case "initialized":
		// SendMessage(wr, methods.WorkspaceFoldersRequest())
		server.Init()

	// Notifications
	case "textDocument/didChange":
		logging.HandledMessage(string(contents))
		methods.HandleTextDocumentDidChange(contents)
		// TODO: After implementing concurrency, do workspacediagnostics concurrently with a timer and not after every change
		methods.SendWorkspacesDiagnostics(wr)
		
	case "textDocument/didOpen":
		logging.HandledMessage(string(contents))
		methods.HandleTextDocumentDidOpen(contents)

		
	case "textDocument/didClose":
		logging.HandledMessage(string(contents))
		methods.HandleTextDocumentDidClose(contents)

	case "textDocument/hover":
		methods.HandleHover(wr, contents)

	case "shutdown":
		server.Close()
	case "exit":
		os.Exit(0)
	}
}

func SendMessage(wr *os.File, msg any){
	resp := rpc.EncodeMessage(msg)
	wr.Write([]byte(resp))
	logging.SentMessage(string(resp))
}
