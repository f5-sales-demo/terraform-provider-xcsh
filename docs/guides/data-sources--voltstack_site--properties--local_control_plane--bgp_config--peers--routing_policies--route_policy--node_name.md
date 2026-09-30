---
page_title: "local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name"
subcategory: ""
description: "local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1945, "body_sha256": "sha256:81711f4efbff7a967f1309d3eaa1ea3ef1b1b64c2f21e4bd6f1be6b5c4e4628f", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "child_ids": [], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "path": "docs/guides/data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "node_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/node_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [local_control_plane](data-sources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

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

<a id="schema-local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name--node"></a>

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

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
