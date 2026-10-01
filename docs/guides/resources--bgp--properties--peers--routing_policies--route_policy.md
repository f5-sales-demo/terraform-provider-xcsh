---
page_title: "peers.routing_policies.route_policy"
subcategory: ""
description: "peers.routing_policies.route_policy for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 3040, "body_sha256": "sha256:51c036d69cc3a4a8e3eb222067a4b7a3d849fde9aceff5f35ca2a75e7fb06a81", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:inbound", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:object_refs", "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:outbound"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "parent_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies", "path": "docs/guides/resources--bgp--properties--peers--routing_policies--route_policy.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/route_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.routing_policies.route_policy for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.routing_policies](resources--bgp--properties--peers--routing_policies.md)
- peers.routing_policies.route_policy

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

- [all_nodes](resources--bgp--properties--peers--routing_policies--route_policy--all_nodes.md): complete subsection reference.

- [inbound](resources--bgp--properties--peers--routing_policies--route_policy--inbound.md): complete subsection reference.

- [node_name](resources--bgp--properties--peers--routing_policies--route_policy--node_name.md): complete subsection reference.

- [object_refs](resources--bgp--properties--peers--routing_policies--route_policy--object_refs.md): complete subsection reference.

- [outbound](resources--bgp--properties--peers--routing_policies--route_policy--outbound.md): complete subsection reference.

## Next pages

- [peers.routing_policies.route_policy.all_nodes](resources--bgp--properties--peers--routing_policies--route_policy--all_nodes.md)
- [peers.routing_policies.route_policy.inbound](resources--bgp--properties--peers--routing_policies--route_policy--inbound.md)
- [peers.routing_policies.route_policy.node_name](resources--bgp--properties--peers--routing_policies--route_policy--node_name.md)
- [peers.routing_policies.route_policy.object_refs](resources--bgp--properties--peers--routing_policies--route_policy--object_refs.md)
- [peers.routing_policies.route_policy.outbound](resources--bgp--properties--peers--routing_policies--route_policy--outbound.md)
- [peers.routing_policies](resources--bgp--properties--peers--routing_policies.md)
- [xcsh_bgp](../resources/bgp.md)
