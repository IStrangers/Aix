package parser

import (
	"aix/ast"
	"aix/token"
)

func (self parser) parseExpression() ast.Expression {
	return nil
}

func (self parser) parseVarDeclarationList() []*ast.Binding {
	return nil
}

func (self parser) parseAssignmentExpression() ast.Expression {
	left := self.parseConditionalExpression()
	var operator token.Token
	switch self.token {
	case token.ASSIGN:
		operator = token.ASSIGN
	}
	return &ast.AssignExpression{
		Left:     left,
		Operator: operator,
		Right:    self.parseAssignmentExpression(),
	}
}

func (self *parser) parseConditionalExpression() ast.Expression {
	return self.parseAdditiveExpression()
}

func (self *parser) parseAdditiveExpression() ast.Expression {
	left := self.parseMultiplicativeExpression()

	for self.token == token.PLUS || self.token == token.MINUS {
		operator := self.token
		self.next()
		left = &ast.BinaryExpression{
			Operator: operator,
			Left:     left,
			Right:    self.parseMultiplicativeExpression(),
		}
	}

	return left
}

func (self *parser) parseMultiplicativeExpression() ast.Expression {
	left := self.parsePrimaryExpression()

	for self.token == token.MULTIPLY || self.token == token.SLASH || self.token == token.REMAINDER {
		operator := self.token
		self.next()
		left = &ast.BinaryExpression{
			Operator: operator,
			Left:     left,
			Right:    self.parsePrimaryExpression(),
		}
	}

	return left
}

func (self *parser) parsePrimaryExpression() ast.Expression {
	index, literal, parsedLiteral := self.index, self.literal, self.parsedLiteral
	switch self.token {
	case token.STRING:
		return &ast.StringLiteral{
			Index:   index,
			Literal: literal,
			Value:   parsedLiteral,
		}
	}

	self.errorUnexpectedToken(self.token)
	self.nextStatement()
	return &ast.BadExpression{Start: index, End: self.index}
}
