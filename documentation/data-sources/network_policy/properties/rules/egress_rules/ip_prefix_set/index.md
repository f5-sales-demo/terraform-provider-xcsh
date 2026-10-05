---
page_title: "rules.egress_rules.ip_prefix_set"
subcategory: "Security"
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["rules egress rules ip prefix set"], "body_bytes": 1650, "body_sha256": "sha256:46560244ec79fc74a9df67f3b658c6cc1f5d91a5aa56be0118c47fbc77656ba8", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:ip_prefix_set", "parent_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "path": "documentation/data-sources/network_policy/properties/rules/egress_rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1312312110232031-1203301113033300-2032231212221031-2013232311020211-0320030010121223-3332230221103331-1111333303323003-3320220200210220", "registry_path": "docs/guides/data-sources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "egress_rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["rules egress rules ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:ip_prefix_set:ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "egress_rules", "ip_prefix_set", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/rules/egress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/)
- rules.egress_rules.ip_prefix_set

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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/): complete subsection reference.

## Next pages

- [rules.egress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
