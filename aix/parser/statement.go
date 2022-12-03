package parser

import (
	"aix/ast"
	"aix/token"
)

func (self parser) parseStatementList() []ast.Statement {
	var statementList []ast.Statement
	for self.token != token.EOF {
		statementList = append(statementList, self.parseStatement())
	}
	return statementList
}

func (self parser) parseStatement() ast.Statement {
	return nil
}
