---
page_title: "slo_to_global_dr"
subcategory: "Networking"
description: "slo_to_global_dr for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1074, "body_sha256": "sha256:66a595e614322e3bb8e5d76a0d98d9cc806fcfe86a9c883ffa620a129436a12b", "canonical_id": "xcsh-docs:resources:network_connector:properties:slo_to_global_dr", "child_ids": ["xcsh-docs:resources:network_connector:properties:slo_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:properties:slo_to_global_dr", "parent_id": "xcsh-docs:resources:network_connector:reference", "path": "docs/guides/resources--network_connector--properties--slo_to_global_dr.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["slo_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/properties/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "slo_to_global_dr for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slo_to_global_dr

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md)
- [Property reference](resources--network_connector--reference.md)
- slo_to_global_dr

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](resources--network_connector--properties--slo_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [slo_to_global_dr.global_vn](resources--network_connector--properties--slo_to_global_dr--global_vn.md)
- [Property reference](resources--network_connector--reference.md)
- [xcsh_network_connector](../resources/network_connector.md)
