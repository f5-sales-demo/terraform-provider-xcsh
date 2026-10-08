---
page_title: "cache_rules.eligible_for_cache"
subcategory: ""
description: "List of OPTIONS for Cache Action."
xcsh_docs: {"aliases": ["cache rules eligible for cache"], "body_bytes": 1438, "body_sha256": "sha256:bfa2e9a1e124a2ad5052868bddbefd972bd1dc50e89af36af9f7c2fe1b60b726", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache"], "schema_version": 1, "sections": [{"aliases": ["cache rules eligible for cache scheme proxy host request uri"], "anchor": "section", "description": "Cache TTL Enable Values.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_request_uri"], "syntax": "attribute", "type": "object"}, {"aliases": ["cache rules eligible for cache scheme proxy host uri"], "anchor": "section", "description": "Cache TTL Enable Values.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of OPTIONS for Cache Action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.eligible_for_cache

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- cache_rules.eligible_for_cache

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for eligible for cache.

Additional upstream details:

List of OPTIONS for Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-eligible_for_cache": "[\"scheme_proxy_host_request_uri\",\"scheme_proxy_host_uri\"]"
}
```

## Direct properties

- [scheme_proxy_host_request_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/): complete subsection reference.

- [scheme_proxy_host_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/): complete subsection reference.
