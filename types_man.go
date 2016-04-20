package fit

import "fmt"

// Manually generated types: The 'Bool' type is not found in the SDK
// specification, so it won't be auto-generated, but it is also not a base
// type.

type Bool byte

const (
	BoolFalse   Bool = 0
	BoolTrue    Bool = 1
	BoolInvalid Bool = 255
)
const (
	_Bool_name_0 = "FalseTrue"
	_Bool_name_1 = "Invalid"
)

var (
	_Bool_index_0 = [...]uint8{0, 5, 9}
)

func (i Bool) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _Bool_name_0[_Bool_index_0[i]:_Bool_index_0[i+1]]
	case i == 255:
		return _Bool_name_1
	default:
		return fmt.Sprintf("Bool(%d)", i)
	}
}
