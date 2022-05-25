// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/starttls/v3/starttls.proto

package starttlsv3

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

// Validate checks the field values on StartTlsConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StartTlsConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartTlsConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StartTlsConfigMultiError,
// or nil if none found.
func (m *StartTlsConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *StartTlsConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCleartextSocketConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StartTlsConfigValidationError{
					field:  "CleartextSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StartTlsConfigValidationError{
					field:  "CleartextSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCleartextSocketConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StartTlsConfigValidationError{
				field:  "CleartextSocketConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTlsSocketConfig() == nil {
		err := StartTlsConfigValidationError{
			field:  "TlsSocketConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTlsSocketConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StartTlsConfigValidationError{
					field:  "TlsSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StartTlsConfigValidationError{
					field:  "TlsSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTlsSocketConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StartTlsConfigValidationError{
				field:  "TlsSocketConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StartTlsConfigMultiError(errors)
	}
	return nil
}

// StartTlsConfigMultiError is an error wrapping multiple validation errors
// returned by StartTlsConfig.ValidateAll() if the designated constraints
// aren't met.
type StartTlsConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartTlsConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartTlsConfigMultiError) AllErrors() []error { return m }

// StartTlsConfigValidationError is the validation error returned by
// StartTlsConfig.Validate if the designated constraints aren't met.
type StartTlsConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartTlsConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartTlsConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartTlsConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartTlsConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartTlsConfigValidationError) ErrorName() string { return "StartTlsConfigValidationError" }

// Error satisfies the builtin error interface
func (e StartTlsConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartTlsConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartTlsConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartTlsConfigValidationError{}

// Validate checks the field values on UpstreamStartTlsConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpstreamStartTlsConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpstreamStartTlsConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpstreamStartTlsConfigMultiError, or nil if none found.
func (m *UpstreamStartTlsConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *UpstreamStartTlsConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCleartextSocketConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamStartTlsConfigValidationError{
					field:  "CleartextSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamStartTlsConfigValidationError{
					field:  "CleartextSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCleartextSocketConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamStartTlsConfigValidationError{
				field:  "CleartextSocketConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTlsSocketConfig() == nil {
		err := UpstreamStartTlsConfigValidationError{
			field:  "TlsSocketConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTlsSocketConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamStartTlsConfigValidationError{
					field:  "TlsSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamStartTlsConfigValidationError{
					field:  "TlsSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTlsSocketConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamStartTlsConfigValidationError{
				field:  "TlsSocketConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpstreamStartTlsConfigMultiError(errors)
	}
	return nil
}

// UpstreamStartTlsConfigMultiError is an error wrapping multiple validation
// errors returned by UpstreamStartTlsConfig.ValidateAll() if the designated
// constraints aren't met.
type UpstreamStartTlsConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpstreamStartTlsConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpstreamStartTlsConfigMultiError) AllErrors() []error { return m }

// UpstreamStartTlsConfigValidationError is the validation error returned by
// UpstreamStartTlsConfig.Validate if the designated constraints aren't met.
type UpstreamStartTlsConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamStartTlsConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamStartTlsConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamStartTlsConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamStartTlsConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamStartTlsConfigValidationError) ErrorName() string {
	return "UpstreamStartTlsConfigValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamStartTlsConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamStartTlsConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamStartTlsConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamStartTlsConfigValidationError{}