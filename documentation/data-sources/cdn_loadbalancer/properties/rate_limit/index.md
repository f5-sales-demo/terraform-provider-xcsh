---
page_title: "rate_limit"
subcategory: "Load Balancing"
description: "RateLimitConfigType."
xcsh_docs: {"aliases": ["rate limit"], "body_bytes": 3090, "body_sha256": "sha256:6e1e0475a7880b662bb753157d3ac93359272ab000b2e255555f5aaa3c71b892", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:ip_allowed_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:no_policies", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:policies", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/rate_limit/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit"], "schema_version": 1, "sections": [{"aliases": ["custom ip allowed list"], "anchor": "section", "description": "IP Allowed list using existing ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "custom_ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip allowed list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:ip_allowed_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["no ip allowed list"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "no_ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["no policies"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:no_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "no_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["policies"], "anchor": "section", "description": "List of rate limiter policies to be applied.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["rate limiter"], "anchor": "section", "description": "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/rate_limit/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "RateLimitConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- rate_limit

<a id="section"></a>

Type: `"single"`. Computed.

RateLimitConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

## Direct properties

- [custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/): complete subsection reference.

- [ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/ip_allowed_list/): complete subsection reference.

- [no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/no_ip_allowed_list/): complete subsection reference.

- [no_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/no_policies/): complete subsection reference.

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/policies/): complete subsection reference.

- [rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/): complete subsection reference.

## Next pages

- [rate_limit.custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/)
- [rate_limit.ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/ip_allowed_list/)
- [rate_limit.no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/no_ip_allowed_list/)
- [rate_limit.no_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/no_policies/)
- [rate_limit.policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/policies/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
