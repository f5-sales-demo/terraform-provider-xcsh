---
page_title: "peers.routing_policies.route_policy.node_name"
subcategory: ""
description: "peers.routing_policies.route_policy.node_name for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1488, "body_sha256": "sha256:b1ad14a14da3bb21c49b76d2f0881a2bb05f4a9585ba925b6f53b94c0a91e3d3", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "parent_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "path": "docs/guides/resources--bgp--properties--peers--routing_policies--route_policy--node_name.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy", "node_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/route_policy/node_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.routing_policies.route_policy.node_name for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy.node_name

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.routing_policies](resources--bgp--properties--peers--routing_policies.md)
- [peers.routing_policies.route_policy](resources--bgp--properties--peers--routing_policies--route_policy.md)
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

- [peers.routing_policies.route_policy](resources--bgp--properties--peers--routing_policies--route_policy.md)
- [xcsh_bgp](../resources/bgp.md)
