---
page_title: "peers.routing_policies"
subcategory: ""
description: "peers.routing_policies for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1064, "body_sha256": "sha256:1760aaecd15881228974ed77cc04b13459f905c92087eebcf056c081efc2a33b", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "docs/guides/resources--bgp--properties--peers--routing_policies.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "routing_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.routing_policies for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- peers.routing_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
routing_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [route_policy](resources--bgp--properties--peers--routing_policies--route_policy.md): complete subsection reference.

## Next pages

- [peers.routing_policies.route_policy](resources--bgp--properties--peers--routing_policies--route_policy.md)
- [peers](resources--bgp--properties--peers.md)
- [xcsh_bgp](../resources/bgp.md)
