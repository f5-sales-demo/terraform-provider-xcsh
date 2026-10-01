---
page_title: "local_control_plane.bgp_config.peers.routing_policies.route_policy"
subcategory: ""
description: "local_control_plane.bgp_config.peers.routing_policies.route_policy for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 4174, "body_sha256": "sha256:f041aa91c9ba8bae32ea351e667f1dadee6963d9bb41e02bc796e87a8adb3a7c", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:all_nodes", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:inbound", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:node_name", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:object_refs", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy:outbound"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies:route_policy", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:routing_policies", "path": "docs/guides/resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "routing_policies", "route_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/routing_policies/route_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.routing_policies.route_policy for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.routing_policies.route_policy

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [local_control_plane](resources--voltstack_site--properties--local_control_plane.md)
- [local_control_plane.bgp_config](resources--voltstack_site--properties--local_control_plane--bgp_config.md)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--properties--local_control_plane--bgp_config--peers.md)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md)
- local_control_plane.bgp_config.peers.routing_policies.route_policy

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("object_refs"),
  validators.ConflictingListObjectAttributes("all_nodes",
    "node_name"),
  validators.ConflictingListObjectAttributes("inbound",
    "outbound")}
```

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

Terraform syntax:

```terraform
route_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_nodes](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--all_nodes.md): complete subsection reference.

- [inbound](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--inbound.md): complete subsection reference.

- [node_name](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name.md): complete subsection reference.

- [object_refs](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--object_refs.md): complete subsection reference.

- [outbound](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--outbound.md): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--all_nodes.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--inbound.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--node_name.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--object_refs.md)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies--route_policy--outbound.md)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--properties--local_control_plane--bgp_config--peers--routing_policies.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
