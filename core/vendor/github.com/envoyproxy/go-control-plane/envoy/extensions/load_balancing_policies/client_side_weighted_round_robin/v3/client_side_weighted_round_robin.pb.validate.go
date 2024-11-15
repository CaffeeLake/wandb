//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/load_balancing_policies/client_side_weighted_round_robin/v3/client_side_weighted_round_robin.proto

package client_side_weighted_round_robinv3

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

// Validate checks the field values on ClientSideWeightedRoundRobin with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ClientSideWeightedRoundRobin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientSideWeightedRoundRobin with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClientSideWeightedRoundRobinMultiError, or nil if none found.
func (m *ClientSideWeightedRoundRobin) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientSideWeightedRoundRobin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEnableOobLoadReport()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "EnableOobLoadReport",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "EnableOobLoadReport",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnableOobLoadReport()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientSideWeightedRoundRobinValidationError{
				field:  "EnableOobLoadReport",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOobReportingPeriod()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "OobReportingPeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "OobReportingPeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOobReportingPeriod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientSideWeightedRoundRobinValidationError{
				field:  "OobReportingPeriod",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBlackoutPeriod()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "BlackoutPeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "BlackoutPeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBlackoutPeriod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientSideWeightedRoundRobinValidationError{
				field:  "BlackoutPeriod",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWeightExpirationPeriod()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "WeightExpirationPeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "WeightExpirationPeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWeightExpirationPeriod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientSideWeightedRoundRobinValidationError{
				field:  "WeightExpirationPeriod",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWeightUpdatePeriod()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "WeightUpdatePeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClientSideWeightedRoundRobinValidationError{
					field:  "WeightUpdatePeriod",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWeightUpdatePeriod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientSideWeightedRoundRobinValidationError{
				field:  "WeightUpdatePeriod",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetErrorUtilizationPenalty(); wrapper != nil {

		if wrapper.GetValue() < 0 {
			err := ClientSideWeightedRoundRobinValidationError{
				field:  "ErrorUtilizationPenalty",
				reason: "value must be greater than or equal to 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ClientSideWeightedRoundRobinMultiError(errors)
	}

	return nil
}

// ClientSideWeightedRoundRobinMultiError is an error wrapping multiple
// validation errors returned by ClientSideWeightedRoundRobin.ValidateAll() if
// the designated constraints aren't met.
type ClientSideWeightedRoundRobinMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientSideWeightedRoundRobinMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientSideWeightedRoundRobinMultiError) AllErrors() []error { return m }

// ClientSideWeightedRoundRobinValidationError is the validation error returned
// by ClientSideWeightedRoundRobin.Validate if the designated constraints
// aren't met.
type ClientSideWeightedRoundRobinValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientSideWeightedRoundRobinValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientSideWeightedRoundRobinValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientSideWeightedRoundRobinValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientSideWeightedRoundRobinValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientSideWeightedRoundRobinValidationError) ErrorName() string {
	return "ClientSideWeightedRoundRobinValidationError"
}

// Error satisfies the builtin error interface
func (e ClientSideWeightedRoundRobinValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientSideWeightedRoundRobin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientSideWeightedRoundRobinValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientSideWeightedRoundRobinValidationError{}