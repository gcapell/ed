package cli

import (
	"bytes"
	"testing"
)

var singleLineCommands = []struct {
	input string
	cmd   Command
}{
	{"e /tmp/demo.txt", Command{name: 'e', arg: "/tmp/demo.txt"}},
	{"1,3p", Command{name: 'p', addr1: "1", addr2: "3"}},
	{"p", Command{name: 'p'}},
	{"1 p", Command{name: 'p', addr1: "1"}},
}

func TestParseLine(t *testing.T) {
	for _, td := range singleLineCommands {
		if cmd := parseLine(td.input); cmd != td.cmd {
			t.Errorf("parseLine(%q) got %s, want %s", td.input, cmd, td.cmd)
		}
	}
}

func TestParse(t *testing.T) {
	input := ""
	for _, td := range singleLineCommands {
		input += td.input + "\n"
	}

	r := bytes.NewBufferString(input)
	ch := make(chan Command)
	go Parse(r, ch)

	pos := 0
	for cmd := range ch {
		if pos == len(singleLineCommands) {
			t.Fatal("more commands returned than expected")
		}
		td := singleLineCommands[pos]
		pos++
		if cmd != td.cmd {
			t.Errorf("parse(%q) got %s, want %s", td.input, cmd, td.cmd)
		}
	}
}
