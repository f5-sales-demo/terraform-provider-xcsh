---
page_title: "cache_rules.eligible_for_cache"
subcategory: ""
description: "List of OPTIONS for Cache Action."
xcsh_docs: {"aliases": ["cache rules eligible for cache"], "body_bytes": 2112, "body_sha256": "sha256:1a2b54a08442b69ec3fcf2c841daaf97e56f8f37aae01487d8cf3861b879cdfa", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache"], "schema_version": 1, "sections": [{"aliases": ["scheme proxy host request uri"], "anchor": "section", "description": "Cache TTL Enable Values.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_request_uri"], "syntax": "attribute", "type": "object"}, {"aliases": ["scheme proxy host uri"], "anchor": "section", "description": "Cache TTL Enable Values.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of OPTIONS for Cache Action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

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

## Next pages

- [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/)
- [cache_rules.eligible_for_cache.scheme_proxy_host_uri](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
