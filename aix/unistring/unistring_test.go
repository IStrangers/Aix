package unistring

import (
	"testing"
)

func TestNewFromString(t *testing.T) {
	const str = "abc阿瓦达123"
	uniStr := NewFromString(str)
	println(uniStr.String() == str)
}
