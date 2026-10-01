---
page_title: "waf_exclusion.waf_exclusion_inline_rules"
subcategory: "Load Balancing"
description: "waf_exclusion.waf_exclusion_inline_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1257, "body_sha256": "sha256:54dc866fce7ae00d58bd02d589391b897c1bd42ef7f1d4d6d2130129b1279185", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:waf_exclusion", "path": "docs/guides/data-sources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion.waf_exclusion_inline_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [waf_exclusion](data-sources--http_loadbalancer--properties--waf_exclusion.md)
- waf_exclusion.waf_exclusion_inline_rules

<a id="section"></a>

Type: `"single"`. Computed.

List of WAF exclusion rules that will be applied inline.

Upstream description:

A list of WAF exclusion rules that will be applied inline.

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

- [rules](data-sources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules.md): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules.md)
- [waf_exclusion](data-sources--http_loadbalancer--properties--waf_exclusion.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
