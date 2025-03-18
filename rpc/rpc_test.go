package rpc_test

import (
	"riscv-lsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s\n Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 71\r\n\r\n{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"textDocument/completion\",\"params\":{}}"
	method, content, err:= rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 71 {
		t.Fatalf("Expected: 71, Got: %d", contentLength)
	}

	if method != "textDocument/completion" {
		t.Fatalf("Expected: 'textDocument/completion', Got: '%s'", method)
	}
}
