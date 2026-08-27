// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package validators

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Int64Range is one inclusive interval accepted by Int64RangeSetValidator.
type Int64Range struct {
	Minimum int64
	Maximum int64
}

// Int64RangeSetValidator returns a validator that accepts a value when it is in
// any supplied inclusive interval.
func Int64RangeSetValidator(ranges ...Int64Range) validator.Int64 {
	return int64RangeSetValidator{ranges: append([]Int64Range(nil), ranges...)}
}

type int64RangeSetValidator struct {
	ranges []Int64Range
}

func (v int64RangeSetValidator) Description(ctx context.Context) string {
	parts := make([]string, 0, len(v.ranges))
	for _, span := range v.ranges {
		if span.Minimum == span.Maximum {
			parts = append(parts, fmt.Sprintf("%d", span.Minimum))
		} else {
			parts = append(parts, fmt.Sprintf("%d-%d", span.Minimum, span.Maximum))
		}
	}
	return "must be in one of these inclusive ranges: " + strings.Join(parts, ", ")
}

func (v int64RangeSetValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v int64RangeSetValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueInt64()
	for _, span := range v.ranges {
		if value >= span.Minimum && value <= span.Maximum {
			return
		}
	}
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Value Outside Allowed Integer Ranges",
		fmt.Sprintf("Value %d %s.", value, v.Description(ctx)),
	)
}
