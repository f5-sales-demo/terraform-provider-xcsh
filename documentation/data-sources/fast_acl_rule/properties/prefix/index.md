---
page_title: "prefix"
subcategory: ""
description: "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes."
xcsh_docs: {"aliases": ["prefix"], "body_bytes": 1697, "body_sha256": "sha256:eb7f8b0765699740d0be2923dcfd3b5a5676838ccc9bbb6f7a3c39d22f825327", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:properties:prefix", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:reference", "path": "documentation/data-sources/fast_acl_rule/properties/prefix/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2002033300333000-1223112020311322-1210001022020221-0123333022303330-0030013300310300-3210021023233013-1021201300130211-2203221310300103", "registry_path": "docs/guides/data-sources--fast_acl_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["prefix"], "schema_version": 1, "sections": [{"aliases": ["prefix prefix"], "anchor": "schema-prefix--prefix", "description": "IP Address prefix in string format. String must contain both prefix and prefix-length.", "document_id": "xcsh-docs:data-sources:fast_acl_rule:properties:prefix", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["prefix", "prefix"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/properties/prefix/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# prefix

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/)
- prefix

<a id="section"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="schema-prefix--prefix"></a>

### prefix property

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-17T10:10:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```
