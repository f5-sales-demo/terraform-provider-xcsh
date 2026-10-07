---
page_title: "ipv4_prefixes"
subcategory: ""
description: "List of IPv4 prefixes with description."
xcsh_docs: {"aliases": ["ipv4 prefixes"], "body_bytes": 2356, "body_sha256": "sha256:587b86dca1bf264a39cebcc364f2179a54b3d66d8678a9138801b6564abffef0", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "parent_id": "xcsh-docs:data-sources:ip_prefix_set:reference", "path": "documentation/data-sources/ip_prefix_set/properties/ipv4_prefixes/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2230202101210010-2331000201030332-0321131110201103-3123221101201312-1323101201322322-0023022032221311-2000312233223302-0121112122331233", "registry_path": "docs/guides/data-sources--ip_prefix_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipv4_prefixes"], "schema_version": 1, "sections": [{"aliases": ["ipv4 prefixes description spec"], "anchor": "schema-ipv4_prefixes--description_spec", "description": "Description. Human-readable description text", "document_id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4_prefixes", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipv4 prefixes ipv4 prefix"], "anchor": "schema-ipv4_prefixes--ipv4_prefix", "description": "IP address configuration", "document_id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4_prefixes", "ipv4_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/properties/ipv4_prefixes/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "List of IPv4 prefixes with description.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipv4_prefixes

Breadcrumbs:

- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/properties/)
- ipv4_prefixes

<a id="section"></a>

Type: `"list"`. Computed.

IPv4 Prefixes. List of IPv4 prefixes with description.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-ipv4_prefixes--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Human-readable description text

<a id="schema-ipv4_prefixes--ipv4_prefix"></a>

### ipv4_prefix property

Type: `"string"`. Computed.

IPv4 Prefix. IP address configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```
