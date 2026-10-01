---
page_title: "local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name"
subcategory: ""
description: "local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2135, "body_sha256": "sha256:092744dab9549a1e6ea822edc75c3306043b8d2321aa3d3511bfe21f4f8f7337", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "path": "docs/guides/resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy", "node_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/node_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](resources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

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

<a id="schema-local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name--node"></a>

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

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
