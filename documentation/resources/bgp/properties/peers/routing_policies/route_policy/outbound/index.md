---
page_title: "peers.routing_policies.route_policy.outbound"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["peers routing policies route policy outbound"], "body_bytes": 1254, "body_sha256": "sha256:596c4744e24980ba6eacbcfdda354873330b4fa81b92f53b509ae559a641c2f3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:outbound", "parent_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "path": "documentation/resources/bgp/properties/peers/routing_policies/route_policy/outbound/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3221333323021330-0323000002011212-1023032220130110-0122031210303123-3212000211131011-0330300333213322-0130121001222101-1102121322001321", "registry_path": "docs/guides/resources--bgp--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy", "outbound"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/route_policy/outbound/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["bgpCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy.outbound

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/)
- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/)
- peers.routing_policies.route_policy.outbound

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
outbound = {}
```

This is an empty object or choice marker. It has no direct properties.
