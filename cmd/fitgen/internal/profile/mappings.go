package profile

var typeQuirks = map[string]string{
	"activity": "activity_mode",
}

var timestampTypes = map[string]bool{
	"date_time":       true,
	"local_date_time": true,
}

var baseTypeToGoType = map[string]string{
	"enum":    "byte",
	"sint8":   "int8",
	"uint8":   "uint8",
	"sint16":  "int16",
	"uint16":  "uint16",
	"sint32":  "int32",
	"uint32":  "uint32",
	"string":  "string",
	"float32": "float32",
	"float64": "float64",
	"uint8z":  "uint8",
	"uint16z": "uint16",
	"uint32z": "uint32",
	"byte":    "byte",
}

var baseTypeToInvalidValue = map[string]string{
	"enum":    "0xFF",
	"sint8":   "0x7F",
	"uint8":   "0xFF",
	"sint16":  "0x7FFF",
	"uint16":  "0xFFFF",
	"sint32":  "0x7FFFFFFF",
	"uint32":  "0xFFFFFFFF",
	"string":  "0x00",
	"float32": "0xFFFFFFFF",
	"float64": "0xFFFFFFFFFFFFFFFF",
	"uint8z":  "0x00",
	"uint16z": "0x0000",
	"uint32z": "0x00000000",
	"byte":    "0xFF",
}

var goBaseTypeToInvalidValue = map[string]string{
	"int8":    "0x7F",
	"uint8":   "0xFF",
	"int16":   "0x7FFF",
	"uint16":  "0xFFFF",
	"int32":   "0x7FFFFFFF",
	"uint32":  "0xFFFFFFFF",
	"string":  "\"\"",
	"float32": "0xFFFFFFFF",
	"float64": "0xFFFFFFFFFFFFFFFF",
	"byte":    "0xFF",
}

var knownMesgNumButNoMsg = map[string]bool{
	"Pad":         true,
	"GpsMetadata": true,
}
