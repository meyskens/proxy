// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/auth/v3/external_auth.proto

package envoy_service_auth_v3

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

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckRequestMultiError, or
// nil if none found.
func (m *CheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAttributes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckRequestValidationError{
					field:  "Attributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckRequestValidationError{
					field:  "Attributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckRequestValidationError{
				field:  "Attributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CheckRequestMultiError(errors)
	}
	return nil
}

// CheckRequestMultiError is an error wrapping multiple validation errors
// returned by CheckRequest.ValidateAll() if the designated constraints aren't met.
type CheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckRequestMultiError) AllErrors() []error { return m }

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckRequestValidationError) ErrorName() string { return "CheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckRequestValidationError{}

// Validate checks the field values on DeniedHttpResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeniedHttpResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeniedHttpResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeniedHttpResponseMultiError, or nil if none found.
func (m *DeniedHttpResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeniedHttpResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetStatus() == nil {
		err := DeniedHttpResponseValidationError{
			field:  "Status",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeniedHttpResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeniedHttpResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeniedHttpResponseValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeniedHttpResponseValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeniedHttpResponseValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeniedHttpResponseValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Body

	if len(errors) > 0 {
		return DeniedHttpResponseMultiError(errors)
	}
	return nil
}

// DeniedHttpResponseMultiError is an error wrapping multiple validation errors
// returned by DeniedHttpResponse.ValidateAll() if the designated constraints
// aren't met.
type DeniedHttpResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeniedHttpResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeniedHttpResponseMultiError) AllErrors() []error { return m }

// DeniedHttpResponseValidationError is the validation error returned by
// DeniedHttpResponse.Validate if the designated constraints aren't met.
type DeniedHttpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeniedHttpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeniedHttpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeniedHttpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeniedHttpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeniedHttpResponseValidationError) ErrorName() string {
	return "DeniedHttpResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeniedHttpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeniedHttpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeniedHttpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeniedHttpResponseValidationError{}

// Validate checks the field values on OkHttpResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OkHttpResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OkHttpResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OkHttpResponseMultiError,
// or nil if none found.
func (m *OkHttpResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OkHttpResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OkHttpResponseValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OkHttpResponseValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OkHttpResponseValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetDynamicMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OkHttpResponseValidationError{
					field:  "DynamicMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OkHttpResponseValidationError{
					field:  "DynamicMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDynamicMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OkHttpResponseValidationError{
				field:  "DynamicMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetResponseHeadersToAdd() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OkHttpResponseValidationError{
						field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OkHttpResponseValidationError{
						field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OkHttpResponseValidationError{
					field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OkHttpResponseMultiError(errors)
	}
	return nil
}

// OkHttpResponseMultiError is an error wrapping multiple validation errors
// returned by OkHttpResponse.ValidateAll() if the designated constraints
// aren't met.
type OkHttpResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OkHttpResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OkHttpResponseMultiError) AllErrors() []error { return m }

// OkHttpResponseValidationError is the validation error returned by
// OkHttpResponse.Validate if the designated constraints aren't met.
type OkHttpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OkHttpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OkHttpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OkHttpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OkHttpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OkHttpResponseValidationError) ErrorName() string { return "OkHttpResponseValidationError" }

// Error satisfies the builtin error interface
func (e OkHttpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOkHttpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OkHttpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OkHttpResponseValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckResponseMultiError, or
// nil if none found.
func (m *CheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResponseValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDynamicMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "DynamicMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckResponseValidationError{
					field:  "DynamicMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDynamicMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResponseValidationError{
				field:  "DynamicMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.HttpResponse.(type) {

	case *CheckResponse_DeniedResponse:

		if all {
			switch v := interface{}(m.GetDeniedResponse()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  "DeniedResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  "DeniedResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeniedResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					field:  "DeniedResponse",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CheckResponse_OkResponse:

		if all {
			switch v := interface{}(m.GetOkResponse()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  "OkResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckResponseValidationError{
						field:  "OkResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOkResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					field:  "OkResponse",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CheckResponseMultiError(errors)
	}
	return nil
}

// CheckResponseMultiError is an error wrapping multiple validation errors
// returned by CheckResponse.ValidateAll() if the designated constraints
// aren't met.
type CheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckResponseMultiError) AllErrors() []error { return m }

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResponseValidationError) ErrorName() string { return "CheckResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResponseValidationError{}
