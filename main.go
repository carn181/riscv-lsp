package main

import (
	"bufio"
	"os"
	"riscv-lsp/logging"
	"riscv-lsp/methods"
	"riscv-lsp/rpc"
	"riscv-lsp/store"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	logging.Init()
	store.Init()
	writer := os.Stdout
	
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
	switch method{
		case "initialize":
		resp := rpc.EncodeMessage(methods.HandleInitialize(contents))
		wr.Write([]byte(resp))		
		logging.SentMessage(string(resp))
		case "textDocument/didChange":
		methods.HandleTextDocumentDidChange(contents)
		case "textDocument/didOpen":
		methods.HandleTextDocumentDidOpen(contents)
		case "textDocument/didClose":
		methods.HandleTextDocumentDidClose(contents)
		case "shutdown":
		
		case "exit":
		os.Exit(0)
	}
}

