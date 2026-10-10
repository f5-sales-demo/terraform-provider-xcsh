---
page_title: "site_acl.fast_acl_rules.prefix"
subcategory: ""
description: "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes."
xcsh_docs: {"aliases": ["site acl fast acl rules prefix"], "body_bytes": 2012, "body_sha256": "sha256:9a04c6c6de7f02f12fbf73f90be7a506e06fe0d7d96e8612e993b2a6ab8606b3", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:prefix", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules", "path": "documentation/data-sources/fast_acl/properties/site_acl/fast_acl_rules/prefix/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3032222320121131-1200233103132111-1211223013221331-3032033230313000-3302333110003003-2131103313133032-2332113111233312-3100002230232033", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "prefix"], "schema_version": 1, "sections": [{"aliases": ["site acl fast acl rules prefix prefix"], "anchor": "schema-site_acl--fast_acl_rules--prefix--prefix", "description": "IP Address prefix in string format. String must contain both prefix and prefix-length.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:prefix", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "prefix", "prefix"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/site_acl/fast_acl_rules/prefix/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fast_aclCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.prefix

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/)
- site_acl.fast_acl_rules.prefix

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

<a id="schema-site_acl--fast_acl_rules--prefix--prefix"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
