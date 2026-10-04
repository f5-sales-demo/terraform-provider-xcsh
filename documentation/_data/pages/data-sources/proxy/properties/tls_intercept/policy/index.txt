---
page_title: "tls_intercept.policy"
subcategory: ""
description: "Policy to enable or disable TLS interception."
xcsh_docs: {"aliases": ["tls intercept policy"], "body_bytes": 1362, "body_sha256": "sha256:30b9eaa7ea7dbb17a10d53b092024a92185ef77f6ca3c8258e2a512a4f9c9e04", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy", "parent_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept", "path": "documentation/data-sources/proxy/properties/tls_intercept/policy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100", "registry_path": "docs/guides/data-sources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "policy"], "schema_version": 1, "sections": [{"aliases": ["tls intercept policy interception rules"], "anchor": "section", "description": "List of ordered rules to enable or disable for TLS interception.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["tls_intercept", "policy", "interception_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/policy/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Policy to enable or disable TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/)
- tls_intercept.policy

<a id="section"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

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

- [interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/interception_rules/): complete subsection reference.

## Next pages

- [tls_intercept.policy.interception_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/interception_rules/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
