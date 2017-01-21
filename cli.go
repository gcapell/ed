package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	return Command{name: 'e', arg: "banana"}, nil, false
}
