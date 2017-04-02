package fit

type decodeOptions struct {
	logger          Logger
	unknownFields   bool
	unknownMessages bool
}

// DecodeOption configures a decoder.
type DecodeOption func(*decodeOptions)

// WithLogger configures the decoder to enable debug logging using the provided
// logger.
func WithLogger(logger Logger) DecodeOption {
	return func(o *decodeOptions) {
		o.logger = logger
	}
}

// WithUnknownFields configures the decoder to record information about unknown
// fields encountered when decoding a known message type. Currently message
// number, field number and number of occurrences are recorded.
func WithUnknownFields() DecodeOption {
	return func(o *decodeOptions) {
		o.unknownFields = true
	}
}

// WithUnknownMessages configures the decoder to record information about unknown
// messages encountered during decoding of a FIT file. Currently message
// number and number of occurrences are recorded.
func WithUnknownMessages() DecodeOption {
	return func(o *decodeOptions) {
		o.unknownMessages = true
	}
}
