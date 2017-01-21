package cli

import (
	"bytes"
	"testing"
)

func TestE(t *testing.T) {
	testData := []struct {
		input string
		cmd   Command
	}{
		{"e /tmp/demo.txt", Command{name: 'e', arg: "/tmp/demo.txt"}},
	}
	for _, td := range testData {
		cmds := parseAll(td.input)
		if len(cmds) != 1 {
			t.Errorf("parseAll(%q) expected one cmd, got %v", td.input, cmds)
		}
		if cmds[0] != td.cmd {
			t.Errorf("parse(%q) got %#v, want %#v", td.input, cmds[0], td.cmd)
		}
	}
}

func parseAll(s string) []Command {
	r := bytes.NewBufferString(s)
	ch := make(chan Command)
	go Parse(r, ch)
	var reply []Command
	for cmd := range ch {
		reply = append(reply, cmd)
	}
	return reply
}