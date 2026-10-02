---
page_title: "custom_auth_types"
subcategory: ""
description: "Select your custom authentication types to be detected in the API discovery."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "custom auth types"], "body_bytes": 4103, "body_sha256": "sha256:e6275c12b5aaf531fbb65beaba4e8aee9ca45deafcba88ff05fc126c800c41d9", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "parent_id": "xcsh-docs:resources:api_discovery:reference", "path": "documentation/resources/api_discovery/properties/custom_auth_types/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2212121133030123-3313103211030133-1332101312033102-0023310100201330-1131130101331031-0031112010301020-0220030121333230-3300021031132021", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [{"anchor": "schema-custom_auth_types--parameter_name", "enforcement": "provider-schema", "group": "custom_auth_types:RequiredListObjectAttributes:parameter_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_auth_types"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential setup", "credentials", "parameter name"], "anchor": "schema-custom_auth_types--parameter_name", "description": "The authentication parameter name.", "document_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_auth_types", "parameter_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials", "parameter type"], "anchor": "schema-custom_auth_types--parameter_type", "description": "Enumeration for authentication parameter types.", "document_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_auth_types", "parameter_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/custom_auth_types/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select your custom authentication types to be detected in the API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
