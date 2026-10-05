---
page_title: "client_side_defense.policy.js_insertion_rules"
subcategory: "Load Balancing"
description: "This defines custom JavaScript insertion rules for Client-Side Defense Policy."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules"], "body_bytes": 2301, "body_sha256": "sha256:76f8b318e20c6a40e7054b1f0db0db17130517c4b59ed3027f56caf67c8ac236", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy", "path": "documentation/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy js insertion rules exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:exclude_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "exclude_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insertion rules rules"], "anchor": "section", "description": "Required list of pages to insert Client-Side Defense client JavaScript.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This defines custom JavaScript insertion rules for Client-Side Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/)
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

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/): complete subsection reference.

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/exclude_list/)
- [client_side_defense.policy.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
