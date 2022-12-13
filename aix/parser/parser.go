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

func newParser(fileName, script string, baseOffset int) *parser {
	return &parser{
		chr:          ' ',
		script:       script,
		scriptLength: len(script),
		baseOffset:   baseOffset,
		file:         file.NewSourceFile(fileName, script, baseOffset),
	}
}

func ParseFileByPath(path string) *ast.Program {
	fsFile, _ := os.Open(path)
	return ParseFile(fsFile)
}

func ParseFile(fsFile fs.File) *ast.Program {
	defer fsFile.Close()
	fileInfo, _ := fsFile.Stat()
	script := make([]byte, fileInfo.Size())
	_, err := fsFile.Read(script)
	if err != nil && err != io.EOF {
		fmt.Println("read buf fail", err)
	}
	return ParseScript(fileInfo.Name(), string(script))
}

func ParseScript(fileName, script string) *ast.Program {
	parser := newParser(fileName, script, 1)
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
		Body:            self.parseScriptStatementList(),
		DeclarationList: self.scope.declarationList,
		File:            self.file,
	}
	return program
}

func (self parser) expect(tkn token.Token) file.Index {
	index := self.index
	if self.token != tkn {
		self.errorUnexpectedToken(self.token)
	}
	self.next()
	return index
}

func (self parser) next() {
	self.token, self.literal, self.parsedLiteral, self.index = self.scan()
}
