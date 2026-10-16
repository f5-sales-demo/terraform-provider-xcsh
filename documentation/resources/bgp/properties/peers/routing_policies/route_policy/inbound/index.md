---
page_title: "peers.routing_policies.route_policy.inbound"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["peers routing policies route policy inbound"], "body_bytes": 1251, "body_sha256": "sha256:79e20953ef097a2257bc94fec5a526a84bf373d5258e4dbaf7f78a5795148ba8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:inbound", "parent_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "path": "documentation/resources/bgp/properties/peers/routing_policies/route_policy/inbound/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0331112133300211-0220113302003321-0003030201203110-3331322233022123-2313003003223123-3331312211120012-2213002203022013-2002322121000333", "registry_path": "docs/guides/resources--bgp--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy", "inbound"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/route_policy/inbound/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["bgpCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy.inbound

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/)
- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/)
- peers.routing_policies.route_policy.inbound

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
inbound = {}
```

This is an empty object or choice marker. It has no direct properties.
