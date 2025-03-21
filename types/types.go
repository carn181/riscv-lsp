package types

import (
	"encoding/json"
)

// General

type Position struct {
	Line uint      `json:"line"`
	Character uint `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End Position   `json:"end"`
}


type ServerInfo struct {
	Name string    `json:"name"`
	Version string `json:"version"`
}

// Base

type Request struct {
	RPC string `json:"jsonrpc"`
	ID int `json:"id"`
	Method string `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	RPC string `json:"jsonrpc"`	
	ID *int `json:"id,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	Error json.RawMessage `json:"error,omitempty"`
}

type Notification struct {
	RPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}


// TextDocument

type TextDocumentIdentifier struct{
	Uri string `json:"uri"`
}

type VersionedTextDocumentIdentifier struct{
	Uri string `json:"uri"`
	Version int `json:"version"`
}

type TextDocumentItem struct{
	Uri string `json:"uri"`
	LanguageId string `json:"languageId"`
	Version int `json:"version"`
	Text string `json:"text"`
}

type TextDocumentContentChangeEvent struct {
	Range         Range  `json:"range"`
	RangeLength   uint   `json:"rangeLength"`
	Text          string `json:"text"`
}

// DidChange

type DidChangeTextDocumentParams struct {
	TextDocument VersionedTextDocumentIdentifier             `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent   `json:"contentChanges"`
}

// DidOpen
type DidOpenDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// Capabilities

type ServerCapabilities struct {
	HoverProvider bool `json:"hoverProvider"`
	TextDocumentSync TextDocumentSync `json:"textDocumentSync"`
	Workspace Workspace `json:"workspace"`
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

// Initialize Params
type InitializeParams struct{
	ProcessId int `json:"processId"`
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

// Initialize 
type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities  `json:"capabilities"`
	ServerInfo ServerInfo `json:"serverInfo"`
}

// Workspaces
type Workspace struct {
	WorkspaceFolders WorkspaceFoldersServerCapabilities `json:"workspaceFolders"`
}

type WorkspaceFoldersServerCapabilities struct {
	Supported bool `json:"supported"`
	//	ChangeNotifications string `json:"changeNotifications"`
}

type WorkspaceFolder struct{
	Uri string `json:"uri"`
	Name string `json:"name"`
}

// Diagnostics
type Diagnostic struct {
	Range           Range `json:"range"`
	Severity        int    `json:"severity,omitempty"`
	Code            string `json:"code,omitempty"`
	CodeDescription struct {
		Description string `json:"description"`
	} `json:"codeDescription,omitempty"`
	Source             string   `json:"source.omitempty"`
	Message            string   `json:"message,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	RelatedInformation []struct {
		RelatedInfo string `json:"relatedInfo"`
	} `json:"relatedInformation,omitempty"`
	Data               string `json:"data"`
}

type PublishDiagnosticsParams struct {
	URI        string  `json:"uri"`
	Version    *int          `json:"version,omitempty"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

// Hover Support
type TextDocumentPositionParams struct {
  TextDocument TextDocumentIdentifier `json:"textDocument"`
  Position     Position               `json:"position"`
}

// Markup
// A string for now
type MarkedString string

// Define the Hover struct
type Hover struct {
  Contents MarkedString `json:"contents"` // Can be MarkedString, []MarkedString, or MarkupContent
  Range    *Range      `json:"range,omitempty"` // Optional field
}


// Symbols

type Register struct {
	Register    string `json:"register"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type RegistersList struct {
	Registers []Register `json:"Registers"`
}

type Opcode struct {
	Name           string `json:"name"`
	Format         string `json:"format"`
	Description    string `json:"description"`
	Implementation string `json:"implementation,omitempty"`
	Expansion      string `json:"expansion,omitempty"`
	Sdescription   string `json:"sdescription,omitempty"`
} 

type OpcodesList struct {
	Opcodes []Opcode `json:"Opcodes"`
}
