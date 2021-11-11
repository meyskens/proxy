// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/compressor/v3/compressor.proto

package envoy_extensions_filters_http_compressor_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Compressor with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Compressor) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Compressor with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CompressorMultiError, or
// nil if none found.
func (m *Compressor) ValidateAll() error {
	return m.validate(true)
}

func (m *Compressor) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetContentLength()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "ContentLength",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "ContentLength",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContentLength()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompressorValidationError{
				field:  "ContentLength",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DisableOnEtagHeader

	// no validation rules for RemoveAcceptEncodingHeader

	if all {
		switch v := interface{}(m.GetRuntimeEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "RuntimeEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "RuntimeEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRuntimeEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompressorValidationError{
				field:  "RuntimeEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCompressorLibrary()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "CompressorLibrary",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "CompressorLibrary",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompressorLibrary()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompressorValidationError{
				field:  "CompressorLibrary",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRequestDirectionConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "RequestDirectionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "RequestDirectionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequestDirectionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompressorValidationError{
				field:  "RequestDirectionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResponseDirectionConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "ResponseDirectionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompressorValidationError{
					field:  "ResponseDirectionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponseDirectionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompressorValidationError{
				field:  "ResponseDirectionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CompressorMultiError(errors)
	}
	return nil
}

// CompressorMultiError is an error wrapping multiple validation errors
// returned by Compressor.ValidateAll() if the designated constraints aren't met.
type CompressorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompressorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompressorMultiError) AllErrors() []error { return m }

// CompressorValidationError is the validation error returned by
// Compressor.Validate if the designated constraints aren't met.
type CompressorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompressorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompressorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompressorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompressorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompressorValidationError) ErrorName() string { return "CompressorValidationError" }

// Error satisfies the builtin error interface
func (e CompressorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompressor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompressorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompressorValidationError{}

// Validate checks the field values on Compressor_CommonDirectionConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *Compressor_CommonDirectionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Compressor_CommonDirectionConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// Compressor_CommonDirectionConfigMultiError, or nil if none found.
func (m *Compressor_CommonDirectionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *Compressor_CommonDirectionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Compressor_CommonDirectionConfigValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Compressor_CommonDirectionConfigValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Compressor_CommonDirectionConfigValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMinContentLength()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Compressor_CommonDirectionConfigValidationError{
					field:  "MinContentLength",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Compressor_CommonDirectionConfigValidationError{
					field:  "MinContentLength",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMinContentLength()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Compressor_CommonDirectionConfigValidationError{
				field:  "MinContentLength",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Compressor_CommonDirectionConfigMultiError(errors)
	}
	return nil
}

// Compressor_CommonDirectionConfigMultiError is an error wrapping multiple
// validation errors returned by
// Compressor_CommonDirectionConfig.ValidateAll() if the designated
// constraints aren't met.
type Compressor_CommonDirectionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Compressor_CommonDirectionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Compressor_CommonDirectionConfigMultiError) AllErrors() []error { return m }

// Compressor_CommonDirectionConfigValidationError is the validation error
// returned by Compressor_CommonDirectionConfig.Validate if the designated
// constraints aren't met.
type Compressor_CommonDirectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Compressor_CommonDirectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Compressor_CommonDirectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Compressor_CommonDirectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Compressor_CommonDirectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Compressor_CommonDirectionConfigValidationError) ErrorName() string {
	return "Compressor_CommonDirectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Compressor_CommonDirectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompressor_CommonDirectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Compressor_CommonDirectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Compressor_CommonDirectionConfigValidationError{}

// Validate checks the field values on Compressor_RequestDirectionConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *Compressor_RequestDirectionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Compressor_RequestDirectionConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// Compressor_RequestDirectionConfigMultiError, or nil if none found.
func (m *Compressor_RequestDirectionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *Compressor_RequestDirectionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Compressor_RequestDirectionConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Compressor_RequestDirectionConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Compressor_RequestDirectionConfigValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Compressor_RequestDirectionConfigMultiError(errors)
	}
	return nil
}

// Compressor_RequestDirectionConfigMultiError is an error wrapping multiple
// validation errors returned by
// Compressor_RequestDirectionConfig.ValidateAll() if the designated
// constraints aren't met.
type Compressor_RequestDirectionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Compressor_RequestDirectionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Compressor_RequestDirectionConfigMultiError) AllErrors() []error { return m }

// Compressor_RequestDirectionConfigValidationError is the validation error
// returned by Compressor_RequestDirectionConfig.Validate if the designated
// constraints aren't met.
type Compressor_RequestDirectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Compressor_RequestDirectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Compressor_RequestDirectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Compressor_RequestDirectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Compressor_RequestDirectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Compressor_RequestDirectionConfigValidationError) ErrorName() string {
	return "Compressor_RequestDirectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Compressor_RequestDirectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompressor_RequestDirectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Compressor_RequestDirectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Compressor_RequestDirectionConfigValidationError{}

// Validate checks the field values on Compressor_ResponseDirectionConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *Compressor_ResponseDirectionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Compressor_ResponseDirectionConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// Compressor_ResponseDirectionConfigMultiError, or nil if none found.
func (m *Compressor_ResponseDirectionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *Compressor_ResponseDirectionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Compressor_ResponseDirectionConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Compressor_ResponseDirectionConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Compressor_ResponseDirectionConfigValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DisableOnEtagHeader

	// no validation rules for RemoveAcceptEncodingHeader

	if len(errors) > 0 {
		return Compressor_ResponseDirectionConfigMultiError(errors)
	}
	return nil
}

// Compressor_ResponseDirectionConfigMultiError is an error wrapping multiple
// validation errors returned by
// Compressor_ResponseDirectionConfig.ValidateAll() if the designated
// constraints aren't met.
type Compressor_ResponseDirectionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Compressor_ResponseDirectionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Compressor_ResponseDirectionConfigMultiError) AllErrors() []error { return m }

// Compressor_ResponseDirectionConfigValidationError is the validation error
// returned by Compressor_ResponseDirectionConfig.Validate if the designated
// constraints aren't met.
type Compressor_ResponseDirectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Compressor_ResponseDirectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Compressor_ResponseDirectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Compressor_ResponseDirectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Compressor_ResponseDirectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Compressor_ResponseDirectionConfigValidationError) ErrorName() string {
	return "Compressor_ResponseDirectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Compressor_ResponseDirectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompressor_ResponseDirectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Compressor_ResponseDirectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Compressor_ResponseDirectionConfigValidationError{}
