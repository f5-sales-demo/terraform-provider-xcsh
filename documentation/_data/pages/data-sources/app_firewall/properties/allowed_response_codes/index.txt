---
page_title: "allowed_response_codes"
subcategory: "Security"
description: "List of HTTP response status codes that are allowed."
xcsh_docs: {"aliases": ["allowed response codes"], "body_bytes": 2104, "body_sha256": "sha256:944f5dc069e1c585815e291c8f2d431b7944ad947a8638b799bbcebd56622a2b", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:allowed_response_codes", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/allowed_response_codes/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2120023301001102-2103000021001323-3230113311223023-3220101220122000-2300110311220000-2121131310320100-2130312102313212-2200031122131102", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowed_response_codes"], "schema_version": 1, "sections": [{"aliases": ["allowed response codes response code"], "anchor": "schema-allowed_response_codes--response_code", "description": "List of HTTP response status codes that are allowed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:allowed_response_codes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_response_codes", "response_code"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/allowed_response_codes/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of HTTP response status codes that are allowed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["app_firewallCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowed_response_codes

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- allowed_response_codes

<a id="section"></a>

Type: `"single"`. Computed.

List of HTTP response status codes that are allowed.

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

<a id="schema-allowed_response_codes--response_code"></a>

### response_code property

Type: `["list", "number"]`. Computed.

List of HTTP response status codes that are allowed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 48,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
