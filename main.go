package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	var b Buffer
	ch := make(chan Command)
	go Parse(os.Stdin, ch)
	for cmd := range ch {
		if len(cmd.parseError) > 0 {
			fmt.Fprintf(os.Stderr, "problem: %s", cmd.parseError)
			continue
		}
		if err := cmd.exec(&b); err != nil {
			fmt.Fprintf(os.Stderr, "problem:%s executing %s", err, cmd)
		}
	}
}

func (c Command) exec(b *Buffer) error {
	switch c.name {
	case 'e':
		return c.edit(b)
	case 'p':
		return c.print(b)
	}
	return fmt.Errorf("unrecognised command: %s", string(c.name))
}

func (c Command) edit(b *Buffer) error {
	filename := c.arg // FIXME
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	b.edit(filename, data)
	return nil
}

func (c Command) print(b *Buffer) error {
	a1, err := strconv.Atoi(c.addr1)
	if err != nil {
		return err
	}
	a2, err := strconv.Atoi(c.addr2)
	if err != nil {
		return err
	}
	return b.print(a1, a2)
}
