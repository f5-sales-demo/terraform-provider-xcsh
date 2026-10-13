---
page_title: "http_access"
subcategory: ""
description: "Configuration parameter for http access."
xcsh_docs: {"aliases": ["http access"], "body_bytes": 1538, "body_sha256": "sha256:d06e401cf1f6079bff3f2c491e52041791cad4b288dbb4720d79c08fe3f98f9a", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:crl:properties:http_access", "parent_id": "xcsh-docs:data-sources:crl:reference", "path": "documentation/data-sources/crl/properties/http_access/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:1b9ea59a790d92083efb68ff1abebe57ebbb3f7b716b91cdb7b0416134676174", "provider_type": "data-sources", "registry_anchor": "canonical-1022222013321321-2111220010000022-2232313002220321-3121302011213012-0210002333232201-0302202003231231-1033320101311130-3211121130113211", "registry_path": "docs/guides/data-sources--crl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_access"], "schema_version": 1, "sections": [{"aliases": ["http access path"], "anchor": "schema-http_access--path", "description": "CRL file location.", "document_id": "xcsh-docs:data-sources:crl:properties:http_access", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_access", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/properties/http_access/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Configuration parameter for http access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["crlCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_access

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/properties/)
- http_access

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-http_access--path"></a>

### path property

Type: `"string"`. Computed.

CRL File path. CRL file location.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-13T13:24:52+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
