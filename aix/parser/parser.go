package parser

import (
	"aix/ast"
	"aix/file"
	"aix/token"
	"aix/unistring"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type parser struct {
	script       string
	scriptLength int
	baseOffset   int

	chr       rune
	chrOffset int
	offset    int

	index         file.Index
	token         token.Token
	literal       string
	parsedLiteral unistring.UniString

	scope *scope

	errorList ErrorList

	file *file.SourceFile
}

func NewParser(fileName, script string, baseOffset int) *parser {
	parser := &parser{
		chr:          ' ',
		script:       script,
		scriptLength: len(script),
		baseOffset:   baseOffset,
		file:         file.NewSourceFile(fileName, script, baseOffset),
	}
	return parser
}

func ParseFileByPath(path string) *ast.Program {
	file, _ := os.Open(path)
	return ParseFile(file)
}

func ParseFile(file fs.File) *ast.Program {
	defer file.Close()
	fileInfo, _ := file.Stat()
	script := make([]byte, fileInfo.Size())
	_, err := file.Read(script)
	if err != nil && err != io.EOF {
		fmt.Println("read buf fail", err)
	}
	return ParseScript(fileInfo.Name(), string(script))
}

func ParseScript(fileName, script string) *ast.Program {
	parser := NewParser(fileName, script, 1)
	return parser.parseProgram()
}

func (self parser) parse() (*ast.Program, error) {
	program := self.parseProgram()
	return program, self.errorList.PeekErr()
}

func (self parser) parseProgram() *ast.Program {
	defer self.closeScope()
	self.openScope()
	program := &ast.Program{
		Body:            self.parseStatementList(),
		DeclarationList: self.scope.declarationList,
		File:            self.file,
	}
	return program
}
