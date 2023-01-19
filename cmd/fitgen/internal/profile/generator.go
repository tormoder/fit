package profile

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type Profile struct {
	TypesSource            []byte
	MessagesSource         []byte
	ProfileSource          []byte
	StringerInput          []string
	MesgNumsWithoutMessage []string
}

type generatorOptions struct {
	genTimestamp    bool
	logger          *log.Logger
	debug           bool
	handleHRSTQuirk bool
}

type GeneratorOption func(*generatorOptions)

func WithGenerationTimestamp(gt bool) GeneratorOption {
	return func(o *generatorOptions) {
		o.genTimestamp = gt
	}
}

func WithLogger(logger *log.Logger) GeneratorOption {
	return func(o *generatorOptions) {
		o.logger = logger
	}
}

func WithDebugOutput() GeneratorOption {
	return func(o *generatorOptions) {
		o.debug = true
	}
}

func WithHandleHeartRateSourceTypeQuirk() GeneratorOption {
	return func(o *generatorOptions) {
		o.handleHRSTQuirk = true
	}
}

type Generator struct {
	opts                 generatorOptions
	sdkMajVer, sdkMinVer int

	typesData [][]string
	msgsData  [][]string

	types map[string]*Type
	msgs  []*Msg

	p *Profile
}

func NewGenerator(sdkMajVer, sdkMinVer int, workbookData []byte, opts ...GeneratorOption) (*Generator, error) {
	g := &Generator{
		sdkMajVer: sdkMajVer,
		sdkMinVer: sdkMinVer,
		p:         new(Profile),
	}

	for _, opt := range opts {
		opt(&g.opts)
	}

	// The code generation is not performance critical,
	// so we can avoid nil-checks when logging.
	if g.opts.logger == nil {
		g.opts.logger = log.New(io.Discard, "", 0)
	}

	g.logf("sdk version: %d.%d", sdkMajVer, sdkMinVer)
	g.logln("parsing workbook")

	var err error
	g.typesData, g.msgsData, err = parseWorkbook(workbookData)
	if err != nil {
		return nil, fmt.Errorf("error creating generator: %w", err)
	}

	return g, nil
}

func (g *Generator) GenerateProfile() (*Profile, error) {
	if err := g.parseTypes(); err != nil {
		return nil, fmt.Errorf("error parsing types: %w", err)
	}
	if err := g.parseMsgs(); err != nil {
		return nil, fmt.Errorf("error parsing msgs: %w", err)
	}
	if err := g.genCode(); err != nil {
		return nil, fmt.Errorf("code generation error: %w", err)
	}
	if err := g.genStringerTypeInput(); err != nil {
		return nil, fmt.Errorf("error generating stringer input: %w", err)
	}
	if err := g.genMsgNumsVsMsgs(); err != nil {
		return nil, fmt.Errorf("error comparing msgnums vs messages: %w", err)
	}
	return g.p, nil
}

func (g *Generator) parseTypes() error {
	g.logln("parsing types")

	parser, err := NewTypeParser(g.typesData)
	if err != nil {
		return fmt.Errorf("error creating parser: %w", err)
	}

	var ptypes []*PType
	for {
		t, perr := parser.ParseType()
		if perr == io.EOF {
			break
		}
		if perr != nil {
			return perr
		}
		ptypes = append(ptypes, t)
	}

	g.types, err = TransformTypes(ptypes)
	if err != nil {
		return fmt.Errorf("error transforming types: %w", err)
	}

	return nil
}

func (g *Generator) parseMsgs() error {
	g.logln("parsing messages")

	parser, err := NewMsgParser(g.msgsData)
	if err != nil {
		return fmt.Errorf("parser error: %w", err)
	}

	var pmsgs []*PMsg
	for {
		m, perr := parser.ParseMsg()
		if perr == io.EOF {
			break
		}
		if perr != nil {
			return fmt.Errorf("parsing error: %w", perr)
		}
		pmsgs = append(pmsgs, m)
	}

	g.msgs, err = TransformMsgs(pmsgs, g.types, g.opts.handleHRSTQuirk, g.opts.logger)
	if err != nil {
		return fmt.Errorf("transform error: %w", err)
	}

	return nil
}

func (g *Generator) genCode() error {
	g.logln("generating code")

	var err error
	codeg := newCodeGenerator(g.sdkMajVer, g.sdkMinVer, g.opts.genTimestamp, g.opts.logger)
	g.p.TypesSource, err = codeg.generateTypes(g.types)
	if err != nil {
		return err
	}
	g.p.MessagesSource, err = codeg.generateMsgs(g.msgs)
	if err != nil {
		return err
	}
	g.p.ProfileSource, err = codeg.generateProfile(g.types, g.msgs)
	return err
}

func (g *Generator) genStringerTypeInput() error {
	g.logln("generating stringer input")

	tkeys := make([]string, 0, len(g.types))
	for tkey := range g.types {
		tkeys = append(tkeys, tkey)
	}
	sort.Strings(tkeys)

	allTypes := make([]string, 0, len(tkeys))
	for _, tkey := range tkeys {
		t := g.types[tkey]
		allTypes = append(allTypes, t.Name)
	}

	g.p.StringerInput = allTypes

	return nil
}

func (g *Generator) genMsgNumsVsMsgs() error {
	g.logln("generating messages nums vs messages")

	mesgNum, found := g.types["MesgNum"]
	if !found {
		return errors.New("MesgNum type not found")
	}

	nMesgNum := len(mesgNum.Values) - 2 // skip range min/max
	if nMesgNum-len(g.msgs) == 0 {
		return nil
	}

	msgsMap := make(map[string]*Msg)
	for _, msg := range g.msgs {
		msgsMap[msg.CCName] = msg
	}

	var diff []string
	for _, mnv := range mesgNum.Values {
		if strings.HasPrefix(mnv.Name, "MfgRange") {
			continue
		}
		_, ok := msgsMap[mnv.Name]
		if !ok {
			diff = append(diff, mnv.Name)
		}
	}
	g.p.MesgNumsWithoutMessage = diff

	return nil
}

func (g *Generator) logf(format string, v ...interface{}) {
	if g.opts.logger != nil {
		g.opts.logger.Printf(format, v...)
	}
}

func (g *Generator) logln(v ...interface{}) {
	if g.opts.logger != nil {
		g.opts.logger.Println(v...)
	}
}
