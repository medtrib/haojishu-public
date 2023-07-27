// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/common/v1/common.proto

package v1

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

// Validate checks the field values on GetCaptchaRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCaptchaRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCaptchaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCaptchaRequestMultiError, or nil if none found.
func (m *GetCaptchaRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCaptchaRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetLastIp()); ip == nil {
		err := GetCaptchaRequestValidationError{
			field:  "LastIp",
			reason: "value must be a valid IP address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCaptchaRequestMultiError(errors)
	}

	return nil
}

// GetCaptchaRequestMultiError is an error wrapping multiple validation errors
// returned by GetCaptchaRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCaptchaRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCaptchaRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCaptchaRequestMultiError) AllErrors() []error { return m }

// GetCaptchaRequestValidationError is the validation error returned by
// GetCaptchaRequest.Validate if the designated constraints aren't met.
type GetCaptchaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCaptchaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCaptchaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCaptchaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCaptchaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCaptchaRequestValidationError) ErrorName() string {
	return "GetCaptchaRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCaptchaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCaptchaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCaptchaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCaptchaRequestValidationError{}

// Validate checks the field values on GetCaptchaReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCaptchaReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCaptchaReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCaptchaReplyMultiError, or nil if none found.
func (m *GetCaptchaReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCaptchaReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CaptchaId

	// no validation rules for Data

	if len(errors) > 0 {
		return GetCaptchaReplyMultiError(errors)
	}

	return nil
}

// GetCaptchaReplyMultiError is an error wrapping multiple validation errors
// returned by GetCaptchaReply.ValidateAll() if the designated constraints
// aren't met.
type GetCaptchaReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCaptchaReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCaptchaReplyMultiError) AllErrors() []error { return m }

// GetCaptchaReplyValidationError is the validation error returned by
// GetCaptchaReply.Validate if the designated constraints aren't met.
type GetCaptchaReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCaptchaReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCaptchaReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCaptchaReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCaptchaReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCaptchaReplyValidationError) ErrorName() string { return "GetCaptchaReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetCaptchaReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCaptchaReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCaptchaReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCaptchaReplyValidationError{}

// Validate checks the field values on VerifyCaptchaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyCaptchaRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyCaptchaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyCaptchaRequestMultiError, or nil if none found.
func (m *VerifyCaptchaRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyCaptchaRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetLastIp()); ip == nil {
		err := VerifyCaptchaRequestValidationError{
			field:  "LastIp",
			reason: "value must be a valid IP address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return VerifyCaptchaRequestMultiError(errors)
	}

	return nil
}

// VerifyCaptchaRequestMultiError is an error wrapping multiple validation
// errors returned by VerifyCaptchaRequest.ValidateAll() if the designated
// constraints aren't met.
type VerifyCaptchaRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyCaptchaRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyCaptchaRequestMultiError) AllErrors() []error { return m }

// VerifyCaptchaRequestValidationError is the validation error returned by
// VerifyCaptchaRequest.Validate if the designated constraints aren't met.
type VerifyCaptchaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyCaptchaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyCaptchaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyCaptchaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyCaptchaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyCaptchaRequestValidationError) ErrorName() string {
	return "VerifyCaptchaRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyCaptchaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyCaptchaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyCaptchaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyCaptchaRequestValidationError{}

// Validate checks the field values on VerifyCaptchaReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyCaptchaReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyCaptchaReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyCaptchaReplyMultiError, or nil if none found.
func (m *VerifyCaptchaReply) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyCaptchaReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return VerifyCaptchaReplyMultiError(errors)
	}

	return nil
}

// VerifyCaptchaReplyMultiError is an error wrapping multiple validation errors
// returned by VerifyCaptchaReply.ValidateAll() if the designated constraints
// aren't met.
type VerifyCaptchaReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyCaptchaReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyCaptchaReplyMultiError) AllErrors() []error { return m }

// VerifyCaptchaReplyValidationError is the validation error returned by
// VerifyCaptchaReply.Validate if the designated constraints aren't met.
type VerifyCaptchaReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyCaptchaReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyCaptchaReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyCaptchaReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyCaptchaReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyCaptchaReplyValidationError) ErrorName() string {
	return "VerifyCaptchaReplyValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyCaptchaReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyCaptchaReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyCaptchaReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyCaptchaReplyValidationError{}

// Validate checks the field values on FireWallRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FireWallRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FireWallRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FireWallRequestMultiError, or nil if none found.
func (m *FireWallRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FireWallRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetLastIp()); ip == nil {
		err := FireWallRequestValidationError{
			field:  "LastIp",
			reason: "value must be a valid IP address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FireWallRequestMultiError(errors)
	}

	return nil
}

// FireWallRequestMultiError is an error wrapping multiple validation errors
// returned by FireWallRequest.ValidateAll() if the designated constraints
// aren't met.
type FireWallRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FireWallRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FireWallRequestMultiError) AllErrors() []error { return m }

// FireWallRequestValidationError is the validation error returned by
// FireWallRequest.Validate if the designated constraints aren't met.
type FireWallRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FireWallRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FireWallRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FireWallRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FireWallRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FireWallRequestValidationError) ErrorName() string { return "FireWallRequestValidationError" }

// Error satisfies the builtin error interface
func (e FireWallRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFireWallRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FireWallRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FireWallRequestValidationError{}

// Validate checks the field values on FireWallReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FireWallReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FireWallReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FireWallReplyMultiError, or
// nil if none found.
func (m *FireWallReply) ValidateAll() error {
	return m.validate(true)
}

func (m *FireWallReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return FireWallReplyMultiError(errors)
	}

	return nil
}

// FireWallReplyMultiError is an error wrapping multiple validation errors
// returned by FireWallReply.ValidateAll() if the designated constraints
// aren't met.
type FireWallReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FireWallReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FireWallReplyMultiError) AllErrors() []error { return m }

// FireWallReplyValidationError is the validation error returned by
// FireWallReply.Validate if the designated constraints aren't met.
type FireWallReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FireWallReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FireWallReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FireWallReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FireWallReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FireWallReplyValidationError) ErrorName() string { return "FireWallReplyValidationError" }

// Error satisfies the builtin error interface
func (e FireWallReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFireWallReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FireWallReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FireWallReplyValidationError{}