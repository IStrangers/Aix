package ast

import (
	"aix/file"
	"aix/token"
	"aix/unistring"
)

type (
	Node interface {
		StartIndex() file.Index
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

	BinaryExpression struct {
		Operator   token.Token
		Left       Expression
		Right      Expression
		Comparison bool
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
		BindingList []*Binding
	}

	LexicalDeclaration struct {
		Index       file.Index
		Token       token.Token
		BindingList []*Binding
	}

	BlockStatement struct {
		LeftBrace  file.Index
		List       []Statement
		RightBrace file.Index
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
		BindingList []*Binding
	}
)

func (self *BadStatement) statementNode()        {}
func (self *ExpressionStatement) statementNode() {}
func (self *VariableStatement) statementNode()   {}
func (self *LexicalDeclaration) statementNode()  {}
func (self *Binding) expressionNode()            {}
func (self *Identifier) expressionNode()         {}
func (self *BadExpression) expressionNode()      {}
func (self *AssignExpression) expressionNode() {
}
func (self *StringLiteral) expressionNode() {
}
func (self BinaryExpression) expressionNode() {
}

func (self *BadStatement) StartIndex() file.Index       { return self.Start }
func (self *Binding) StartIndex() file.Index            { return self.Target.StartIndex() }
func (self *VariableStatement) StartIndex() file.Index  { return self.Var }
func (self *LexicalDeclaration) StartIndex() file.Index { return self.Index }
func (self *Identifier) StartIndex() file.Index         { return self.Index }
func (self *BadExpression) StartIndex() file.Index      { return self.Start }
func (self *AssignExpression) StartIndex() file.Index {
	return self.Left.StartIndex()
}
func (self *StringLiteral) StartIndex() file.Index {
	return self.Index
}
func (self BinaryExpression) StartIndex() file.Index {
	return self.Left.StartIndex()
}

func (self *BadStatement) EndIndex() file.Index { return self.End }
func (self *Binding) EndIndex() file.Index      { return self.Target.EndIndex() }
func (self *VariableStatement) EndIndex() file.Index {
	return self.BindingList[len(self.BindingList)-1].EndIndex()
}
func (self *LexicalDeclaration) EndIndex() file.Index {
	return self.BindingList[len(self.BindingList)-1].EndIndex()
}
func (self *Identifier) EndIndex() file.Index {
	return file.Index(int(self.Index) + len(self.Name))
}
func (self *BadExpression) EndIndex() file.Index { return self.End }
func (self *AssignExpression) EndIndex() file.Index {
	return self.Right.EndIndex()
}
func (self *StringLiteral) EndIndex() file.Index {
	return file.Index(int(self.Index) + len(self.Literal))
}
func (self BinaryExpression) EndIndex() file.Index {
	return self.Right.EndIndex()
}

func (*BadExpression) bindingTarget() {}
func (*Identifier) bindingTarget()    {}
