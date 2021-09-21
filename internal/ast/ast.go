package ast

type Node interface {
	TokenLiteral() string
}

type Statment interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statments []Statment
}

func (p *Program) TokenLiteral() string {
	if len(p.Statments) > 0 {
		return p.Statments[0].TokenLiteral()
	} else {
		return ""
	}
}
