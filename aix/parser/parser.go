package parser

import (
	"aix/ast"
	"aix/file"
	"aix/token"
	"aix/unistring"
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

func ParseFileByPath(path string) (*ast.Program, error) {
	script, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseScript(path, string(script))
}

func ParseScript(fileName, script string) (*ast.Program, error) {
	parser := newParser(fileName, script, 1)
	return parser.parse()
}

func (self parser) parse() (*ast.Program, error) {
	defer self.closeScope()
	self.openScope()
	program := self.parseProgram()
	return program, self.errorList.PeekErr()
}

func (self parser) parseProgram() *ast.Program {
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

func (self parser) position(index file.Index) file.Position {
	return self.file.Position(int(index) - self.baseOffset)
}

func (self parser) next() {
	self.token, self.literal, self.parsedLiteral, self.index = self.scan()
}
