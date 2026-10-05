---
page_title: "rules.criteria.tcp.destination_port.no_port_match"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules criteria tcp destination port no port match"], "body_bytes": 1723, "body_sha256": "sha256:56766c9ca37858322460f56cbb9f4832f794791984b2f99fd7a22ee1063c80ab", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port:no_port_match", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "path": "documentation/resources/nat_policy/properties/rules/criteria/tcp/destination_port/no_port_match/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0023103200122233-2030011202213110-0331310103333012-0121301122133133-1313110212003031-2003012021111102-0311333212121033-2200022221020302", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria", "tcp", "destination_port", "no_port_match"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/tcp/destination_port/no_port_match/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.tcp.destination_port.no_port_match

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/)
- [rules.criteria.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/)
- [rules.criteria.tcp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/)
- rules.criteria.tcp.destination_port.no_port_match

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
no_port_match = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.criteria.tcp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
