package parser

import (
	"aix/file"
	"aix/token"
	"fmt"
)

type Error struct {
	Position file.Position
	Message  string
}

func (self Error) Error() string {
	fileName := self.Position.FileName
	if fileName == "" {
		fileName = "(anonymous)"
	}
	return fmt.Sprintf("%s: Line %d:%d %s",
		fileName,
		self.Position.Line,
		self.Position.Column,
		self.Message,
	)
}

type ErrorList []*Error

func (self ErrorList) Error() string {
	switch len(self) {
	case 0:
		return "no errors"
	case 1:
		return self[0].Error()
	}
	return fmt.Sprintf("%s (and %d more errors)", self[0].Error(), len(self)-1)
}

func (self ErrorList) PeekErr() error {
	if len(self) == 0 {
		return nil
	}
	return self
}

func (self *ErrorList) Add(position file.Position, msg string) {
	*self = append(*self, &Error{position, msg})
}

func (self parser) error(place file.Index, msg string, msgValues ...any) *Error {
	position := self.position(place)
	msg = fmt.Sprintf(msg, msgValues...)
	self.errorList.Add(position, msg)
	return self.errorList[len(self.errorList)-1]
}

func (self parser) errorUnexpectedToken(tkn token.Token) *Error {
	switch tkn {
	case token.EOF:
		return self.error(file.Index(0), "Unexpected end of input")
	}
	value := tkn.String()
	switch tkn {
	case token.BOOLEAN, token.NULL:
		value = self.literal
	case token.IDENTIFIER:
		return self.error(self.index, "Unexpected identifier")
	case token.KEYWORD:
		return self.error(self.index, "Unexpected reserved word")
	case token.NUMBER:
		return self.error(self.index, "Unexpected number")
	case token.STRING:
		return self.error(self.index, "Unexpected string")
	}
	return self.error(self.index, "Unexpected token %v", value)
}
