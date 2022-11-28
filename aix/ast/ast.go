package ast

// Node
type Node interface {
	//节点在解析文本的开始索引
	StartIndex()
	//节点在解析文本的结束索引
	EndIndex()
}

// Statement
type (
	Statement interface {
		Node
		statement()
	}
)

// Expression
type (
	Expression interface {
		Node
		expression()
	}
)
