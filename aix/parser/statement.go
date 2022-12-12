package parser

import (
	"aix/ast"
	"aix/token"
)

func (self parser) parseScriptStatementList() []ast.Statement {
	return self.parseStatementList(func(tkn token.Token) bool {
		return tkn != token.EOF
	})
}

func (self parser) parseBlockStatementList() []ast.Statement {
	return self.parseStatementList(func(tkn token.Token) bool {
		return tkn != token.RIGHT_BRACE && tkn != token.EOF
	})
}

func (self parser) parseStatementList(endCondition func(token.Token) bool) []ast.Statement {
	var statementList []ast.Statement
	for endCondition(self.token) {
		statementList = append(statementList, self.parseStatement())
	}
	return statementList
}

func (self parser) parseStatement() ast.Statement {
	if self.token == token.EOF {
		return &ast.BadStatement{
			Start: self.index,
			End:   self.index + 1,
		}
	}

	switch self.token {
	case token.VAR:
		return self.parseVariableStatement()
	case token.CONST:
		return self.parseLexicalDeclaration()
	}

	expression := self.parseExpression()

	return &ast.ExpressionStatement{
		Expression: expression,
	}
}

func (self parser) parseVariableStatement() ast.Statement {

}

func (self parser) parseLexicalDeclaration() ast.Statement {

}
