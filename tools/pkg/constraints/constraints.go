// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package constraints

// Parsed represents extracted x-f5xc-constraints data.
type Parsed struct {
	MinLength int
	MaxLength int
	Pattern   string
	MinItems  int
	MaxItems  int
}

// Parse extracts constraint data from x-f5xc-constraints map.
// Returns nil for nil input, low-confidence, or non-deterministic constraints.
// Only uses constraints where metadata.deterministic==true OR metadata.confidence >= 0.9.
func Parse(raw map[string]interface{}) *Parsed {
	if raw == nil {
		return nil
	}
	deterministic := false
	confidence := 0.0
	if meta, ok := raw["metadata"].(map[string]interface{}); ok {
		if d, ok := meta["deterministic"].(bool); ok {
			deterministic = d
		}
		if c, ok := meta["confidence"].(float64); ok {
			confidence = c
		}
	}
	if !deterministic && confidence < 0.9 {
		return nil
	}
	p := &Parsed{}
	if v, ok := raw["minLength"].(float64); ok {
		p.MinLength = int(v)
	}
	if v, ok := raw["maxLength"].(float64); ok {
		p.MaxLength = int(v)
	}
	if v, ok := raw["pattern"].(string); ok {
		p.Pattern = v
	}
	if v, ok := raw["minItems"].(float64); ok {
		p.MinItems = int(v)
	}
	if v, ok := raw["maxItems"].(float64); ok {
		p.MaxItems = int(v)
	}
	return p
}
