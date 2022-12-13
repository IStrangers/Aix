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

func (self parser) errorUnexpectedToken(tkn token.Token) {

}
