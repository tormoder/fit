package profile

import (
	"errors"
	"fmt"
	"io"
)

type Type struct {
	Header []string
	Fields [][]string
}

type Msg struct {
	Header []string
	Fields []*Field
}

type Field struct {
	RegField  []string
	DynFields [][]string
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

func NewTypeParser(r io.Reader) (*Parser, error) {
	s, err := NewTypeScanner(r)
	if err != nil {
		return nil, err
	}
	p := &Parser{s: s, typeParser: true}
	err = p.init()
	return p, nil
}

func NewMsgParser(r io.Reader) (*Parser, error) {
	s, err := NewMsgScanner(r)
	if err != nil {
		return nil, err
	}
	p := &Parser{s: s}
	err = p.init()
	return p, nil
}

func (p *Parser) init() error {
	tok, lit := p.scan()
	if tok != CSVHDR {
		return fmt.Errorf(
			"got %v, expect %v as first token. Input: %s",
			CSVHDR, tok, lit,
		)
	}
	return nil
}

func (p *Parser) ParseType() (*Type, error) {
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
	t := &Type{}
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
		case CSVHDR, ILLEGAL:
			return nil, fmt.Errorf("unexpected %s when parsing type, input: %s", tok, lit)
		case EMPTY, EOF:
			if len(t.Fields) == 0 {
				return nil, fmt.Errorf("unexpected %s due to no seen type fields", tok)
			}
			return t, nil
		}
	}
}

func (p *Parser) ParseMsg() (*Msg, error) {
	if p.typeParser {
		return nil, errors.New("illegal operation: this is a type parser")
	}
	tok, lit := p.scan()
	if tok == EOF {
		return nil, io.EOF
	}
	if tok == FMSGSHDR {
		tok, lit = p.scan()
	}
	if tok != MSGHDR {
		return nil, fmt.Errorf("got %s, expected %s, input %v", tok, MSGHDR, lit)
	}

	m := &Msg{}
	m.Header = lit
	var lf *Field

	for {
		tok, lit = p.scan()
		switch tok {
		case MSGFIELD:
			lf = &Field{RegField: lit}
			m.Fields = append(m.Fields, lf)
		case DYNMSGFIELD:
			if lf == nil {
				return nil, fmt.Errorf("unexpected %s due to no seen field for message yet, input: %s", tok, lit)
			}
			lf.DynFields = append(lf.DynFields, lit)
		case CSVHDR, MSGHDR, FMSGSHDR, ILLEGAL:
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
