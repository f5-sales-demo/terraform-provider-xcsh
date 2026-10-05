---
page_title: "egress_rules.ip_prefix_set"
subcategory: ""
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["egress rules ip prefix set"], "body_bytes": 1528, "body_sha256": "sha256:af9b38d4307cadca068ef8c846b9d676238c819bba309e0b349a27ebcafb88f3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set", "parent_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "path": "documentation/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323", "registry_path": "docs/guides/data-sources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["egress_rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["egress rules ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set:ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["egress_rules", "ip_prefix_set", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
- [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/)
- egress_rules.ip_prefix_set

<a id="section"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/): complete subsection reference.

## Next pages

- [egress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/ref/)
- [egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/egress_rules/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
