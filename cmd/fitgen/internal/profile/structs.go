package profile

type Type struct {
	Name         string
	CCName       string
	OrigName     string
	PkgName      string
	BaseType     string
	GoBaseType   string
	InvalidValue string
	Values       []ValueTriple

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

	BaseType  string
	BTInvalid string
	GoType    string
	GoInvalid string

	data []string
}

type Component struct {
	Name       string
	Bits       string
	BitsInt    int
	Scale      string
	Offset     string
	Accumulate bool
}
