---
page_title: "local_control_plane.bgp_config.peers.routing_policies.route_policy"
subcategory: ""
description: "local_control_plane.bgp_config.peers.routing_policies.route_policy for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3707, "body_sha256": "sha256:2dad156d731c2bffafbbca49abfa95f08f285ec2a58589592b5274a928f42d27", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:all_nodes", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:inbound", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:object_refs", "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:outbound"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "path": "docs/guides/data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.routing_policies.route_policy for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_control_plane.bgp_config.peers.routing_policies.route_policy

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [local_control_plane](data-sources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md)
- local_control_plane.bgp_config.peers.routing_policies.route_policy

<a id="section"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

## Direct properties

- [all_nodes](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--all_nodes.md): complete subsection reference.

- [inbound](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--inbound.md): complete subsection reference.

- [node_name](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name.md): complete subsection reference.

- [object_refs](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--object_refs.md): complete subsection reference.

- [outbound](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--outbound.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--all_nodes.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--inbound.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--object_refs.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--outbound.md)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
