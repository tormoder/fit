package profile

type Token int

const (
	// Common
	ILLEGAL Token = iota
	EOF
	EMPTY
	PROFILEHDR

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
	"PROFILEHDR",

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
	scan  func() (int, Token, []string)
}

func NewTypeScanner(input [][]string) (*Scanner, error) {
	s := new(Scanner)
	s.input = input
	s.scan = s.tscan
	return s, nil
}

func NewMsgScanner(input [][]string) (*Scanner, error) {
	s := new(Scanner)
	s.input = input
	s.scan = s.mscan
	return s, nil
}

func (s *Scanner) Scan() (line int, token Token, literal []string) {
	return s.scan()
}

func (s *Scanner) read() (int, []string) {
	if s.i > len(s.input)-1 {
		return 0, nil
	}
	line, ch := s.i, s.input[s.i]
	s.i++
	return line, ch
}

func (s *Scanner) tscan() (line int, token Token, literal []string) {
	line, ch := s.read()
	if ch == nil {
		return line, EOF, nil
	}
	if ch[tNAME] != "" {
		if ch[tVALNAME] != "" {
			return line, PROFILEHDR, ch
		}
		return line, THDR, ch
	}
	if ch[tVALNAME] == "" {
		return line, EMPTY, ch
	}
	return line, TFIELD, ch
}

func (s *Scanner) mscan() (line int, token Token, literal []string) {
	line, ch := s.read()

	if ch == nil {
		return line, EOF, nil
	}

	if ch[mMSGNAME] != "" {
		// not empty: CSVHDR, MSGHDR
		if ch[mFDEFN] == "" {
			return line, MSGHDR, ch
		}
		return line, PROFILEHDR, ch
	}

	if ch[mFDEFN] == "" {
		// fdefn empty: can be FMSGHDR, EMPTY, DYNMSGFIELD.
		if ch[mFNAME] == "" {
			// fname empty: FMSGSHDR, EMPTY
			switch {
			case ch[mFTYPE] != "":
				return line, FMSGSHDR, ch
			case isempty(ch[mFTYPE:]):
				return line, EMPTY, ch
			default:
				return line, ILLEGAL, ch
			}
		} else {
			// fname not empty, must be DYNMSGFIELD.
			return line, DYNMSGFIELD, ch
		}
	} else {
		// fdefn not empty: must be MSGFIELD.
		return line, MSGFIELD, ch
	}
}

func isempty(ss []string) bool {
	if len(ss) == 0 {
		return true
	}
	for _, s := range ss {
		if s != "" {
			return false
		}
	}
	return true
}
