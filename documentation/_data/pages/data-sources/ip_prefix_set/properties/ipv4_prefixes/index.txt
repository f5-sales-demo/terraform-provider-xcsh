---
page_title: "ipv4_prefixes"
subcategory: ""
description: "List of IPv4 prefixes with description."
xcsh_docs: {"aliases": ["ipv4 prefixes"], "body_bytes": 2356, "body_sha256": "sha256:126107b7d9114d770499ddad691c6edfa6ce4c965ad41b5e8833101fd1f63678", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "parent_id": "xcsh-docs:data-sources:ip_prefix_set:reference", "path": "documentation/data-sources/ip_prefix_set/properties/ipv4_prefixes/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2230202101210010-2331000201030332-0321131110201103-3123221101201312-1323101201322322-0023022032221311-2000312233223302-0121112122331233", "registry_path": "docs/guides/data-sources--ip_prefix_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipv4_prefixes"], "schema_version": 1, "sections": [{"aliases": ["ipv4 prefixes description spec"], "anchor": "schema-ipv4_prefixes--description_spec", "description": "Description. Human-readable description text", "document_id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4_prefixes", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipv4 prefixes ipv4 prefix"], "anchor": "schema-ipv4_prefixes--ipv4_prefix", "description": "IP address configuration", "document_id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4_prefixes", "ipv4_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/properties/ipv4_prefixes/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of IPv4 prefixes with description.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
