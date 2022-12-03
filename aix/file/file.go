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
	baseOffset        int
	name              string
	script            string
	lineOffsets       []int
	lastScannedOffset int
}

func NewSourceFile(fileName, script string, baseOffset int) *SourceFile {
	return &SourceFile{
		baseOffset: baseOffset,
		script:     script,
		name:       fileName,
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
		p := findNextLineStart(self.script[o:])
		if p == -1 {
			self.lastScannedOffset = len(self.script)
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

func (self *SourceFileSet) nextBaseOffset() int {
	if self.last == nil {
		return 1
	}
	return self.last.baseOffset + len(self.last.script) + 1
}

func (self *SourceFileSet) Add(fileName, content string) int {
	baseOffset := self.nextBaseOffset()
	sourceFile := NewSourceFile(fileName, content, baseOffset)
	self.files = append(self.files, sourceFile)
	self.last = sourceFile
	return baseOffset
}

func (self *SourceFileSet) get(index Index) *SourceFile {
	for _, file := range self.files {
		if index <= Index(file.baseOffset+len(file.script)) {
			return file
		}
	}
	return nil
}

func (self *SourceFileSet) Position(index Index) Position {
	for _, file := range self.files {
		if index <= Index(file.baseOffset+len(file.script)) {
			return file.Position(int(index) - file.baseOffset)
		}
	}
	return Position{}
}
