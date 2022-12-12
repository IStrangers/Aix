package ast

import (
	"aix/file"
	"aix/token"
	"aix/unistring"
)

type (
	Node interface {
		//节点在解析文本的开始索引
		StartIndex() file.Index
		//节点在解析文本的结束索引
		EndIndex() file.Index
	}

	Expression interface {
		Node
		expressionNode()
	}

	BindingTarget interface {
		Expression
		bindingTarget()
	}

	Binding struct {
		Target      BindingTarget
		Initializer Expression
	}

	Pattern interface {
		BindingTarget
		pattern()
	}

	Statement interface {
		Node
		statementNode()
	}

	Program struct {
		Body            []Statement
		DeclarationList []*VariableDeclaration
		File            *file.SourceFile
	}
)

// Expression
type (
	BadExpression struct {
		Start file.Index
		End   file.Index
	}

	ExpressionStatement struct {
		Expression
	}

	AssignExpression struct {
		Left     Expression
		Operator token.Token
		Right    Expression
	}

	ParameterList struct {
		Opening     file.Index
		BindingList []*Binding
		Closing     file.Index
	}

	StringLiteral struct {
		Index   file.Index
		Literal string
		Value   unistring.UniString
	}

	FunctionLiteral struct {
		Function      file.Index
		Name          *Identifier
		ParameterList *ParameterList
		Body          *BlockStatement
	}
)

// Statement
type (
	BadStatement struct {
		Start file.Index
		End   file.Index
	}

	Identifier struct {
		Index file.Index
		Name  unistring.UniString
	}

	VariableStatement struct {
		Var         file.Index
		BindingList []Binding
	}

	BlockStatement struct {
		LeftBrace     file.Index
		StatementList []Statement
		RightBrace    file.Index
	}

	ReturnStatement struct {
		Return   file.Index
		Argument Expression
	}
)

// Declaration
type (
	VariableDeclaration struct {
		Var         file.Index
		BindingList []Binding
	}

	ClassDefinition struct {
		Node
	}

	ClassStaticBlock struct {
		Static file.Index
		Block  *BlockStatement
	}

	FieldDefinition struct {
		Index       file.Index
		Key         Expression
		Initializer Expression
		Static      bool
	}

	MethodDefinition struct {
		Index  file.Index
		Key    Expression
		Body   *FunctionLiteral
		Static bool
	}
)

func (self BadStatement) StartIndex() file.Index {
	return self.Start
}
func (self BadStatement) EndIndex() file.Index {
	return self.End
}

func (self BadStatement) statementNode()     {}
func (e ExpressionStatement) statementNode() {}

func (*BadExpression) bindingTarget() {}
func (*Identifier) bindingTarget()    {}
