---
page_title: "segment_vrf.segment_config.static_v6_routes.static_routes.node_interface"
subcategory: ""
description: "segment_vrf.segment_config.static_v6_routes.static_routes.node_interface for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1972, "body_sha256": "sha256:229eb7078cbe73cf67c59ca9239767dfbdde38540f69831a5cb237244f607dce", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes:node_interface", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes:node_interface:list"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes:node_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes", "path": "docs/guides/resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes--static_routes--node_interface.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_v6_routes", "static_routes", "node_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_vrf.segment_config.static_v6_routes.static_routes.node_interface for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [segment_vrf](resources--securemesh_site_v2--properties--segment_vrf.md)
- [segment_vrf.segment_config](resources--securemesh_site_v2--properties--segment_vrf--segment_config.md)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes.md)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes--static_routes.md)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

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

- [list](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes--static_routes--node_interface--list.md): complete subsection reference.

## Next pages

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes--static_routes--node_interface--list.md)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--properties--segment_vrf--segment_config--static_v6_routes--static_routes.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
