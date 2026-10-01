---
page_title: "local_vrf.slo_config.static_routes.static_routes.node_interface"
subcategory: ""
description: "local_vrf.slo_config.static_routes.static_routes.node_interface for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1857, "body_sha256": "sha256:66dfc3d98598fec8cafaaaa1b87487d0060bf85928d540bd9ce30a3296f91654", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:node_interface", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:node_interface:list"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:node_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "path": "docs/guides/resources--securemesh_site_v2--properties--local_vrf--slo_config--static_routes--static_routes--node_interface.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_vrf", "slo_config", "static_routes", "static_routes", "node_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_vrf.slo_config.static_routes.static_routes.node_interface for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.slo_config.static_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [local_vrf](resources--securemesh_site_v2--properties--local_vrf.md)
- [local_vrf.slo_config](resources--securemesh_site_v2--properties--local_vrf--slo_config.md)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--properties--local_vrf--slo_config--static_routes.md)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--properties--local_vrf--slo_config--static_routes--static_routes.md)
- local_vrf.slo_config.static_routes.static_routes.node_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [list](resources--securemesh_site_v2--properties--local_vrf--slo_config--static_routes--static_routes--node_interface--list.md): complete subsection reference.

## Next pages

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--properties--local_vrf--slo_config--static_routes--static_routes--node_interface--list.md)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--properties--local_vrf--slo_config--static_routes--static_routes.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
