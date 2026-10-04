---
page_title: "custom_auth_types"
subcategory: ""
description: "Select your custom authentication types to be detected in the API discovery."
xcsh_docs: {"aliases": ["custom auth types"], "body_bytes": 4103, "body_sha256": "sha256:d554eb9238b733ac96c5c1fe54977430306af5fe9663deedb0c12f4b47178a6c", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "parent_id": "xcsh-docs:resources:api_discovery:reference", "path": "documentation/resources/api_discovery/properties/custom_auth_types/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2212121133030123-3313103211030133-1332101312033102-0023310100201330-1131130101331031-0031112010301020-0220030121333230-3300021031132021", "registry_path": "docs/guides/resources--api_discovery--reference--group-001.md", "relationships": [{"anchor": "schema-custom_auth_types--parameter_name", "enforcement": "provider-schema", "group": "custom_auth_types:RequiredListObjectAttributes:parameter_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_auth_types"], "schema_version": 1, "sections": [{"aliases": ["custom auth types parameter name"], "anchor": "schema-custom_auth_types--parameter_name", "description": "The authentication parameter name.", "document_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_auth_types", "parameter_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom auth types parameter type"], "anchor": "schema-custom_auth_types--parameter_type", "description": "Enumeration for authentication parameter types.", "document_id": "xcsh-docs:resources:api_discovery:properties:custom_auth_types", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_auth_types", "parameter_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/properties/custom_auth_types/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Select your custom authentication types to be detected in the API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
