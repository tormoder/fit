package profile

import (
	"errors"
	"fmt"
	"io"
)

type PType struct {
	Header []string
	Fields [][]string
}

type PMsg struct {
	Header []string
	Fields []*PField
}

type PField struct {
	Field     []string
	Subfields [][]string
}

type Parser struct {
	s          *Scanner
	typeParser bool
	buf        struct {
		line int
		tok  Token
		lit  []string
		n    int
	}
}

func NewTypeParser(input [][]string) (*Parser, error) {
	s, err := NewTypeScanner(input)
	if err != nil {
		return nil, err
	}
	p := &Parser{s: s, typeParser: true}
	err = p.init()
	return p, err
}

func NewMsgParser(input [][]string) (*Parser, error) {
	s, err := NewMsgScanner(input)
	if err != nil {
		return nil, err
	}
	p := &Parser{s: s}
	err = p.init()
	return p, err
}

func (p *Parser) init() error {
	line, tok, _ := p.scan()
	if tok != PROFILEHDR {
		return unexpectedErr{line: line, token: tok, reason: fmt.Sprintf("expected %v as first token", PROFILEHDR)}
	}
	return nil
}

func (p *Parser) ParseType() (*PType, error) {
	if !p.typeParser {
		return nil, errors.New("illegal operation: this is a message parser")
	}

	line, tok, lit := p.scan()
	for {
		if tok == EOF {
			return nil, io.EOF
		}
		if tok != EMPTY {
			break
		}
		line, tok, lit = p.scan()
	}

	if tok != THDR {
		return nil, unexpectedErr{line: line, token: tok, reason: fmt.Sprintf("expected %v", MSGHDR)}
	}

	t := &PType{}
	t.Header = lit
	for {
		line, tok, lit := p.scan()
		switch tok {
		case TFIELD:
			t.Fields = append(t.Fields, lit)
		case THDR:
			// Allow this due to bug (?) in the SDK 16.10 workbook.
			// One message is missing the empty row (1166).
			p.unscan()
			return t, nil
		case PROFILEHDR, ILLEGAL:
			return nil, unexpectedErr{line: line, token: tok, reason: "parsing type"}
		case EMPTY, EOF:
			return t, nil
		}
	}
}

func (p *Parser) ParseMsg() (*PMsg, error) {
	if p.typeParser {
		return nil, errors.New("illegal operation: this is a type parser")
	}
	line, tok, lit := p.scan()
	if tok == EOF || tok == EMPTY {
		return nil, io.EOF
	}
	if tok == FMSGSHDR {
		line, tok, lit = p.scan()
	}
	if tok != MSGHDR {
		return nil, unexpectedErr{line: line, token: tok, reason: fmt.Sprintf("expected %v", MSGHDR)}
	}

	m := &PMsg{}
	m.Header = lit
	var lf *PField

	for {
		line, tok, lit = p.scan()
		switch tok {
		case MSGFIELD:
			lf = &PField{Field: lit}
			m.Fields = append(m.Fields, lf)
		case DYNMSGFIELD:
			if lf == nil {
				return nil, unexpectedErr{line: line, token: tok, reason: "no seen field for message yet"}
			}
			lf.Subfields = append(lf.Subfields, lit)
		case MSGHDR, FMSGSHDR:
			// No empty row between current and next message/header.
			p.unscan()
			return m, nil
		case PROFILEHDR, ILLEGAL:
			return nil, unexpectedErr{line: line, token: tok, reason: "parsing message"}
		case EMPTY, EOF:
			if len(m.Header) == 0 {
				return nil, unexpectedErr{line: line, token: tok, reason: "no seen message header"}
			}
			if len(m.Fields) == 0 {
				return nil, unexpectedErr{line: line, token: tok, reason: "no seen message fields"}
			}
			return m, nil
		}
	}

}

type unexpectedErr struct {
	line   int
	token  Token
	reason string
}

func (ue unexpectedErr) Error() string {
	return fmt.Sprintf("line %d: unexpected %s - reason: %s", ue.line, ue.token, ue.reason)
}

func (p *Parser) scan() (line int, token Token, literal []string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.line, p.buf.tok, p.buf.lit
	}
	line, token, literal = p.s.Scan()
	p.buf.line, p.buf.tok, p.buf.lit = line, token, literal
	return
}

func (p *Parser) unscan() { p.buf.n = 1 }
