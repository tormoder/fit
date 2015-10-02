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
	if ch[tNAME] != "" {
		if ch[tVALNAME] != "" {
			return CSVHDR, ch
		}
		return THDR, ch
	}
	if ch[tVALNAME] == "" {
		return EMPTY, ch
	}
	return TFIELD, ch
}

func (s *Scanner) mscan() (tok Token, lit []string) {
	ch := s.read()

	if ch == nil {
		return EOF, nil
	}

	if ch[mMSGNAME] != "" {
		// not empty: CSVHDR, MSGHDR
		if ch[mFDEFN] == "" {
			return MSGHDR, ch
		}
		return CSVHDR, ch
	}

	if ch[mFDEFN] == "" {
		// fdefn empty: can be FMSGHDR, EMPTY, DYNMSGFIELD
		if ch[mFNAME] == "" {
			// fname empty: FMSGSHDR, EMPTY
			switch {
			case ch[mFTYPE] != "":
				return FMSGSHDR, ch
			case isempty(ch[mFTYPE:]):
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
