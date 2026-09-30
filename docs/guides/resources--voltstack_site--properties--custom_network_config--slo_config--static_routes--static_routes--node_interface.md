---
page_title: "custom_network_config.slo_config.static_routes.static_routes.node_interface"
subcategory: ""
description: "custom_network_config.slo_config.static_routes.static_routes.node_interface for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1890, "body_sha256": "sha256:47484e8188e941ed4227b9c0e5c9bb64d5fe7ea110f42c5f69b9f9aee18f35dd", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_routes", "static_routes", "node_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.slo_config.static_routes.static_routes.node_interface for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.slo_config.static_routes.static_routes.node_interface

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.slo_config](resources--voltstack_site--properties--custom_network_config--slo_config.md)
- [custom_network_config.slo_config.static_routes](resources--voltstack_site--properties--custom_network_config--slo_config--static_routes.md)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--properties--custom_network_config--slo_config--static_routes--static_routes.md)
- custom_network_config.slo_config.static_routes.static_routes.node_interface

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

- [list](resources--voltstack_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list.md): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--voltstack_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list.md)
- [custom_network_config.slo_config.static_routes.static_routes](resources--voltstack_site--properties--custom_network_config--slo_config--static_routes--static_routes.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
