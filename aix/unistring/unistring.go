package unistring

import (
	"unicode/utf8"
)

const (
	BOM       = 0xFEFF
	BOMLength = 3
)

type UniString string

func NewFromBytes(bytes []byte) UniString {
	return UniString(bytes)
}

func NewFromString(s string) UniString {
	utf8Size := 0
	for _, chr := range s {
		utf8Size += utf8.RuneLen(chr)
	}
	headBuf := make([]byte, BOMLength)
	utf8.EncodeRune(headBuf, BOM)

	buf := make([]byte, utf8Size+BOMLength)
	index := 0
	for _, b := range headBuf {
		buf[index] = b
		index++
	}

	for _, chr := range s {
		chrLen := utf8.RuneLen(chr)
		chrBuf := make([]byte, chrLen)
		utf8.EncodeRune(chrBuf, chr)
		for _, b := range chrBuf {
			buf[index] = b
			index++
		}
	}

	return NewFromBytes(buf)
}

func NewFromRunes(s []rune) UniString {
	return NewFromString(string(s))
}

func (s UniString) AsBytes() []byte {
	buf := []byte(s)
	if BOMLength > len(buf) {
		return nil
	}
	head, _ := utf8.DecodeRune(buf[0:BOMLength])
	if head == BOM {
		return buf
	}
	return nil
}

func (s UniString) String() string {
	if bytes := s.AsBytes(); bytes != nil {
		return string(bytes[BOMLength:])
	}
	return string(s)
}
