---
page_title: "custom_auth_types"
subcategory: ""
description: "Select your custom authentication types to be detected in the API discovery."
xcsh_docs: {"aliases": ["custom auth types"], "body_bytes": 3048, "body_sha256": "sha256:c9d9211f657efad80ea7dc51fb2023fe8191620850ea62120c25f842f9e5371e", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:custom_auth_types", "parent_id": "xcsh-docs:data-sources:api_discovery:reference", "path": "documentation/data-sources/api_discovery/properties/custom_auth_types/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0030011001132310-0130023101132022-3322210030220302-0321012110320203-0313220020332330-0100101001100333-1013101323033200-3133301202233011", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_auth_types"], "schema_version": 1, "sections": [{"aliases": ["custom auth types parameter name"], "anchor": "schema-custom_auth_types--parameter_name", "description": "The authentication parameter name.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:custom_auth_types", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_auth_types", "parameter_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom auth types parameter type"], "anchor": "schema-custom_auth_types--parameter_type", "description": "Enumeration for authentication parameter types.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:custom_auth_types", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_auth_types", "parameter_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/custom_auth_types/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Select your custom authentication types to be detected in the API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_auth_types

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/)
- custom_auth_types

<a id="section"></a>

Type: `"list"`. Computed.

Select your custom authentication types to be detected in the API discovery. Defaults to \`\[\]\`.
Server applies default when omitted.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

## Direct properties

<a id="schema-custom_auth_types--parameter_name"></a>

### parameter_name property

Type: `"string"`. Computed.

Parameter Name. The authentication parameter name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Type: `"string"`. Computed.

\[Enum: QUERY\_PARAMETER|HEADER|COOKIE\] Enumeration for authentication parameter types. Possible
values are \`QUERY\_PARAMETER\`, \`HEADER\`, \`COOKIE\`. Defaults to \`QUERY\_PARAMETER\`.

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
