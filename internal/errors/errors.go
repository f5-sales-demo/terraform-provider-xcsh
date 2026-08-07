// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package errors provides structured error types for F5 XC Terraform provider
// following Terraform Plugin Framework best practices.
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ErrorCode represents specific error types for better handling
type ErrorCode string

const (
	// API errors
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden    ErrorCode = "FORBIDDEN"
	ErrCodeConflict     ErrorCode = "CONFLICT"
	ErrCodeRateLimit    ErrorCode = "RATE_LIMIT"
	ErrCodeServerError  ErrorCode = "SERVER_ERROR"
	ErrCodeBadRequest   ErrorCode = "BAD_REQUEST"
	ErrCodeTimeout      ErrorCode = "TIMEOUT"
	ErrCodeNetworkError ErrorCode = "NETWORK_ERROR"

	// Resource errors
	ErrCodeValidation    ErrorCode = "VALIDATION"
	ErrCodeStateRead     ErrorCode = "STATE_READ"
	ErrCodeStateWrite    ErrorCode = "STATE_WRITE"
	ErrCodeConfiguration ErrorCode = "CONFIGURATION"
)

// XCSHError is a structured error type for the provider
type XCSHError struct {
	Code       ErrorCode
	Message    string
	Resource   string
	Operation  string
	StatusCode int
	Details    map[string]interface{}
	Wrapped    error
}

// Error implements the error interface
func (e *XCSHError) Error() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "[%s] %s", e.Code, e.Message)

	if e.Resource != "" {
		fmt.Fprintf(&sb, " (resource: %s)", e.Resource)
	}
	if e.Operation != "" {
		fmt.Fprintf(&sb, " (operation: %s)", e.Operation)
	}
	if e.StatusCode != 0 {
		fmt.Fprintf(&sb, " (status: %d)", e.StatusCode)
	}

	if apiCode, ok := e.Details["api_code"]; ok {
		fmt.Fprintf(&sb, " [api_code: %v]", apiCode)
	}

	var detailStr string
	if details, ok := e.Details["api_details"].([]APIErrorDetail); ok && len(details) > 0 {
		var parts []string
		for _, d := range details {
			dReason := redactSensitive(sanitizeText(d.Reason))
			dDomain := redactSensitive(sanitizeText(d.Domain))
			dReason = truncateUTF8(dReason, 256)
			dDomain = truncateUTF8(dDomain, 256)
			parts = append(parts, fmt.Sprintf("[%s: %s]", dDomain, dReason))
		}
		detailStr = strings.Join(parts, " ")
	} else if raw, ok := e.Details["raw_response"].(string); ok {
		detailStr = redactSensitive(sanitizeText(raw))
	}

	if detailStr != "" {
		cappedStr := truncateUTF8(detailStr, 1024)
		fmt.Fprintf(&sb, " details: %s", cappedStr)
	}

	if e.Wrapped != nil {
		fmt.Fprintf(&sb, ": %v", e.Wrapped)
	}

	return sb.String()
}

// Unwrap returns the wrapped error
func (e *XCSHError) Unwrap() error {
	return e.Wrapped
}

// IsRetryable returns true if the error is potentially transient
func (e *XCSHError) IsRetryable() bool {
	switch e.Code {
	case ErrCodeRateLimit, ErrCodeTimeout, ErrCodeNetworkError, ErrCodeServerError:
		return true
	default:
		return false
	}
}

// IsNotFound returns true if the resource was not found
func (e *XCSHError) IsNotFound() bool {
	return e.Code == ErrCodeNotFound
}

// APIErrorDetail represents a single detail object in an API error response
type APIErrorDetail struct {
	Type   string `json:"@type"`
	Reason string `json:"reason"`
	Domain string `json:"domain"`
}

// APIErrorResponse represents the F5 XC API error response structure
type APIErrorResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Details []APIErrorDetail `json:"details"`
}

// NewAPIError creates an error from an API response
func NewAPIError(statusCode int, body []byte, resource, operation string) *XCSHError {
	err := &XCSHError{
		Resource:   resource,
		Operation:  operation,
		StatusCode: statusCode,
		Details:    make(map[string]interface{}),
	}

	// Set error code based on status
	switch statusCode {
	case http.StatusNotFound:
		err.Code = ErrCodeNotFound
		err.Message = fmt.Sprintf("%s not found", resource)
	case http.StatusUnauthorized:
		err.Code = ErrCodeUnauthorized
		err.Message = "API authentication failed - check your API token"
	case http.StatusForbidden:
		err.Code = ErrCodeForbidden
		err.Message = "Access denied - insufficient permissions"
	case http.StatusConflict:
		err.Code = ErrCodeConflict
		err.Message = fmt.Sprintf("%s already exists or has a conflicting state", resource)
	case http.StatusTooManyRequests:
		err.Code = ErrCodeRateLimit
		err.Message = "API rate limit exceeded - retry after a delay"
	case http.StatusBadRequest:
		err.Code = ErrCodeBadRequest
		err.Message = "Invalid request parameters"
	default:
		if statusCode >= 500 {
			err.Code = ErrCodeServerError
			err.Message = "F5 XC API server error"
		} else {
			err.Code = ErrCodeBadRequest
			err.Message = fmt.Sprintf("API request failed with status %d", statusCode)
		}
	}

	// Try to parse API error response for more details
	if len(body) > 0 {
		var apiErr APIErrorResponse
		if jsonErr := json.Unmarshal(body, &apiErr); jsonErr == nil {
			if apiErr.Message != "" {
				err.Message = apiErr.Message
			}
			if apiErr.Code != "" {
				err.Details["api_code"] = apiErr.Code
			}
			if len(apiErr.Details) > 0 {
				err.Details["api_details"] = apiErr.Details
			}
		} else {
			// Store raw response if JSON parsing fails
			err.Details["raw_response"] = string(body)
		}
	}

	return err
}

// NewNotFoundError creates a not found error
func NewNotFoundError(resource, name, namespace string) *XCSHError {
	return &XCSHError{
		Code:     ErrCodeNotFound,
		Message:  fmt.Sprintf("%s '%s' not found in namespace '%s'", resource, name, namespace),
		Resource: resource,
		Details: map[string]interface{}{
			"name":      name,
			"namespace": namespace,
		},
	}
}

// NewValidationError creates a validation error
func NewValidationError(resource, field, message string) *XCSHError {
	return &XCSHError{
		Code:     ErrCodeValidation,
		Message:  fmt.Sprintf("validation failed for %s.%s: %s", resource, field, message),
		Resource: resource,
		Details: map[string]interface{}{
			"field": field,
		},
	}
}

// NewTimeoutError creates a timeout error
func NewTimeoutError(resource, operation string, wrapped error) *XCSHError {
	return &XCSHError{
		Code:      ErrCodeTimeout,
		Message:   fmt.Sprintf("operation timed out: %s %s", operation, resource),
		Resource:  resource,
		Operation: operation,
		Wrapped:   wrapped,
	}
}

// NewNetworkError creates a network error
func NewNetworkError(wrapped error) *XCSHError {
	return &XCSHError{
		Code:    ErrCodeNetworkError,
		Message: "network error communicating with F5 XC API",
		Wrapped: wrapped,
	}
}

// NewConfigurationError creates a configuration error
func NewConfigurationError(message string) *XCSHError {
	return &XCSHError{
		Code:    ErrCodeConfiguration,
		Message: message,
	}
}

// DiagnosticHelpers provides methods to add errors to diagnostics

// AddError adds a structured error to diagnostics
func AddError(diags *diag.Diagnostics, err *XCSHError) {
	caser := cases.Title(language.English)
	summary := fmt.Sprintf("%s Error", caser.String(string(err.Code)))
	diags.AddError(summary, err.Error())
}

// AddWarning adds a warning to diagnostics
func AddWarning(diags *diag.Diagnostics, summary, detail string) {
	diags.AddWarning(summary, detail)
}

// AddAttributeError adds an attribute-specific error
func AddAttributeError(diags *diag.Diagnostics, path, summary, detail string) {
	diags.AddError(fmt.Sprintf("%s: %s", path, summary), detail)
}

// CreateDiagnostic creates a simple error diagnostic from a resource operation error
func CreateDiagnostic(operation, resourceType string, err error) diag.Diagnostics {
	var diags diag.Diagnostics

	if xcshErr, ok := err.(*XCSHError); ok {
		AddError(&diags, xcshErr)
	} else {
		diags.AddError(
			fmt.Sprintf("Error %s %s", operation, resourceType),
			err.Error(),
		)
	}

	return diags
}

// WrapError wraps an error with additional context
func WrapError(err error, resource, operation string) *XCSHError {
	if xcshErr, ok := err.(*XCSHError); ok {
		xcshErr.Resource = resource
		xcshErr.Operation = operation
		return xcshErr
	}

	return &XCSHError{
		Code:      ErrCodeServerError,
		Message:   err.Error(),
		Resource:  resource,
		Operation: operation,
		Wrapped:   err,
	}
}

var (
	// Redaction patterns
	bearerPattern = regexp.MustCompile(`(?i)(bearer|authorization)(["']?\s*[:=]\s*["']?|[\s:=]+["']?)([^\s"']{10,})["']?`)
	kvPattern     = regexp.MustCompile(`(?i)(token|api_?key|password|cookie|secret)(["']?\s*[:=]\s*["']?|[\s:=]+["']?)([^\s"']{5,})["']?`)
	pemPattern    = regexp.MustCompile(`(?s)-----BEGIN.*?-----[\s\S]*?-----END.*?-----`)
)

func sanitizeText(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		w = width
		if r == utf8.RuneError {
			continue
		}
		if r == '\n' || r == '\r' || r == '\t' || r == ' ' {
			if sb.Len() == 0 || sb.String()[sb.Len()-1] != ' ' {
				sb.WriteRune(' ')
			}
			continue
		}
		if r < 32 || r == 127 {
			continue
		}
		sb.WriteRune(r)
	}
	return strings.TrimSpace(sb.String())
}

func truncateUTF8(s string, maxBytes int) string {
	if len(s) <= maxBytes {
		return s
	}
	// Walk back from maxBytes to find valid rune boundary
	for i := maxBytes; i > 0; i-- {
		if utf8.RuneStart(s[i]) {
			return s[:i]
		}
	}
	return ""
}

func redactSensitive(s string) string {
	s = bearerPattern.ReplaceAllString(s, "${1}${2}[REDACTED]")
	s = kvPattern.ReplaceAllString(s, "${1}${2}[REDACTED]")
	s = pemPattern.ReplaceAllString(s, "[REDACTED_PEM]")
	return s
}
