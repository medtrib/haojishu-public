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

// Validate checks the field values on GetCaptchaReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCaptchaReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCaptchaReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCaptchaReqMultiError, or
// nil if none found.
func (m *GetCaptchaReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCaptchaReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetLastIp()); ip == nil {
		err := GetCaptchaReqValidationError{
			field:  "LastIp",
			reason: "value must be a valid IP address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCaptchaReqMultiError(errors)
	}

	return nil
}

// GetCaptchaReqMultiError is an error wrapping multiple validation errors
// returned by GetCaptchaReq.ValidateAll() if the designated constraints
// aren't met.
type GetCaptchaReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCaptchaReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCaptchaReqMultiError) AllErrors() []error { return m }

// GetCaptchaReqValidationError is the validation error returned by
// GetCaptchaReq.Validate if the designated constraints aren't met.
type GetCaptchaReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCaptchaReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCaptchaReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCaptchaReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCaptchaReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCaptchaReqValidationError) ErrorName() string { return "GetCaptchaReqValidationError" }

// Error satisfies the builtin error interface
func (e GetCaptchaReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCaptchaReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCaptchaReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCaptchaReqValidationError{}

// Validate checks the field values on GetCaptchaRep with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCaptchaRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCaptchaRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCaptchaRepMultiError, or
// nil if none found.
func (m *GetCaptchaRep) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCaptchaRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CaptchaId

	// no validation rules for Data

	if len(errors) > 0 {
		return GetCaptchaRepMultiError(errors)
	}

	return nil
}

// GetCaptchaRepMultiError is an error wrapping multiple validation errors
// returned by GetCaptchaRep.ValidateAll() if the designated constraints
// aren't met.
type GetCaptchaRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCaptchaRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCaptchaRepMultiError) AllErrors() []error { return m }

// GetCaptchaRepValidationError is the validation error returned by
// GetCaptchaRep.Validate if the designated constraints aren't met.
type GetCaptchaRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCaptchaRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCaptchaRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCaptchaRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCaptchaRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCaptchaRepValidationError) ErrorName() string { return "GetCaptchaRepValidationError" }

// Error satisfies the builtin error interface
func (e GetCaptchaRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCaptchaRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCaptchaRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCaptchaRepValidationError{}

// Validate checks the field values on VerifyCaptchaReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *VerifyCaptchaReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyCaptchaReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyCaptchaReqMultiError, or nil if none found.
func (m *VerifyCaptchaReq) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyCaptchaReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetLastIp()); ip == nil {
		err := VerifyCaptchaReqValidationError{
			field:  "LastIp",
			reason: "value must be a valid IP address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return VerifyCaptchaReqMultiError(errors)
	}

	return nil
}

// VerifyCaptchaReqMultiError is an error wrapping multiple validation errors
// returned by VerifyCaptchaReq.ValidateAll() if the designated constraints
// aren't met.
type VerifyCaptchaReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyCaptchaReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyCaptchaReqMultiError) AllErrors() []error { return m }

// VerifyCaptchaReqValidationError is the validation error returned by
// VerifyCaptchaReq.Validate if the designated constraints aren't met.
type VerifyCaptchaReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyCaptchaReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyCaptchaReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyCaptchaReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyCaptchaReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyCaptchaReqValidationError) ErrorName() string { return "VerifyCaptchaReqValidationError" }

// Error satisfies the builtin error interface
func (e VerifyCaptchaReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyCaptchaReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyCaptchaReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyCaptchaReqValidationError{}

// Validate checks the field values on VerifyCaptchaRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *VerifyCaptchaRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyCaptchaRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyCaptchaRepMultiError, or nil if none found.
func (m *VerifyCaptchaRep) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyCaptchaRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return VerifyCaptchaRepMultiError(errors)
	}

	return nil
}

// VerifyCaptchaRepMultiError is an error wrapping multiple validation errors
// returned by VerifyCaptchaRep.ValidateAll() if the designated constraints
// aren't met.
type VerifyCaptchaRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyCaptchaRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyCaptchaRepMultiError) AllErrors() []error { return m }

// VerifyCaptchaRepValidationError is the validation error returned by
// VerifyCaptchaRep.Validate if the designated constraints aren't met.
type VerifyCaptchaRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyCaptchaRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyCaptchaRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyCaptchaRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyCaptchaRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyCaptchaRepValidationError) ErrorName() string { return "VerifyCaptchaRepValidationError" }

// Error satisfies the builtin error interface
func (e VerifyCaptchaRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyCaptchaRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyCaptchaRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyCaptchaRepValidationError{}

// Validate checks the field values on FireWallVerifyReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FireWallVerifyReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FireWallVerifyReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FireWallVerifyReqMultiError, or nil if none found.
func (m *FireWallVerifyReq) ValidateAll() error {
	return m.validate(true)
}

func (m *FireWallVerifyReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetLastIp()); ip == nil {
		err := FireWallVerifyReqValidationError{
			field:  "LastIp",
			reason: "value must be a valid IP address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FireWallVerifyReqMultiError(errors)
	}

	return nil
}

// FireWallVerifyReqMultiError is an error wrapping multiple validation errors
// returned by FireWallVerifyReq.ValidateAll() if the designated constraints
// aren't met.
type FireWallVerifyReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FireWallVerifyReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FireWallVerifyReqMultiError) AllErrors() []error { return m }

// FireWallVerifyReqValidationError is the validation error returned by
// FireWallVerifyReq.Validate if the designated constraints aren't met.
type FireWallVerifyReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FireWallVerifyReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FireWallVerifyReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FireWallVerifyReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FireWallVerifyReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FireWallVerifyReqValidationError) ErrorName() string {
	return "FireWallVerifyReqValidationError"
}

// Error satisfies the builtin error interface
func (e FireWallVerifyReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFireWallVerifyReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FireWallVerifyReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FireWallVerifyReqValidationError{}

// Validate checks the field values on FireWallVerifyRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FireWallVerifyRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FireWallVerifyRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FireWallVerifyRepMultiError, or nil if none found.
func (m *FireWallVerifyRep) ValidateAll() error {
	return m.validate(true)
}

func (m *FireWallVerifyRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return FireWallVerifyRepMultiError(errors)
	}

	return nil
}

// FireWallVerifyRepMultiError is an error wrapping multiple validation errors
// returned by FireWallVerifyRep.ValidateAll() if the designated constraints
// aren't met.
type FireWallVerifyRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FireWallVerifyRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FireWallVerifyRepMultiError) AllErrors() []error { return m }

// FireWallVerifyRepValidationError is the validation error returned by
// FireWallVerifyRep.Validate if the designated constraints aren't met.
type FireWallVerifyRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FireWallVerifyRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FireWallVerifyRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FireWallVerifyRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FireWallVerifyRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FireWallVerifyRepValidationError) ErrorName() string {
	return "FireWallVerifyRepValidationError"
}

// Error satisfies the builtin error interface
func (e FireWallVerifyRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFireWallVerifyRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FireWallVerifyRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FireWallVerifyRepValidationError{}

// Validate checks the field values on FireWallListReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FireWallListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FireWallListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FireWallListReqMultiError, or nil if none found.
func (m *FireWallListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *FireWallListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := FireWallListReqValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPageSize() <= 0 {
		err := FireWallListReqValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FireWallListReqMultiError(errors)
	}

	return nil
}

// FireWallListReqMultiError is an error wrapping multiple validation errors
// returned by FireWallListReq.ValidateAll() if the designated constraints
// aren't met.
type FireWallListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FireWallListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FireWallListReqMultiError) AllErrors() []error { return m }

// FireWallListReqValidationError is the validation error returned by
// FireWallListReq.Validate if the designated constraints aren't met.
type FireWallListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FireWallListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FireWallListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FireWallListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FireWallListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FireWallListReqValidationError) ErrorName() string { return "FireWallListReqValidationError" }

// Error satisfies the builtin error interface
func (e FireWallListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFireWallListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FireWallListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FireWallListReqValidationError{}

// Validate checks the field values on FireWallListRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FireWallListRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FireWallListRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FireWallListRepMultiError, or nil if none found.
func (m *FireWallListRep) ValidateAll() error {
	return m.validate(true)
}

func (m *FireWallListRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FireWallListRepMultiError(errors)
	}

	return nil
}

// FireWallListRepMultiError is an error wrapping multiple validation errors
// returned by FireWallListRep.ValidateAll() if the designated constraints
// aren't met.
type FireWallListRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FireWallListRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FireWallListRepMultiError) AllErrors() []error { return m }

// FireWallListRepValidationError is the validation error returned by
// FireWallListRep.Validate if the designated constraints aren't met.
type FireWallListRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FireWallListRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FireWallListRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FireWallListRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FireWallListRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FireWallListRepValidationError) ErrorName() string { return "FireWallListRepValidationError" }

// Error satisfies the builtin error interface
func (e FireWallListRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFireWallListRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FireWallListRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FireWallListRepValidationError{}

// Validate checks the field values on UpTokenRep with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpTokenRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpTokenRep with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpTokenRepMultiError, or
// nil if none found.
func (m *UpTokenRep) ValidateAll() error {
	return m.validate(true)
}

func (m *UpTokenRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return UpTokenRepMultiError(errors)
	}

	return nil
}

// UpTokenRepMultiError is an error wrapping multiple validation errors
// returned by UpTokenRep.ValidateAll() if the designated constraints aren't met.
type UpTokenRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpTokenRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpTokenRepMultiError) AllErrors() []error { return m }

// UpTokenRepValidationError is the validation error returned by
// UpTokenRep.Validate if the designated constraints aren't met.
type UpTokenRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpTokenRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpTokenRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpTokenRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpTokenRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpTokenRepValidationError) ErrorName() string { return "UpTokenRepValidationError" }

// Error satisfies the builtin error interface
func (e UpTokenRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpTokenRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpTokenRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpTokenRepValidationError{}

// Validate checks the field values on UploadFileReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UploadFileReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadFileReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UploadFileReqMultiError, or
// nil if none found.
func (m *UploadFileReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadFileReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for File

	// no validation rules for Path

	if len(errors) > 0 {
		return UploadFileReqMultiError(errors)
	}

	return nil
}

// UploadFileReqMultiError is an error wrapping multiple validation errors
// returned by UploadFileReq.ValidateAll() if the designated constraints
// aren't met.
type UploadFileReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadFileReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadFileReqMultiError) AllErrors() []error { return m }

// UploadFileReqValidationError is the validation error returned by
// UploadFileReq.Validate if the designated constraints aren't met.
type UploadFileReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadFileReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadFileReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadFileReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadFileReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadFileReqValidationError) ErrorName() string { return "UploadFileReqValidationError" }

// Error satisfies the builtin error interface
func (e UploadFileReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadFileReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadFileReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadFileReqValidationError{}

// Validate checks the field values on UploadFileRep with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UploadFileRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadFileRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UploadFileRepMultiError, or
// nil if none found.
func (m *UploadFileRep) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadFileRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for File

	if len(errors) > 0 {
		return UploadFileRepMultiError(errors)
	}

	return nil
}

// UploadFileRepMultiError is an error wrapping multiple validation errors
// returned by UploadFileRep.ValidateAll() if the designated constraints
// aren't met.
type UploadFileRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadFileRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadFileRepMultiError) AllErrors() []error { return m }

// UploadFileRepValidationError is the validation error returned by
// UploadFileRep.Validate if the designated constraints aren't met.
type UploadFileRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadFileRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadFileRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadFileRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadFileRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadFileRepValidationError) ErrorName() string { return "UploadFileRepValidationError" }

// Error satisfies the builtin error interface
func (e UploadFileRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadFileRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadFileRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadFileRepValidationError{}
