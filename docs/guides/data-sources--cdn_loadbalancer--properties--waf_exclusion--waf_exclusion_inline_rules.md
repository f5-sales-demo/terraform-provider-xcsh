---
page_title: "waf_exclusion.waf_exclusion_inline_rules"
subcategory: "Load Balancing"
description: "waf_exclusion.waf_exclusion_inline_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1248, "body_sha256": "sha256:9b619db7bbb7d58f0eadf8ab82d9a0390a6cd748a4105b8029405c090bc5ce45", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion.waf_exclusion_inline_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [waf_exclusion](data-sources--cdn_loadbalancer--properties--waf_exclusion.md)
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

- [rules](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules.md): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules--rules.md)
- [waf_exclusion](data-sources--cdn_loadbalancer--properties--waf_exclusion.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
