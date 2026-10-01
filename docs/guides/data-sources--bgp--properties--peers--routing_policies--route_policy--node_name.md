---
page_title: "peers.routing_policies.route_policy.node_name"
subcategory: ""
description: "peers.routing_policies.route_policy.node_name for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1391, "body_sha256": "sha256:3ebe0b6013b1e1102d5c56646a0b4c7dba3daec6ac50991ecb9ddb3937a000ea", "canonical_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:node_name", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:node_name", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy", "path": "docs/guides/data-sources--bgp--properties--peers--routing_policies--route_policy--node_name.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy", "node_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/routing_policies/route_policy/node_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.routing_policies.route_policy.node_name for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy.node_name

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Property reference](data-sources--bgp--reference.md)
- [peers](data-sources--bgp--properties--peers.md)
- [peers.routing_policies](data-sources--bgp--properties--peers--routing_policies.md)
- [peers.routing_policies.route_policy](data-sources--bgp--properties--peers--routing_policies--route_policy.md)
- peers.routing_policies.route_policy.node_name

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-peers--routing_policies--route_policy--node_name--node"></a>

### node property

Type: `["list", "string"]`. Computed.

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

- [peers.routing_policies.route_policy](data-sources--bgp--properties--peers--routing_policies--route_policy.md)
- [xcsh_bgp](../data-sources/bgp.md)
