// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package errors

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TestXCSHErrorError(t *testing.T) {
	tests := []struct {
		name     string
		err      *XCSHError
		contains []string
	}{
		{
			name: "basic error",
			err: &XCSHError{
				Code:    ErrCodeNotFound,
				Message: "resource not found",
			},
			contains: []string{"NOT_FOUND", "resource not found"},
		},
		{
			name: "error with resource",
			err: &XCSHError{
				Code:     ErrCodeNotFound,
				Message:  "not found",
				Resource: "namespace",
			},
			contains: []string{"resource: namespace"},
		},
		{
			name: "error with operation",
			err: &XCSHError{
				Code:      ErrCodeTimeout,
				Message:   "timeout",
				Operation: "create",
			},
			contains: []string{"operation: create"},
		},
		{
			name: "error with status code",
			err: &XCSHError{
				Code:       ErrCodeServerError,
				Message:    "server error",
				StatusCode: 500,
			},
			contains: []string{"status: 500"},
		},
		{
			name: "error with wrapped error",
			err: &XCSHError{
				Code:    ErrCodeNetworkError,
				Message: "network error",
				Wrapped: errors.New("connection refused"),
			},
			contains: []string{"connection refused"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errStr := tt.err.Error()
			for _, s := range tt.contains {
				if !strings.Contains(errStr, s) {
					t.Errorf("Error() = %q, expected to contain %q", errStr, s)
				}
			}
		})
	}
}

func TestXCSHErrorUnwrap(t *testing.T) {
	innerErr := errors.New("inner error")
	err := &XCSHError{
		Code:    ErrCodeServerError,
		Message: "outer error",
		Wrapped: innerErr,
	}

	unwrapped := err.Unwrap()
	if unwrapped != innerErr {
		t.Errorf("Unwrap() = %v, expected %v", unwrapped, innerErr)
	}

	// Test with nil wrapped
	errNoWrap := &XCSHError{
		Code:    ErrCodeNotFound,
		Message: "no wrap",
	}
	if errNoWrap.Unwrap() != nil {
		t.Error("Unwrap() should return nil when no wrapped error")
	}
}

func TestXCSHErrorIsRetryable(t *testing.T) {
	tests := []struct {
		code     ErrorCode
		expected bool
	}{
		{ErrCodeRateLimit, true},
		{ErrCodeTimeout, true},
		{ErrCodeNetworkError, true},
		{ErrCodeServerError, true},
		{ErrCodeNotFound, false},
		{ErrCodeUnauthorized, false},
		{ErrCodeForbidden, false},
		{ErrCodeConflict, false},
		{ErrCodeBadRequest, false},
		{ErrCodeValidation, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.code), func(t *testing.T) {
			err := &XCSHError{Code: tt.code}
			if err.IsRetryable() != tt.expected {
				t.Errorf("IsRetryable() for %s = %v, expected %v", tt.code, err.IsRetryable(), tt.expected)
			}
		})
	}
}

func TestXCSHErrorIsNotFound(t *testing.T) {
	tests := []struct {
		code     ErrorCode
		expected bool
	}{
		{ErrCodeNotFound, true},
		{ErrCodeUnauthorized, false},
		{ErrCodeServerError, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.code), func(t *testing.T) {
			err := &XCSHError{Code: tt.code}
			if err.IsNotFound() != tt.expected {
				t.Errorf("IsNotFound() for %s = %v, expected %v", tt.code, err.IsNotFound(), tt.expected)
			}
		})
	}
}

func TestNewAPIError(t *testing.T) {
	tests := []struct {
		name         string
		statusCode   int
		body         []byte
		resource     string
		operation    string
		expectedCode ErrorCode
	}{
		{
			name:         "not found",
			statusCode:   http.StatusNotFound,
			body:         nil,
			resource:     "namespace",
			operation:    "read",
			expectedCode: ErrCodeNotFound,
		},
		{
			name:         "unauthorized",
			statusCode:   http.StatusUnauthorized,
			body:         nil,
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeUnauthorized,
		},
		{
			name:         "forbidden",
			statusCode:   http.StatusForbidden,
			body:         nil,
			resource:     "namespace",
			operation:    "delete",
			expectedCode: ErrCodeForbidden,
		},
		{
			name:         "conflict",
			statusCode:   http.StatusConflict,
			body:         nil,
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeConflict,
		},
		{
			name:         "rate limit",
			statusCode:   http.StatusTooManyRequests,
			body:         nil,
			resource:     "namespace",
			operation:    "list",
			expectedCode: ErrCodeRateLimit,
		},
		{
			name:         "bad request",
			statusCode:   http.StatusBadRequest,
			body:         nil,
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeBadRequest,
		},
		{
			name:         "server error",
			statusCode:   http.StatusInternalServerError,
			body:         nil,
			resource:     "namespace",
			operation:    "read",
			expectedCode: ErrCodeServerError,
		},
		{
			name:         "with JSON body",
			statusCode:   http.StatusBadRequest,
			body:         []byte(`{"code": "INVALID_ARGUMENT", "message": "field is required"}`),
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeBadRequest,
		},
		{
			name:         "with invalid JSON body",
			statusCode:   http.StatusBadRequest,
			body:         []byte(`not json`),
			resource:     "namespace",
			operation:    "create",
			expectedCode: ErrCodeBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewAPIError(tt.statusCode, tt.body, tt.resource, tt.operation)

			if err.Code != tt.expectedCode {
				t.Errorf("NewAPIError() code = %v, expected %v", err.Code, tt.expectedCode)
			}
			if err.Resource != tt.resource {
				t.Errorf("NewAPIError() resource = %v, expected %v", err.Resource, tt.resource)
			}
			if err.Operation != tt.operation {
				t.Errorf("NewAPIError() operation = %v, expected %v", err.Operation, tt.operation)
			}
			if err.StatusCode != tt.statusCode {
				t.Errorf("NewAPIError() statusCode = %v, expected %v", err.StatusCode, tt.statusCode)
			}
		})
	}
}

func TestNewAPIError_IncludesDetailsInErrorString(t *testing.T) {
	jsonBody := []byte(`{
		"code": "INVALID_ARGUMENT",
		"message": "Validation failed",
		"details": [
			{"@type": "type.googleapis.com/ves.io.schema.ErrorDetail", "reason": "name is required", "domain": "ves.io"}
		]
	}`)
	err := NewAPIError(400, jsonBody, "http_loadbalancer", "create")
	if !strings.Contains(err.Error(), "name is required") {
		t.Errorf("Expected err.Error() to contain 'name is required', got: %s", err.Error())
	}
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("namespace", "test-ns", "system")

	if err.Code != ErrCodeNotFound {
		t.Errorf("NewNotFoundError() code = %v, expected %v", err.Code, ErrCodeNotFound)
	}
	if !strings.Contains(err.Message, "test-ns") {
		t.Errorf("NewNotFoundError() message should contain name, got %v", err.Message)
	}
	if !strings.Contains(err.Message, "system") {
		t.Errorf("NewNotFoundError() message should contain namespace, got %v", err.Message)
	}
}

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("namespace", "name", "must not be empty")

	if err.Code != ErrCodeValidation {
		t.Errorf("NewValidationError() code = %v, expected %v", err.Code, ErrCodeValidation)
	}
	if !strings.Contains(err.Message, "name") {
		t.Errorf("NewValidationError() message should contain field name, got %v", err.Message)
	}
}

func TestNewTimeoutError(t *testing.T) {
	wrapped := errors.New("context deadline exceeded")
	err := NewTimeoutError("namespace", "create", wrapped)

	if err.Code != ErrCodeTimeout {
		t.Errorf("NewTimeoutError() code = %v, expected %v", err.Code, ErrCodeTimeout)
	}
	if err.Wrapped != wrapped {
		t.Error("NewTimeoutError() should wrap the original error")
	}
}

func TestNewNetworkError(t *testing.T) {
	wrapped := errors.New("connection refused")
	err := NewNetworkError(wrapped)

	if err.Code != ErrCodeNetworkError {
		t.Errorf("NewNetworkError() code = %v, expected %v", err.Code, ErrCodeNetworkError)
	}
	if err.Wrapped != wrapped {
		t.Error("NewNetworkError() should wrap the original error")
	}
}

func TestNewConfigurationError(t *testing.T) {
	err := NewConfigurationError("API URL is required")

	if err.Code != ErrCodeConfiguration {
		t.Errorf("NewConfigurationError() code = %v, expected %v", err.Code, ErrCodeConfiguration)
	}
	if err.Message != "API URL is required" {
		t.Errorf("NewConfigurationError() message = %v, expected %v", err.Message, "API URL is required")
	}
}

func TestAddError(t *testing.T) {
	var diags diag.Diagnostics
	err := &XCSHError{
		Code:    ErrCodeNotFound,
		Message: "resource not found",
	}

	AddError(&diags, err)

	if !diags.HasError() {
		t.Error("AddError() should add an error to diagnostics")
	}
	if len(diags.Errors()) != 1 {
		t.Errorf("AddError() should add exactly one error, got %d", len(diags.Errors()))
	}
}

func TestAddWarning(t *testing.T) {
	var diags diag.Diagnostics

	AddWarning(&diags, "Warning Title", "Warning detail")

	if diags.HasError() {
		t.Error("AddWarning() should not add an error")
	}
	if len(diags.Warnings()) != 1 {
		t.Errorf("AddWarning() should add exactly one warning, got %d", len(diags.Warnings()))
	}
}

func TestAddAttributeError(t *testing.T) {
	var diags diag.Diagnostics

	AddAttributeError(&diags, "name", "Invalid Value", "Name must not be empty")

	if !diags.HasError() {
		t.Error("AddAttributeError() should add an error")
	}
}

func TestCreateDiagnostic(t *testing.T) {
	t.Run("with XCSHError", func(t *testing.T) {
		xcshErr := &XCSHError{
			Code:    ErrCodeNotFound,
			Message: "not found",
		}

		diags := CreateDiagnostic("reading", "namespace", xcshErr)

		if !diags.HasError() {
			t.Error("CreateDiagnostic() should create error diagnostic")
		}
	})

	t.Run("with standard error", func(t *testing.T) {
		stdErr := errors.New("standard error")

		diags := CreateDiagnostic("creating", "namespace", stdErr)

		if !diags.HasError() {
			t.Error("CreateDiagnostic() should create error diagnostic")
		}
	})
}

func TestWrapError(t *testing.T) {
	t.Run("wrap standard error", func(t *testing.T) {
		stdErr := errors.New("original error")

		wrapped := WrapError(stdErr, "namespace", "create")

		if wrapped.Resource != "namespace" {
			t.Errorf("WrapError() resource = %v, expected namespace", wrapped.Resource)
		}
		if wrapped.Operation != "create" {
			t.Errorf("WrapError() operation = %v, expected create", wrapped.Operation)
		}
		if wrapped.Wrapped != stdErr {
			t.Error("WrapError() should preserve original error")
		}
	})

	t.Run("wrap XCSHError", func(t *testing.T) {
		xcshErr := &XCSHError{
			Code:    ErrCodeNotFound,
			Message: "not found",
		}

		wrapped := WrapError(xcshErr, "namespace", "read")

		if wrapped.Resource != "namespace" {
			t.Errorf("WrapError() should update resource, got %v", wrapped.Resource)
		}
		if wrapped.Operation != "read" {
			t.Errorf("WrapError() should update operation, got %v", wrapped.Operation)
		}
		if wrapped.Code != ErrCodeNotFound {
			t.Error("WrapError() should preserve original code")
		}
	})
}

func TestAPIErrorResponseParsing(t *testing.T) {
	jsonBody := `{
		"code": "INVALID_ARGUMENT",
		"message": "Field validation failed",
		"details": [
			{
				"@type": "type.googleapis.com/google.rpc.BadRequest",
				"reason": "name is required",
				"domain": "xcsh.io"
			}
		]
	}`

	var apiErr APIErrorResponse
	if err := json.Unmarshal([]byte(jsonBody), &apiErr); err != nil {
		t.Fatalf("Failed to unmarshal APIErrorResponse: %v", err)
	}

	if apiErr.Code != "INVALID_ARGUMENT" {
		t.Errorf("APIErrorResponse.Code = %v, expected INVALID_ARGUMENT", apiErr.Code)
	}
	if apiErr.Message != "Field validation failed" {
		t.Errorf("APIErrorResponse.Message = %v, expected Field validation failed", apiErr.Message)
	}
	if len(apiErr.Details) != 1 {
		t.Errorf("APIErrorResponse.Details length = %d, expected 1", len(apiErr.Details))
	}
}

func TestErrorCodeConstants(t *testing.T) {
	// Verify error codes are distinct
	codes := []ErrorCode{
		ErrCodeNotFound,
		ErrCodeUnauthorized,
		ErrCodeForbidden,
		ErrCodeConflict,
		ErrCodeRateLimit,
		ErrCodeServerError,
		ErrCodeBadRequest,
		ErrCodeTimeout,
		ErrCodeNetworkError,
		ErrCodeValidation,
		ErrCodeStateRead,
		ErrCodeStateWrite,
		ErrCodeConfiguration,
	}

	seen := make(map[ErrorCode]bool)
	for _, code := range codes {
		if seen[code] {
			t.Errorf("Duplicate error code: %v", code)
		}
		seen[code] = true
	}
}

func TestSafeAPIDiagnostics_DeterministicOrderingAndStructuredDetails(t *testing.T) {
	err := &XCSHError{
		Code:    ErrCodeBadRequest,
		Message: "Invalid input",
		Details: map[string]interface{}{
			"api_details": []APIErrorDetail{
				{Domain: "domainB", Reason: "reasonB"},
				{Domain: "domainA", Reason: "reasonA"},
				{Domain: "domainC", Reason: "reasonC"},
			},
		},
	}
	output := err.Error()
	expectedSub := "[domainB: reasonB] [domainA: reasonA] [domainC: reasonC]"
	if !strings.Contains(output, expectedSub) {
		t.Errorf("expected deterministic ordering %q in output, got: %s", expectedSub, output)
	}
}

func TestSafeAPIDiagnostics_AbsentMessages(t *testing.T) {
	err := &XCSHError{
		Code: ErrCodeServerError,
	}
	output := err.Error()
	if strings.Contains(output, "details:") {
		t.Errorf("did not expect details: tag for empty details")
	}
	if strings.Contains(output, "[]") {
		t.Errorf("did not expect empty brackets")
	}
}

func TestSafeAPIDiagnostics_CappingLimitsAndUTF8Splitting(t *testing.T) {
	// A single UTF-8 character that is 3 bytes long
	multiByte := "🔥"
	longDomain := strings.Repeat("A", 255) + multiByte // 258 bytes
	err := &XCSHError{
		Code:    ErrCodeBadRequest,
		Message: "Invalid input",
		Details: map[string]interface{}{
			"api_details": []APIErrorDetail{
				{Domain: longDomain, Reason: "reason"},
			},
		},
	}
	output := err.Error()
	// Domain should be truncated to 256 bytes.
	// 255 A's + 1 byte of the 3-byte emoji would split it.
	// The truncator should backtrack to 255 A's.
	if strings.Contains(output, multiByte) {
		t.Errorf("expected multi-byte character to be truncated")
	}
	if strings.Contains(output, "\ufffd") {
		t.Errorf("expected no replacement characters from bad UTF-8 slicing")
	}

	// Test 1024 cap
	longRaw := strings.Repeat("B", 1023) + multiByte
	err2 := &XCSHError{
		Code:    ErrCodeBadRequest,
		Message: "Invalid input",
		Details: map[string]interface{}{
			"raw_response": longRaw,
		},
	}
	output2 := err2.Error()
	detailsIndex := strings.Index(output2, "details: ")
	if detailsIndex == -1 {
		t.Fatalf("expected details in output")
	}
	detailsStr := output2[detailsIndex+len("details: "):]
	if len(detailsStr) > 1024 {
		t.Errorf("expected details to be capped at 1024 bytes, got %d", len(detailsStr))
	}
	if strings.Contains(detailsStr, multiByte) {
		t.Errorf("expected multi-byte character to be truncated at 1024 cap")
	}
}

func TestSafeAPIDiagnostics_NormalizationAndControlCharacterRemoval(t *testing.T) {
	err := &XCSHError{
		Code:    ErrCodeBadRequest,
		Message: "Bad",
		Details: map[string]interface{}{
			"raw_response": "line1\r\n\tline2\x00\x1b\xff end",
		},
	}
	output := err.Error()
	if strings.Contains(output, "\n") || strings.Contains(output, "\t") || strings.Contains(output, "\r") {
		t.Errorf("expected whitespace to be normalized")
	}
	if strings.Contains(output, "\x00") || strings.Contains(output, "\x1b") {
		t.Errorf("expected control characters to be removed")
	}
	if strings.Contains(output, "\xff") {
		t.Errorf("expected invalid utf-8 to be removed")
	}
	expectedEnd := "line1 line2 end"
	if !strings.Contains(output, expectedEnd) {
		t.Errorf("expected normalized string %q, got: %s", expectedEnd, output)
	}
}

func TestSafeAPIDiagnostics_RedactionClasses(t *testing.T) {
	const secretSentinel = "S3CR3T_S3NT1N3L"
	payloads := []string{
		"Authorization: Bearer " + secretSentinel,
		"api_key=" + secretSentinel,
		"password : " + secretSentinel,
		"cookie=session=" + secretSentinel,
		"-----BEGIN PRIVATE KEY-----\n" + secretSentinel + "\n-----END PRIVATE KEY-----",
		"{\"api_key\": \"" + secretSentinel + "\"}",
		"password='" + secretSentinel + "'",
		"Authorization: \"Bearer " + secretSentinel + "\"",
	}

	for _, payload := range payloads {
		t.Run(payload[:10], func(t *testing.T) {
			err := &XCSHError{
				Code:    ErrCodeBadRequest,
				Message: "Bad",
				Details: map[string]interface{}{
					"raw_response": payload,
				},
			}
			output := err.Error()
			if strings.Contains(output, secretSentinel) {
				t.Errorf("failed to redact secret sentinel in payload: %s\nOutput: %s", payload, output)
			}
			if !strings.Contains(output, "[REDACTED") {
				t.Errorf("expected redaction marker in output: %s", output)
			}
		})
	}
}

func TestSafeAPIDiagnostics_RedactionPreTruncation(t *testing.T) {
	// A secret that gets split exactly at the 256 byte mark
	const secretSentinel = "S3CR3T_S3NT1N3L"

	// Create a string that puts the secret right at the truncation boundary
	// dReason is truncated at 256.
	padding := strings.Repeat("A", 250)
	payload := padding + "api_key=" + secretSentinel

	err := &XCSHError{
		Code:    ErrCodeBadRequest,
		Message: "Bad request",
		Details: map[string]interface{}{
			"api_details": []APIErrorDetail{
				{Domain: "xcsh.io", Reason: payload},
			},
		},
	}
	output := err.Error()
	if strings.Contains(output, secretSentinel) {
		t.Errorf("failed to redact secret near truncation boundary. Output: %s", output)
	}
	// We don't check for [REDACTED] because the marker itself gets truncated.
}

func TestSafeAPIDiagnostics_CreateDiagnosticFinalOutput(t *testing.T) {
	const secretSentinel = "S3CR3T_S3NT1N3L"
	err := &XCSHError{
		Code:    ErrCodeBadRequest,
		Message: "Bad request with secrets",
		Details: map[string]interface{}{
			"raw_response": "Bearer " + secretSentinel + " \x00\x1b\xff \r\n end",
		},
	}

	diags := CreateDiagnostic("Creating", "Resource", err)
	if !diags.HasError() {
		t.Fatalf("expected errors")
	}
	diagMsg := diags.Errors()[0].Detail()

	if strings.Contains(diagMsg, secretSentinel) {
		t.Errorf("secret sentinel leaked to CreateDiagnostic output")
	}
	if strings.Contains(diagMsg, "\x00") || strings.Contains(diagMsg, "\n") {
		t.Errorf("control characters or newlines leaked to CreateDiagnostic output")
	}
	if !strings.Contains(diagMsg, "[REDACTED]") {
		t.Errorf("expected redaction marker in CreateDiagnostic output")
	}
}
