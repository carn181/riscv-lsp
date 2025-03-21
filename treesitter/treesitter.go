package treesitter

import (
	tree_sitter_riscv "github.com/carn181/tree-sitter-riscv/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

var parser *tree_sitter.Parser

func Init(){
	parser = tree_sitter.NewParser()
	language := tree_sitter.NewLanguage(tree_sitter_riscv.Language())
	parser.SetLanguage(language)
}

func Close(){
	parser.Close()
}
