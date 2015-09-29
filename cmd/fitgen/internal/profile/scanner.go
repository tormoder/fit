package profile

import (
	"encoding/csv"
	"io"
)

type Token int

const (
	// Common
	ILLEGAL Token = iota
	EOF
	EMPTY
	CSVHDR

	// Types
	THDR
	TFIELD

	// Messages
	FMSGSHDR
	MSGHDR
	MSGFIELD
	DYNMSGFIELD
)

var tokenString = [...]string{
	"ILLEGAL",
	"EOF",
	"EMPTY",
	"CSVHDR",

	"THDR",
	"TFIELD",

	"FMSGSHDR",
	"MSGHDR",
	"MSGFIELD",
	"DYNMSGFIELD",
}

func (t Token) String() string {
	return tokenString[t]
}

type Scanner struct {
	i     int
	input [][]string
	scan  func() (Token, []string)
}

func NewTypeScanner(r io.Reader) (*Scanner, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 5
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	s := new(Scanner)
	s.input = data
	s.scan = s.tscan
	return s, nil
}

func NewMsgScanner(r io.Reader) (*Scanner, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 16
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	s := new(Scanner)
	s.input = data
	s.scan = s.mscan
	return s, nil
}

func (s *Scanner) Scan() (tok Token, lit []string) {
	return s.scan()
}

func (s *Scanner) read() []string {
	if s.i > len(s.input)-1 {
		return nil
	}
	ch := s.input[s.i]
	s.i++
	return ch
}

func (s *Scanner) tscan() (tok Token, lit []string) {
	ch := s.read()
	if ch == nil {
		return EOF, nil
	}
	if ch[tname] != "" {
		if ch[tvalname] != "" {
			return CSVHDR, ch
		}
		return THDR, ch
	}
	if ch[tvalname] == "" {
		return EMPTY, ch
	}
	return TFIELD, ch
}

func (s *Scanner) mscan() (tok Token, lit []string) {
	ch := s.read()

	if ch == nil {
		return EOF, nil
	}

	if ch[mmsgname] != "" {
		// not empty: CSVHDR, MSGHDR
		if ch[mfdefn] == "" {
			return MSGHDR, ch
		}
		return CSVHDR, ch
	}

	if ch[mfdefn] == "" {
		// fdefn empty: can be FMSGHDR, EMPTY, DYNMSGFIELD
		if ch[mfname] == "" {
			// fname empty: FMSGSHDR, EMPTY
			switch {
			case ch[mftype] != "":
				return FMSGSHDR, ch
			case isempty(ch[mftype:]):
				return EMPTY, ch
			default:
				return ILLEGAL, ch
			}
		} else {
			// fname not empty, must be DYNMSGFIELD
			return DYNMSGFIELD, ch
		}
	} else {
		// fdefn not empty: must be MSGFIELD
		return MSGFIELD, ch
	}
}

func isempty(ss []string) bool {
	if ss == nil || len(ss) == 0 {
		return true
	}
	for _, s := range ss {
		if s != "" {
			return false
		}
	}
	return true
}
