package profile

import "github.com/tormoder/fit/internal/types"

type Type struct {
	Name     string
	OrigName string
	BaseType types.Base
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
	Array  string
	Scale  string
	Offset string
	Units  string

	TypeName string
	FType    types.Fit
	Length   string

	Components []Component

	Subfields     []*Field // Not set for subfields
	RefFieldName  []string // Only set for subfields
	RefFieldValue []string // Only set for subfields

	Comment string
	Example string

	data []string
}

func (f *Field) HasOffset() bool {
	return !(f.Offset == "0" || f.Offset == "")
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
