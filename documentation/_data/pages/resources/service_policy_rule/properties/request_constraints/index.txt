---
page_title: "request_constraints"
subcategory: ""
description: "request_constraints for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 23676, "body_sha256": "sha256:14a56ac3eee38b9f0e2ac468a29f8833a976167cf3cfaa5ecd1fd566bc0ddf3b", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_cookie_count_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_cookie_key_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_cookie_value_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_count_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_key_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_value_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_parameter_count_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_parameter_name_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_parameter_value_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_query_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_request_line_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_request_size_none", "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_url_size_none"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/request_constraints/index.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["request_constraints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/request_constraints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_constraints for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# request_constraints

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- request_constraints

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request constraints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_cookie_count_exceeds",
    "max_cookie_count_none"),
  validators.ConflictingObjectAttributes("max_cookie_key_size_exceeds",
    "max_cookie_key_size_none"),
  validators.ConflictingObjectAttributes("max_cookie_value_size_exceeds",
    "max_cookie_value_size_none"),
  validators.ConflictingObjectAttributes("max_header_count_exceeds",
    "max_header_count_none"),
  validators.ConflictingObjectAttributes("max_header_key_size_exceeds",
    "max_header_key_size_none"),
  validators.ConflictingObjectAttributes("max_header_value_size_exceeds",
    "max_header_value_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_count_exceeds",
    "max_parameter_count_none"),
  validators.ConflictingObjectAttributes("max_parameter_name_size_exceeds",
    "max_parameter_name_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_value_size_exceeds",
    "max_parameter_value_size_none"),
  validators.ConflictingObjectAttributes("max_query_size_exceeds",
    "max_query_size_none"),
  validators.ConflictingObjectAttributes("max_request_line_size_exceeds",
    "max_request_line_size_none"),
  validators.ConflictingObjectAttributes("max_request_size_exceeds",
    "max_request_size_none"),
  validators.ConflictingObjectAttributes("max_url_size_exceeds",
    "max_url_size_none")}
```

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

Terraform syntax:

```terraform
request_constraints {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-request_constraints--max_cookie_count_exceeds"></a>

### max_cookie_count_exceeds property

Type: `"number"`. Optional.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_cookie_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_count_none/): complete subsection reference.

<a id="schema-request_constraints--max_cookie_key_size_exceeds"></a>

### max_cookie_key_size_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_cookie_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_key_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_cookie_value_size_exceeds"></a>

### max_cookie_value_size_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32768),
}
```

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

- [max_cookie_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_value_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_header_count_exceeds"></a>

### max_header_count_exceeds property

Type: `"number"`. Optional.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 40),
}
```

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

- [max_header_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_count_none/): complete subsection reference.

<a id="schema-request_constraints--max_header_key_size_exceeds"></a>

### max_header_key_size_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_header_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_key_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_header_value_size_exceeds"></a>

### max_header_value_size_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 64000),
}
```

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

- [max_header_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_value_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_parameter_count_exceeds"></a>

### max_parameter_count_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_parameter_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_count_none/): complete subsection reference.

<a id="schema-request_constraints--max_parameter_name_size_exceeds"></a>

### max_parameter_name_size_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_parameter_name_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_name_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_parameter_value_size_exceeds"></a>

### max_parameter_value_size_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1073741824),
}
```

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

- [max_parameter_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_value_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_query_size_exceeds"></a>

### max_query_size_exceeds property

Type: `"number"`. Optional.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60000),
}
```

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

- [max_query_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_query_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_request_line_size_exceeds"></a>

### max_request_line_size_exceeds property

Type: `"number"`. Optional.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

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

- [max_request_line_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_request_line_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_request_size_exceeds"></a>

### max_request_size_exceeds property

Type: `"number"`. Optional.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

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

- [max_request_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_request_size_none/): complete subsection reference.

<a id="schema-request_constraints--max_url_size_exceeds"></a>

### max_url_size_exceeds property

Type: `"number"`. Optional.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 128000),
}
```

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

- [max_url_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_url_size_none/): complete subsection reference.

## Next pages

- [request_constraints.max_cookie_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_count_none/)
- [request_constraints.max_cookie_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_key_size_none/)
- [request_constraints.max_cookie_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_cookie_value_size_none/)
- [request_constraints.max_header_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_count_none/)
- [request_constraints.max_header_key_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_key_size_none/)
- [request_constraints.max_header_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_header_value_size_none/)
- [request_constraints.max_parameter_count_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_count_none/)
- [request_constraints.max_parameter_name_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_name_size_none/)
- [request_constraints.max_parameter_value_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_parameter_value_size_none/)
- [request_constraints.max_query_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_query_size_none/)
- [request_constraints.max_request_line_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_request_line_size_none/)
- [request_constraints.max_request_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_request_size_none/)
- [request_constraints.max_url_size_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/max_url_size_none/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
