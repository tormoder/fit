package profile

import "github.com/tormoder/fit/internal/base"

type Type struct {
	Name     string
	CCName   string
	OrigName string
	PkgName  string
	BaseType base.Type
	Values   []ValueTriple

	data *PType
}

type ValueTriple struct {
	Name, Value, Comment string
}

type Msg struct {
	Name        string
	CCName      string
	Fields      []*Field
	FieldByName map[string]*Field
}

type Field struct {
	DefNum string
	Name   string
	CCName string
	Type   string
	Array  string // 255=N
	Scale  string
	Offset string
	Units  string

	Components []Component

	Subfields     []*Field // Not set for subfields
	RefFieldName  []string // Only set for subfields
	RefFieldValue []string // Only set for subfields

	Comment string
	Example string

	BaseType  base.Type
	BTInvalid string
	GoType    string
	GoInvalid string

	data []string
}

func (f *Field) HasOffset() bool {
	return !(f.Offset == "0" || f.Offset == "")
}

func (f *Field) IsArray() bool {
	return f.Array != "0"
}

func (f *Field) HasComment() bool {
	return f.Comment != ""
}

type Component struct {
	Name       string
	Bits       string
	BitsInt    int
	Scale      string
	Offset     string
	Accumulate bool
}
