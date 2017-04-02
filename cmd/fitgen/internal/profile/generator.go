package profile

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Profile struct {
	TypesSource            []byte
	MessagesSource         []byte
	ProfileSource          []byte
	StringerInput          string
	MesgNumsWithoutMessage []string
}

type generatorOptions struct {
	genTimestamp bool
	sdkVersion   string
	logger       *log.Logger
}

type GeneratorOption func(*generatorOptions)

func WithGenerationTimestamp(gt bool) GeneratorOption {
	return func(o *generatorOptions) {
		o.genTimestamp = gt
	}
}

func WithSDKVersionOverride(version string) GeneratorOption {
	return func(o *generatorOptions) {
		o.sdkVersion = version
	}
}

func WithLogger(logger *log.Logger) GeneratorOption {
	return func(o *generatorOptions) {
		o.logger = logger
	}
}

type Generator struct {
	opts generatorOptions

	typesData [][]string
	msgsData  [][]string

	types map[string]*Type
	msgs  []*Msg

	p *Profile
}

func NewGenerator(inputPath string, workbookData []byte, opts ...GeneratorOption) (*Generator, error) {
	g := new(Generator)
	g.p = new(Profile)

	for _, opt := range opts {
		opt(&g.opts)
	}

	if g.opts.sdkVersion == "" {
		switch filepath.Ext(inputPath) {
		case ".zip":
			g.opts.sdkVersion = parseSDKVersionFromZipFilePath(inputPath)
		default:
			g.opts.sdkVersion = "Unknown"
		}
	}

	g.logln("sdk version:", g.opts.sdkVersion)
	g.logln("parsing workbook")
	var err error
	g.typesData, g.msgsData, err = parseWorkbook(workbookData)
	if err != nil {
		return nil, fmt.Errorf("error creating generator: %v", err)
	}

	return g, nil
}

func (g *Generator) GenerateProfile() (*Profile, error) {
	if err := g.parseTypes(); err != nil {
		return nil, fmt.Errorf("error parsing types: %v", err)
	}
	if err := g.parseMsgs(); err != nil {
		return nil, fmt.Errorf("error parsing msgs: %v", err)
	}
	if err := g.genCode(); err != nil {
		return nil, fmt.Errorf("code generation error: %v", err)
	}
	if err := g.genStringerTypeInput(); err != nil {
		return nil, fmt.Errorf("error generating stringer input: %v", err)
	}
	if err := g.genMsgNumsVsMsgs(); err != nil {
		return nil, fmt.Errorf("error comparing msgnums vs messages: %v", err)
	}
	return g.p, nil
}

func (g *Generator) parseTypes() error {
	g.logln("parsing types")

	parser, err := NewTypeParser(g.typesData)
	if err != nil {
		return fmt.Errorf("error creating parser: %v", err)
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
		return fmt.Errorf("error transforming types: %v", err)
	}

	return nil
}

func (g *Generator) parseMsgs() error {
	g.logln("parsing messages")

	parser, err := NewMsgParser(g.msgsData)
	if err != nil {
		return fmt.Errorf("parser error: %v", err)
	}

	var pmsgs []*PMsg
	for {
		m, perr := parser.ParseMsg()
		if perr == io.EOF {
			break
		}
		if perr != nil {
			return fmt.Errorf("parsing error: %v", perr)
		}
		pmsgs = append(pmsgs, m)
	}

	g.msgs, err = TransformMsgs(pmsgs, g.types)
	if err != nil {
		return fmt.Errorf("transform error: %v", err)
	}

	return nil
}

func (g *Generator) genCode() error {
	g.logln("generating code")

	var err error
	codeg := newCodeGenerator(g.opts.sdkVersion, g.opts.genTimestamp)
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

	var allTypesBuf bytes.Buffer
	for _, tkey := range tkeys {
		t := g.types[tkey]
		allTypesBuf.WriteString(t.Name)
		allTypesBuf.WriteByte(',')
	}

	allTypes := allTypesBuf.Bytes()
	g.p.StringerInput = string(allTypes[:len(allTypes)-1]) // last comma

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

func (g *Generator) logln(v ...interface{}) {
	if g.opts.logger != nil {
		g.opts.logger.Println(v...)
	}
}

var debugfg, _ = strconv.ParseBool(os.Getenv("FITGEN_DEBUG"))

func debugln(v ...interface{}) {
	if debugfg {
		fmt.Println(v...)
	}
}

func parseSDKVersionFromZipFilePath(path string) string {
	// Brittle.
	// TODO: Maybe parse 'c/fit.h' with regexp instead.
	_, file := filepath.Split(path)
	ver := strings.TrimSuffix(file, ".zip")
	return strings.TrimPrefix(ver, "FitSDKRelease_")
}
