---
page_title: "local_control_plane.bgp_config.peers.routing_policies"
subcategory: ""
description: "local_control_plane.bgp_config.peers.routing_policies for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1549, "body_sha256": "sha256:d4be1b38e2886686b7361f480fbe8298058762c0553a20123059f87ea01cd9da", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers", "path": "docs/guides/data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.routing_policies for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.routing_policies

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [local_control_plane](data-sources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](data-sources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- local_control_plane.bgp_config.peers.routing_policies

<a id="section"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

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

- [route_policy](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
