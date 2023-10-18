// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: example.proto

package example

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

// Validate checks the field values on HelloRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloRequestMultiError, or
// nil if none found.
func (m *HelloRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return HelloRequestMultiError(errors)
	}

	return nil
}

// HelloRequestMultiError is an error wrapping multiple validation errors
// returned by HelloRequest.ValidateAll() if the designated constraints aren't met.
type HelloRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloRequestMultiError) AllErrors() []error { return m }

// HelloRequestValidationError is the validation error returned by
// HelloRequest.Validate if the designated constraints aren't met.
type HelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloRequestValidationError) ErrorName() string { return "HelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e HelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloRequestValidationError{}

// Validate checks the field values on HelloResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelloResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelloResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelloResponseMultiError, or
// nil if none found.
func (m *HelloResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *HelloResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return HelloResponseMultiError(errors)
	}

	return nil
}

// HelloResponseMultiError is an error wrapping multiple validation errors
// returned by HelloResponse.ValidateAll() if the designated constraints
// aren't met.
type HelloResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloResponseMultiError) AllErrors() []error { return m }

// HelloResponseValidationError is the validation error returned by
// HelloResponse.Validate if the designated constraints aren't met.
type HelloResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloResponseValidationError) ErrorName() string { return "HelloResponseValidationError" }

// Error satisfies the builtin error interface
func (e HelloResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloResponseValidationError{}

// Validate checks the field values on SayRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SayRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SayRequestMultiError, or
// nil if none found.
func (m *SayRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SayRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return SayRequestMultiError(errors)
	}

	return nil
}

// SayRequestMultiError is an error wrapping multiple validation errors
// returned by SayRequest.ValidateAll() if the designated constraints aren't met.
type SayRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayRequestMultiError) AllErrors() []error { return m }

// SayRequestValidationError is the validation error returned by
// SayRequest.Validate if the designated constraints aren't met.
type SayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayRequestValidationError) ErrorName() string { return "SayRequestValidationError" }

// Error satisfies the builtin error interface
func (e SayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayRequestValidationError{}

// Validate checks the field values on SayResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SayResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SayResponseMultiError, or
// nil if none found.
func (m *SayResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SayResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return SayResponseMultiError(errors)
	}

	return nil
}

// SayResponseMultiError is an error wrapping multiple validation errors
// returned by SayResponse.ValidateAll() if the designated constraints aren't met.
type SayResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayResponseMultiError) AllErrors() []error { return m }

// SayResponseValidationError is the validation error returned by
// SayResponse.Validate if the designated constraints aren't met.
type SayResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayResponseValidationError) ErrorName() string { return "SayResponseValidationError" }

// Error satisfies the builtin error interface
func (e SayResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayResponseValidationError{}
