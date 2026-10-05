---
page_title: "rules.criteria.tcp"
subcategory: ""
description: "Action to apply on the packet if the NAT rule is applied."
xcsh_docs: {"aliases": ["rules criteria tcp"], "body_bytes": 1846, "body_sha256": "sha256:1fdb42d3cd9eff7893b64474661bcd668f577712579bffa0799285d207282737", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:destination_port", "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria", "path": "documentation/data-sources/nat_policy/properties/rules/criteria/tcp/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0302030212022102-0022203320220102-2301303202022232-1233212230311300-3123132101033331-0113312012012101-1200202130101220-3132111313230213", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria", "tcp"], "schema_version": 1, "sections": [{"aliases": ["rules criteria tcp destination port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:destination_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "criteria", "tcp", "destination_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria tcp source port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:criteria:tcp:source_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "criteria", "tcp", "source_port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/criteria/tcp/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Action to apply on the packet if the NAT rule is applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.tcp

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/)
- rules.criteria.tcp

<a id="section"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

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

- [destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/destination_port/): complete subsection reference.

- [source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/source_port/): complete subsection reference.

## Next pages

- [rules.criteria.tcp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/destination_port/)
- [rules.criteria.tcp.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/tcp/source_port/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/criteria/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
