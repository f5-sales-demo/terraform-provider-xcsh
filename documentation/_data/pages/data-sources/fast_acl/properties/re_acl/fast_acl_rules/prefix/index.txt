---
page_title: "re_acl.fast_acl_rules.prefix"
subcategory: ""
description: "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes."
xcsh_docs: {"aliases": ["re acl fast acl rules prefix"], "body_bytes": 1998, "body_sha256": "sha256:62a4ec29a32785e7f15bf55fe657422d28598b35f559fe95b50b1a6d9d43224c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:prefix", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules", "path": "documentation/data-sources/fast_acl/properties/re_acl/fast_acl_rules/prefix/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2333121103020211-1132230120300010-3220033202123021-3220113120112312-0301230122021222-3220313031300330-3300303313202000-3130310002001320", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "prefix"], "schema_version": 1, "sections": [{"aliases": ["re acl fast acl rules prefix prefix"], "anchor": "schema-re_acl--fast_acl_rules--prefix--prefix", "description": "IP Address prefix in string format. String must contain both prefix and prefix-length.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:prefix", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "prefix", "prefix"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/fast_acl_rules/prefix/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["fast_aclCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules.prefix

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/re_acl/fast_acl_rules/)
- re_acl.fast_acl_rules.prefix

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

<a id="schema-re_acl--fast_acl_rules--prefix--prefix"></a>

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```
