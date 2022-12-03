package parser

import "aix/ast"

type scope struct {
	outer           *scope
	declarationList []*ast.VariableDeclaration
}

func (self parser) openScope() {
	self.scope = &scope{
		outer: self.scope,
	}
}

func (self parser) closeScope() {
	self.scope = self.scope.outer
}
