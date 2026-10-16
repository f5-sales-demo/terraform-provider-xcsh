---
page_title: "rules.egress_rules.all_udp_traffic"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules egress rules all udp traffic"], "body_bytes": 1155, "body_sha256": "sha256:ae2b7a08cfd0adaeab22a6fc6a161c205766b2ce9fd880e20ac439f3ad548524", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules:all_udp_traffic", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules", "path": "documentation/resources/network_policy/properties/rules/egress_rules/all_udp_traffic/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3122203212023100-0323303031302322-1110302012130131-0303330321001013-0333212110323233-0233000301011311-3103202300223330-1310131031213330", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "egress_rules", "all_udp_traffic"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/egress_rules/all_udp_traffic/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["network_policyCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.all_udp_traffic

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/)
- rules.egress_rules.all_udp_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

Additional upstream details:

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
all_udp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.
