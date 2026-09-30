---
page_title: "waf_exclusion"
subcategory: "Load Balancing"
description: "waf_exclusion for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1278, "body_sha256": "sha256:43e1a139ba57005d39137e8d5cf3153c44fce2873088fda7a1e55e2eee3f5e3e", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_policy"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--waf_exclusion.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_exclusion"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/waf_exclusion/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_exclusion for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# waf_exclusion

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- waf_exclusion

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for waf exclusion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

## Direct properties

- [waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md): complete subsection reference.

- [waf_exclusion_policy](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_policy.md): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_inline_rules.md)
- [waf_exclusion.waf_exclusion_policy](data-sources--cdn_loadbalancer--properties--waf_exclusion--waf_exclusion_policy.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
