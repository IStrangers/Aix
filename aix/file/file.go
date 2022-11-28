package file

import (
	"fmt"
	"sort"
	"sync"
)

type Index int

type Position struct {
	FileName string
	Line     int
	Column   int
}

func (self *Position) isValid() bool {
	return self.Line > 0
}

func (self *Position) String() string {
	str := self.FileName
	if self.isValid() {
		if str != "" {
			str += ":"
		}
		str += fmt.Sprintf("%d:%d", self.Line, self.Column)
	}
	if str == "" {
		str = "-"
	}
	return str
}

type SourceFile struct {
	mutexLock         sync.Mutex
	index             Index
	name              string
	content           string
	lineOffsets       []int
	lastScannedOffset int
}

func NewSourceFile(fileName, content string, index Index) *SourceFile {
	return &SourceFile{
		index:   index,
		content: content,
		name:    fileName,
	}
}

func (self *SourceFile) Position(offset int) Position {
	var line int
	var lineOffsets []int
	self.mutexLock.Lock()
	if offset > self.lastScannedOffset {
		line = self.scanLine(offset)
		lineOffsets = self.lineOffsets
		self.mutexLock.Unlock()
	} else {
		lineOffsets = self.lineOffsets
		self.mutexLock.Unlock()
		line = sort.Search(len(lineOffsets), func(index int) bool {
			return lineOffsets[index] > offset
		}) - 1
	}

	var lineStart int
	if line >= 0 {
		lineStart = lineOffsets[line]
	}

	row := line + 2
	col := offset - lineStart + 1

	return Position{
		FileName: self.name,
		Line:     row,
		Column:   col,
	}
}

func findNextLineStart(s string) int {
	for pos, ch := range s {
		switch ch {
		case '\r':
			if pos < len(s)-1 && s[pos+1] == '\n' {
				return pos + 2
			}
			return pos + 1
		case '\n':
			return pos + 1
		case '\u2028', '\u2029':
			return pos + 3
		}
	}
	return -1
}

func (self *SourceFile) scanLine(offset int) int {
	o := self.lastScannedOffset
	for o < offset {
		p := findNextLineStart(self.content[o:])
		if p == -1 {
			self.lastScannedOffset = len(self.content)
			return len(self.lineOffsets) - 1
		}
		o = o + p
		self.lineOffsets = append(self.lineOffsets, o)
	}
	self.lastScannedOffset = o

	if o == offset {
		return len(self.lineOffsets) - 1
	}

	return len(self.lineOffsets) - 2
}

type SourceFileSet struct {
	files []*SourceFile
	last  *SourceFile
}

func (self *SourceFileSet) nextIndex() Index {
	if self.last == nil {
		return 0
	}
	return self.last.index + 1
}

func (self *SourceFileSet) Add(fileName, content string) Index {
	index := self.nextIndex()
	sourceFile := NewSourceFile(fileName, content, index)
	self.files = append(self.files, sourceFile)
	self.last = sourceFile
	return index
}

func (self *SourceFileSet) get(index Index) *SourceFile {
	return self.files[index]
}
