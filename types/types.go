package types

import (
	"encoding/json"
)

type TextDocumentIdentifier struct{
	Uri string `json:"uri"`
	Version int `json:"version"`
}

type TextDocumentItem struct{
	Uri string `json:"uri"`
	LanguageId string `json:"languageId"`
	Version int `json:"version"`
	Text string `json:"text"`
}

type Position struct {
	Line uint      `json:"line"`
	Character uint `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End Position   `json:"end"`
}

type TextDocumentContentChangeEvent struct {
	Range         Range  `json:"range"`
	RangeLength   uint   `json:"rangeLength"`
	Text          string `json:"text"`
}

type DidChangeTextDocumentParams struct {
	TextDocument TextDocumentIdentifier             `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent   `json:"contentChanges"`
}

type DidOpenDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type Request struct {
	RPC string `json:"jsonrpc"`
	ID int `json:"id"`
	Method string `json:"method"`
	Params json.RawMessage `json:"params"`
}

type Response struct {
	RPC string `json:"jsonrpc"`	
	ID *int `json:"id,omitempty"`

	// Result
	// Error
}

type Notification struct {
	JSONRPC string `json:"jsonrpc"`
	Method string `json:"method"`
}

type TextDocumentSyncKind int

const (
	None = 0
	Full = 1
	Incremental = 2
)

type TextDocumentSync struct {
	OpenClose bool `json:"openClose"`
	Change TextDocumentSyncKind `json:"change"`
}

type ServerCapabilities struct {
	//	HoverProvider bool `json:"hoverProvider"`
	TextDocumentSync TextDocumentSync `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name string    `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities  `json:"capabilities"`
	ServerInfo ServerInfo `json:"serverInfo"`
}



