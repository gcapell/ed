package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
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

	return parseLine(line), false
}

var cmdRegexp = regexp.MustCompile("^ *([0-9]+)? *,? *?([0-9]+)? *([a-zA-Z]) *([^ ]+)?$")

// parseLine matches a string as "[addr1[,addr2]] command" (with optional whitespace)
func parseLine(s string) Command {
	m := cmdRegexp.FindStringSubmatch(s)
	log.Printf("parse(%s)->%q", s, m)
	if len(m) > 0 {
		return Command{
			name:  m[3][0],
			addr1: m[1],
			addr2: m[2],
			arg:   m[4],
		}
	}
	return Command{}

}
func errorf(format string, args ...interface{}) (Command, bool) {
	return Command{parseError: fmt.Sprintf(format, args)}, false
}
