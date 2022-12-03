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
	script string
	baseOffset int

	chr rune
	chrOffset int
	offset int

	index file.Index
	token token.Token
	literal string
	parsedLiteral unistring.UniString

	errorList ErrorList

	file *file.SourceFile
}

func ParseFileByPath(path string) *ast.Program {
	script, _ := os.ReadFile(path)
	return ParseScript(string(script))
}

func ParseFile(file fs.File) *ast.Program {
	defer file.Close()
	var script []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			break
		}
		if n == 0 {
			break
		}
		script = append(script, buf[:n]...)
	}
	return ParseScript(string(script))
}

func ParseScript(script string) *ast.Program {

}

func (self parser) parse() (*ast.Program, error) {
	program := self.parseProgram()
	return program,self.errorList.PeekErr()
}

func (self parser) parseProgram() *ast.Program {
	program := &ast.Program{
		Body: self.parseStatementList(),
		DeclarationList: ,
		File: self.file,
	}
	return program
}