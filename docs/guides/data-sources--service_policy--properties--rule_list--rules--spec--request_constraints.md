---
page_title: "rule_list.rules.spec.request_constraints"
subcategory: "Security"
description: "rule_list.rules.spec.request_constraints for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 20379, "body_sha256": "sha256:1746c37bea4e21363e181b5bc09149b9b9e9a96e6f078d2dff653281ed51173f", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_count_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_key_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_value_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_header_count_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_header_key_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_header_value_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_parameter_count_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_parameter_name_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_parameter_value_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_query_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_request_line_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_request_size_none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints:max_url_size_none"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:request_constraints", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "path": "docs/guides/data-sources--service_policy--properties--rule_list--rules--spec--request_constraints.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "request_constraints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/request_constraints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.request_constraints for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.request_constraints

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- [rule_list](data-sources--service_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md)
- rule_list.rules.spec.request_constraints

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for request constraints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_cookie_count_choice": "[\"max_cookie_count_exceeds\",\"max_cookie_count_none\"]",
  "x-ves-oneof-field-max_cookie_key_size_choice": "[\"max_cookie_key_size_exceeds\",\"max_cookie_key_size_none\"]",
  "x-ves-oneof-field-max_cookie_value_size_choice": "[\"max_cookie_value_size_exceeds\",\"max_cookie_value_size_none\"]",
  "x-ves-oneof-field-max_header_count_choice": "[\"max_header_count_exceeds\",\"max_header_count_none\"]",
  "x-ves-oneof-field-max_header_key_size_choice": "[\"max_header_key_size_exceeds\",\"max_header_key_size_none\"]",
  "x-ves-oneof-field-max_header_value_size_choice": "[\"max_header_value_size_exceeds\",\"max_header_value_size_none\"]",
  "x-ves-oneof-field-max_parameter_count_choice": "[\"max_parameter_count_exceeds\",\"max_parameter_count_none\"]",
  "x-ves-oneof-field-max_parameter_name_size_choice": "[\"max_parameter_name_size_exceeds\",\"max_parameter_name_size_none\"]",
  "x-ves-oneof-field-max_parameter_value_size_choice": "[\"max_parameter_value_size_exceeds\",\"max_parameter_value_size_none\"]",
  "x-ves-oneof-field-max_query_size_choice": "[\"max_query_size_exceeds\",\"max_query_size_none\"]",
  "x-ves-oneof-field-max_request_line_size_choice": "[\"max_request_line_size_exceeds\",\"max_request_line_size_none\"]",
  "x-ves-oneof-field-max_request_size_choice": "[\"max_request_size_exceeds\",\"max_request_size_none\"]",
  "x-ves-oneof-field-max_url_size_choice": "[\"max_url_size_exceeds\",\"max_url_size_none\"]"
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--request_constraints--max_cookie_count_exceeds"></a>

### max_cookie_count_exceeds property

Type: `"number"`. Computed.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_count_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_cookie_key_size_exceeds"></a>

### max_cookie_key_size_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_key_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_key_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_cookie_value_size_exceeds"></a>

### max_cookie_value_size_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

- [max_cookie_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_value_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_header_count_exceeds"></a>

### max_header_count_exceeds property

Type: `"number"`. Computed.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  }
}
```

- [max_header_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_count_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_header_key_size_exceeds"></a>

### max_header_key_size_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_header_key_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_key_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_header_value_size_exceeds"></a>

### max_header_value_size_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [max_header_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_value_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_parameter_count_exceeds"></a>

### max_parameter_count_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_count_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_parameter_name_size_exceeds"></a>

### max_parameter_name_size_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_name_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_name_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_parameter_value_size_exceeds"></a>

### max_parameter_value_size_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1073741824,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  }
}
```

- [max_parameter_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_value_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_query_size_exceeds"></a>

### max_query_size_exceeds property

Type: `"number"`. Computed.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [max_query_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_query_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_request_line_size_exceeds"></a>

### max_request_line_size_exceeds property

Type: `"number"`. Computed.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_request_line_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_request_size_exceeds"></a>

### max_request_size_exceeds property

Type: `"number"`. Computed.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_request_size_none.md): complete subsection reference.

<a id="schema-rule_list--rules--spec--request_constraints--max_url_size_exceeds"></a>

### max_url_size_exceeds property

Type: `"number"`. Computed.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_url_size_none.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.request_constraints.max_cookie_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_count_none.md)
- [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_key_size_none.md)
- [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_value_size_none.md)
- [rule_list.rules.spec.request_constraints.max_header_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_count_none.md)
- [rule_list.rules.spec.request_constraints.max_header_key_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_key_size_none.md)
- [rule_list.rules.spec.request_constraints.max_header_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_header_value_size_none.md)
- [rule_list.rules.spec.request_constraints.max_parameter_count_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_count_none.md)
- [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_name_size_none.md)
- [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_parameter_value_size_none.md)
- [rule_list.rules.spec.request_constraints.max_query_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_query_size_none.md)
- [rule_list.rules.spec.request_constraints.max_request_line_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_request_line_size_none.md)
- [rule_list.rules.spec.request_constraints.max_request_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_request_size_none.md)
- [rule_list.rules.spec.request_constraints.max_url_size_none](data-sources--service_policy--properties--rule_list--rules--spec--request_constraints--max_url_size_none.md)
- [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
