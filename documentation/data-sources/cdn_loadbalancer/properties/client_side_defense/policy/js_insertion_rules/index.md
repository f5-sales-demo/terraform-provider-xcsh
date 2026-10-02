---
page_title: "client_side_defense.policy.js_insertion_rules"
subcategory: "Load Balancing"
description: "This defines custom JavaScript insertion rules for Client-Side Defense Policy."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules"], "body_bytes": 2289, "body_sha256": "sha256:6d2cf7b6d3053f5914df6125464d58cbf0a97dc1b4663d7e659ee70f1e7d0d2a", "capabilities": ["cdn", "security.client-side-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy", "path": "documentation/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules"], "schema_version": 1, "sections": [{"aliases": ["exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules"], "anchor": "section", "description": "Required list of pages to insert Client-Side Defense client JavaScript.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines custom JavaScript insertion rules for Client-Side Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/)
- client_side_defense.policy.js_insertion_rules

<a id="section"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

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

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/): complete subsection reference.

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/)
- [client_side_defense.policy.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/client_side_defense/policy/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
