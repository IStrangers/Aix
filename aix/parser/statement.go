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

func (self parser) parseBlockStatementList() *ast.BlockStatement {
	return &ast.BlockStatement{
		LeftBrace: self.expect(token.LEFT_BRACE),
		List: self.parseStatementList(func(tkn token.Token) bool {
			return tkn != token.RIGHT_BRACE && tkn != token.EOF
		}),
		RightBrace: self.expect(token.RIGHT_BRACE),
	}
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
		return self.parseLexicalDeclaration(token.CONST)
	}

	expression := self.parseExpression()

	return &ast.ExpressionStatement{
		Expression: expression,
	}
}

func (self parser) parseVariableStatement() ast.Statement {
	index := self.expect(token.VAR)
	list := self.parseVarDeclarationList()

	return &ast.VariableStatement{
		Var:         index,
		BindingList: list,
	}
}

func (self parser) parseLexicalDeclaration(tkn token.Token) *ast.LexicalDeclaration {
	index := self.expect(tkn)
	list := self.parseVariableDeclarationList()

	return &ast.LexicalDeclaration{
		Index:       index,
		Token:       tkn,
		BindingList: list,
	}
}

func (self parser) parseVariableDeclarationList() (declarationList []*ast.Binding) {
	for {
		self.parseVariableDeclaration(&declarationList)
		if self.token != token.COMMA {
			break
		}
		self.next()
	}
	return
}

func (self parser) parseVariableDeclaration(declarationList *[]*ast.Binding) ast.Expression {
	node := &ast.Binding{
		Target: self.parseBindingTarget(),
	}

	if declarationList != nil {
		*declarationList = append(*declarationList, node)
	}

	if self.token == token.ASSIGN {
		self.next()
		node.Initializer = self.parseAssignmentExpression()
	}

	return node
}

func (self parser) parseBindingTarget() (target ast.BindingTarget) {
	return
}

func (self parser) nextStatement() {

}
