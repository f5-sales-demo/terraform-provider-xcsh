---
page_title: "peers.routing_policies.route_policy.node_name"
subcategory: ""
description: "List of nodes on which BGP routing policy has to be applied."
xcsh_docs: {"aliases": ["peers routing policies route policy node name"], "body_bytes": 1842, "body_sha256": "sha256:1334929034bef9cf3b2f9562e2cee7a0fa557c1218a5448642535d894d664b47", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "parent_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "path": "documentation/resources/bgp/properties/peers/routing_policies/route_policy/node_name/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0223023301201233-0013302203332313-2100312202110011-3322232232303000-1022321200012330-1323321022220112-0122032220100333-0313033203333220", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy", "node_name"], "schema_version": 1, "sections": [{"aliases": ["peers routing policies route policy node name node"], "anchor": "schema-peers--routing_policies--route_policy--node_name--node", "description": "Select BGP Session on which policy will be applied.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "node_name", "node"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/route_policy/node_name/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of nodes on which BGP routing policy has to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy.node_name

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/)
- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/)
- peers.routing_policies.route_policy.node_name

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of nodes on which BGP routing policy has to be applied.

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
node_name {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-peers--routing_policies--route_policy--node_name--node"></a>

### node property

Type: `["list", "string"]`. Optional.

Select BGP Session on which policy will be applied.

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

## Next pages

- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
