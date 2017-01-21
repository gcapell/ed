package cli

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Command struct {
	parseError          string // We couldn't parse this
	name                byte   // e -> edit, ...
	addr1, addr2, addr3 string // optional address
	arg                 string // e.g. text to insert.  May need to expand on this for complex commands
}

func (c Command) String() string {
	if c.parseError != "" {
		return fmt.Sprintf("Command err %s", c.parseError)
	}
	return fmt.Sprintf("Command{%s,%s,%s,%s [%s]}",
		string(c.name), c.addr1, c.addr2, c.addr3, c.arg)
}

func Parse(r io.Reader, ch chan Command) {
	s := bufio.NewScanner(r)
	defer close(ch)
	for {
		if cmd, eof := parse(s); eof {
			return
		} else {
			ch <- cmd
		}
	}
	close(ch)
}

func parse(s *bufio.Scanner) (Command, bool) {
	if !s.Scan() {
		if err := s.Err(); err != nil {
			return errorf(err.Error())
		}
		return Command{}, true
	}
	line := s.Text()
	log.Printf("parsing %q", line)
	// let's be really stupid
	chunks := strings.Fields(line)
	if len(chunks) < 1 {
		return errorf("need at least one field")
	}
	if len(chunks[0]) != 1 {
		return errorf("first field expected to be one command character, got %q", chunks[0])
	}
	if len(chunks) > 2 {
		return errorf("expected at most one arg")
	}
	cmd := Command{name: chunks[0][0]}
	if len(chunks) == 2 {
		cmd.arg = chunks[1]
	}
	return cmd, false
}

func errorf(format string, args ...interface{}) (Command, bool) {
	return Command{parseError: fmt.Sprintf(format, args)}, false
}
