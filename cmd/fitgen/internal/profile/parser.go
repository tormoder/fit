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
		tok Token
		lit []string
		n   int
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
	tok, lit := p.scan()
	if tok != PROFILEHDR {
		return fmt.Errorf(
			"got %v, expect %v as first token. Input: %s",
			PROFILEHDR, tok, lit,
		)
	}
	return nil
}

func (p *Parser) ParseType() (*PType, error) {
	if !p.typeParser {
		return nil, errors.New("illegal operation: this is a message parser")
	}
	tok, lit := p.scan()
	if tok == EOF {
		return nil, io.EOF
	}
	if tok != THDR {
		return nil, fmt.Errorf("got %s, expected %s, input %v", tok, MSGHDR, lit)
	}
	t := &PType{}
	t.Header = lit
	for {
		tok, lit := p.scan()
		switch tok {
		case TFIELD:
			t.Fields = append(t.Fields, lit)
		case THDR:
			// Allow this due to bug (?) in the SDK 16.10 workbook.
			// One message is missing the empty row (1166).
			p.unscan()
			return t, nil
		case PROFILEHDR, ILLEGAL:
			return nil, fmt.Errorf("unexpected %s when parsing type, input: %s", tok, lit)
		case EMPTY, EOF:
			if len(t.Fields) == 0 {
				return nil, fmt.Errorf("unexpected %s due to no seen type fields", tok)
			}
			return t, nil
		}
	}
}

func (p *Parser) ParseMsg() (*PMsg, error) {
	if p.typeParser {
		return nil, errors.New("illegal operation: this is a type parser")
	}
	tok, lit := p.scan()
	if tok == EOF || tok == EMPTY {
		return nil, io.EOF
	}
	if tok == FMSGSHDR {
		tok, lit = p.scan()
	}
	if tok != MSGHDR {
		return nil, fmt.Errorf("got %s, expected %s, input %v", tok, MSGHDR, lit)
	}

	m := &PMsg{}
	m.Header = lit
	var lf *PField

	for {
		tok, lit = p.scan()
		switch tok {
		case MSGFIELD:
			lf = &PField{Field: lit}
			m.Fields = append(m.Fields, lf)
		case DYNMSGFIELD:
			if lf == nil {
				return nil, fmt.Errorf("unexpected %s due to no seen field for message yet, input: %s", tok, lit)
			}
			lf.Subfields = append(lf.Subfields, lit)
		case PROFILEHDR, MSGHDR, FMSGSHDR, ILLEGAL:
			return nil, fmt.Errorf("unexpected %s when parsing message, input: %s", tok, lit)
		case EMPTY, EOF:
			if len(m.Header) == 0 {
				return nil, fmt.Errorf("unexpected %s due to no seen message header", tok)
			}
			if len(m.Fields) == 0 {
				return nil, fmt.Errorf("unexpected %s due to no seen message fields", tok)
			}
			return m, nil
		}
	}

}

func (p *Parser) scan() (tok Token, lit []string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}
	tok, lit = p.s.Scan()
	p.buf.tok, p.buf.lit = tok, lit
	return
}

func (p *Parser) unscan() { p.buf.n = 1 }
