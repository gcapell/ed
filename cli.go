package cli

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Command struct {
	name                byte   // e -> edit, ...
	addr1, addr2, addr3 string // optional address
	arg                 string // e.g. text to insert.  May need to expand on this for complex commands
}

func Parse(r io.Reader, ch chan Command) {
	s := bufio.NewScanner(r)
	defer close(ch)
	for {
		cmd, err, eof := parse(s)
		switch {
		case eof:
			return
		case err != nil:
			fmt.Fprint(os.Stderr, err)
		default:
			ch <- cmd
		}
	}
	close(ch)
}

func parse(s *bufio.Scanner) (Command, error, bool) {
	if !s.Scan() {
		err := s.Err()
		return Command{}, err, err == nil
	}
	line := s.Text()
	log.Printf("parsing %q", line)
	// let's be really stupid
	chunks := strings.Fields(line)
	var cmd Command
	if len(chunks) < 1 {
		return cmd, fmt.Errorf("need at least one field"), false
	}
	if len(chunks[0]) != 1 {
		return cmd, fmt.Errorf("first field expected to be one command character, got %q", chunks[0]), false
	}
	cmd.name = chunks[0][0]
	if len(chunks) > 2 {
		return cmd, fmt.Errorf("expected at most one arg"), false
	}
	if len(chunks) == 2 {
		cmd.arg = chunks[1]
	}
	return cmd, nil, false
}
