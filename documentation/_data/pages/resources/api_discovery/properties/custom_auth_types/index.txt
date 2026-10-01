---
page_title: "custom_auth_types"
subcategory: ""
description: "custom_auth_types for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 4103, "body_sha256": "sha256:e6275c12b5aaf531fbb65beaba4e8aee9ca45deafcba88ff05fc126c800c41d9", "child_ids": [], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "parent_id": "xcsh-docs:resources:api_discovery:reference", "path": "documentation/resources/api_discovery/properties/custom_auth_types/index.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["custom_auth_types"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/custom_auth_types/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_auth_types for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_auth_types

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- custom_auth_types

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Select your custom authentication types to be detected in the API discovery. Defaults to \`\[\]\`.
Server applies default when omitted.

Upstream description:

Select your custom authentication types to be detected in the API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("parameter_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10"
  }
}
```

Terraform syntax:

```terraform
custom_auth_types {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_auth_types--parameter_name"></a>

### parameter_name property

Type: `"string"`. Optional.

Parameter Name. The authentication parameter name.

Upstream description:

The authentication parameter name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.pattern": "^[!#$%&'*+\\\\-.^_`|~0-9A-Za-z]+$"
  }
}
```

<a id="schema-custom_auth_types--parameter_type"></a>

### parameter_type property

Type: `"string"`. Optional.

\[Enum: QUERY\_PARAMETER|HEADER|COOKIE\] Enumeration for authentication parameter types. Possible
values are \`QUERY\_PARAMETER\`, \`HEADER\`, \`COOKIE\`. Defaults to \`QUERY\_PARAMETER\`.

Upstream description:

Enumeration for authentication parameter types.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("QUERY_PARAMETER",
    "HEADER",
    "COOKIE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "QUERY_PARAMETER",
  "enum": [
    "QUERY_PARAMETER",
    "HEADER",
    "COOKIE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/properties/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_discovery/)
