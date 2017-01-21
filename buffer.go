package main

import (
	"bytes"
	"fmt"
)

type Buffer struct {
	// Excluding newlines.  It's easier.  Files without trailing newlines just don't exist. Shut up.
	// Also to make life easier, we always ignore the 0th entry in lines.
	lines    []string
	dot      int
	filename string
}

func (b *Buffer) edit(filename string, data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	b.lines = make([]string, len(lines)+1)
	for pos, line := range lines {
		b.lines[pos+1] = string(line)
	}
	b.filename = filename
	b.dot = len(b.lines) - 1
}

func (b *Buffer) print(first, last int) error {
	if first > last {
		return fmt.Errorf("%d should not be after %d", first, last)
	}
	if first < 1 {
		return fmt.Errorf("first address (%d) should be >=1", first)
	}
	if last >= len(b.lines) {
		return fmt.Errorf("last address (%d) should be < size (%d)", last, len(b.lines))
	}
	for j := first; j <= last; j++ {
		fmt.Println(j, b.lines[j])
	}
	return nil
}
