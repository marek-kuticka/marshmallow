// Copyright 2022 PerimeterX. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package marshmallow

// Mode dictates the unmarshalling mode.
// Each mode is self documented below.
type Mode uint8

const (
	// ModeFailOnFirstError is the default mode. It makes unmarshalling terminate
	// immediately on any kind of error. This error will then be returned.
	ModeFailOnFirstError Mode = iota

	// ModeAllowMultipleErrors mode makes unmarshalling keep decoding even if
	// errors are encountered. In case of such error, the erroneous value will be omitted from the result.
	// Eventually, all errors will all be returned, alongside the partial result.
	ModeAllowMultipleErrors

	// ModeFailOverToOriginalValue mode makes unmarshalling keep decoding even if
	// errors are encountered. In case of such error, the original external value be placed in the
	// result data, even though it does not meet the schematic requirements.
	// Eventually, all errors will be returned, alongside the full result. Note that the result map
	// will contain values that do not match the struct schema.
	ModeFailOverToOriginalValue
)

// WithMode is an UnmarshalOption function to set the unmarshalling mode.
func WithMode(mode Mode) UnmarshalOption {
	return func(options *unmarshalOptions) {
		options.mode = mode
	}
}

// WithSkipPopulateStruct is an UnmarshalOption function to set the skipPopulateStruct option.
// Skipping populate struct is set to false by default.
// If you do not intend to use the struct value once unmarshalling is finished, set this
// option to true to boost performance. This would mean the struct fields will not be set
// with values, but rather it will only be used as the target schema when populating the result map.
func WithSkipPopulateStruct(skipPopulateStruct bool) UnmarshalOption {
	return func(options *unmarshalOptions) {
		options.skipPopulateStruct = skipPopulateStruct
	}
}

type UnmarshalOption func(*unmarshalOptions)

type unmarshalOptions struct {
	mode               Mode
	skipPopulateStruct bool
}

func buildUnmarshalOptions(options []UnmarshalOption) *unmarshalOptions {
	result := &unmarshalOptions{}
	for _, option := range options {
		option(result)
	}
	return result
}

// JSONDataHandler allow types to handle JSON data as maps.
// Types should implement this interface if they wish to act on the map representation of parsed JSON input.
// This is mainly used to allow nested objects to capture unknown fields and leverage marshmallow's abilities.
type JSONDataHandler interface {
	HandleJSONData(data map[string]interface{})
}
