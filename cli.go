package cli

import "io"

type Command struct {
	name                byte   // e -> edit, ...
	addr1, addr2, addr3 string // optional address
	arg                 string // e.g. text to insert.  May need to expand on this for complex commands
}

func Parse(r io.Reader, ch chan Command) {
	ch <- Command{name: 'e', arg: "banana"}
	close(ch)
}
