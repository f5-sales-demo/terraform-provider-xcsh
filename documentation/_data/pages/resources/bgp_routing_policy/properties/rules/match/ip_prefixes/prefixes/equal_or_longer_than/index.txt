---
page_title: "rules.match.ip_prefixes.prefixes.equal_or_longer_than"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules match ip prefixes prefixes equal or longer than"], "body_bytes": 1841, "body_sha256": "sha256:2b9cc99ed892d91148b58d9b21fe0e351710158684fcab0224de5ee68503f785", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes:equal_or_longer_than", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:ip_prefixes:prefixes", "path": "documentation/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/equal_or_longer_than/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2230123001022111-2113010223032100-2100121322221021-0132123322121303-0330013012310332-0202131023321012-2220031212310130-3321021123002301", "registry_path": "docs/guides/resources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "match", "ip_prefixes", "prefixes", "equal_or_longer_than"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/equal_or_longer_than/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.match.ip_prefixes.prefixes.equal_or_longer_than

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/)
- [rules.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/)
- [rules.match.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/)
- [rules.match.ip_prefixes.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/)
- rules.match.ip_prefixes.prefixes.equal_or_longer_than

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for equal or longer than.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
equal_or_longer_than = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.match.ip_prefixes.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/rules/match/ip_prefixes/prefixes/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
