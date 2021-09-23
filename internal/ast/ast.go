package ast

import "lioslang/internal/token"

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
	Statements []Statment
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type DefStatement struct {
	Token token.Token
	Name  *Idenfier
	Value Expression
}

func (ls *DefStatement) statementNode() {}
func (ls *DefStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type Idenfier struct {
	Token token.Token
	Value string
}

func (i *Idenfier) expressionNode() {}
func (i *Idenfier) TokenLiteral() string {
	return i.Token.Literal
}
