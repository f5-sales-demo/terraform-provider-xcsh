package client

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ResponseField carries source-owned successful-response constraints. It is
// separate from Terraform Required, since response attributes are computed.
type ResponseField struct {
	Name       string
	Type       string
	Required   bool
	MinLength  int
	Format     string
	Pattern    string
	MapFields  []ResponseField
	Properties []ResponseField
}

// ValidateResponseFields runs before assigning response values to state. Errors
// contain schema field names only, never values, signed URLs or parser errors.
func ValidateResponseFields(response map[string]any, fields []ResponseField) error {
	for _, field := range fields {
		value, exists := response[field.Name]
		if !exists || value == nil {
			if field.Required {
				return fmt.Errorf("required response field %q is missing or null", field.Name)
			}
			continue
		}
		if field.Type == "object" && (len(field.MapFields) > 0 || len(field.Properties) > 0) {
			object, ok := value.(map[string]any)
			if !ok {
				return fmt.Errorf("response field %q must be an object", field.Name)
			}
			if len(field.Properties) > 0 {
				if err := ValidateResponseFields(object, field.Properties); err != nil {
					return err
				}
			}
			if len(field.MapFields) > 0 {
				for _, entry := range object {
					item, ok := entry.(map[string]any)
					if !ok {
						return fmt.Errorf("response map %q has an invalid object entry", field.Name)
					}
					if err := ValidateResponseFields(item, field.MapFields); err != nil {
						return err
					}
				}
			}
		}
		if field.Type != "string" {
			continue
		}
		text, ok := value.(string)
		if !ok {
			return fmt.Errorf("response field %q must be a string", field.Name)
		}
		if field.MinLength > 0 && (utf8.RuneCountInString(text) < field.MinLength || strings.TrimSpace(text) == "") {
			return fmt.Errorf("response field %q does not meet its minimum length", field.Name)
		}
		if field.Format == "uri" {
			parsed, err := url.ParseRequestURI(text)
			if err != nil || parsed.Scheme == "" || parsed.Host == "" || parsed.User != nil {
				return fmt.Errorf("response field %q must be an absolute URI without user information", field.Name)
			}
		}
		if field.Pattern != "" {
			pattern, err := regexp.Compile(field.Pattern)
			if err != nil || !pattern.MatchString(text) {
				return fmt.Errorf("response field %q does not match its required pattern", field.Name)
			}
		}
	}
	return nil
}
