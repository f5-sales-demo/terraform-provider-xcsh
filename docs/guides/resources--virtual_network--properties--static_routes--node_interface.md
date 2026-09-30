---
page_title: "static_routes.node_interface"
subcategory: "Networking"
description: "static_routes.node_interface for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 1116, "body_sha256": "sha256:266026a6d4ea3b4af25897a13fa8cbecc37ab5daf055f1ceab82303d1b1bb31f", "canonical_id": "xcsh-docs:resources:virtual_network:properties:static_routes:node_interface", "child_ids": ["xcsh-docs:resources:virtual_network:properties:static_routes:node_interface:list"], "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:properties:static_routes:node_interface", "parent_id": "xcsh-docs:resources:virtual_network:properties:static_routes", "path": "docs/guides/resources--virtual_network--properties--static_routes--node_interface.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["static_routes", "node_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "static_routes.node_interface for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# static_routes.node_interface

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md)
- [Property reference](resources--virtual_network--reference.md)
- [static_routes](resources--virtual_network--properties--static_routes.md)
- static_routes.node_interface

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

- [list](resources--virtual_network--properties--static_routes--node_interface--list.md): complete subsection reference.

## Next pages

- [static_routes.node_interface.list](resources--virtual_network--properties--static_routes--node_interface--list.md)
- [static_routes](resources--virtual_network--properties--static_routes.md)
- [xcsh_virtual_network](../resources/virtual_network.md)
