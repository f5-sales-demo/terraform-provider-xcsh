---
page_title: "http_access"
subcategory: ""
description: "Configuration parameter for http access."
xcsh_docs: {"aliases": ["http access"], "body_bytes": 1538, "body_sha256": "sha256:dc7aab4cd0e58f3ee6186bc7744488c62902629418ad642594cd2036d8ef38a2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:crl:properties:http_access", "parent_id": "xcsh-docs:data-sources:crl:reference", "path": "documentation/data-sources/crl/properties/http_access/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1022222013321321-2111220010000022-2232313002220321-3121302011213012-0210002333232201-0302202003231231-1033320101311130-3211121130113211", "registry_path": "docs/guides/data-sources--crl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_access"], "schema_version": 1, "sections": [{"aliases": ["http access path"], "anchor": "schema-http_access--path", "description": "CRL file location.", "document_id": "xcsh-docs:data-sources:crl:properties:http_access", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_access", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/properties/http_access/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Configuration parameter for http access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["crlCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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
