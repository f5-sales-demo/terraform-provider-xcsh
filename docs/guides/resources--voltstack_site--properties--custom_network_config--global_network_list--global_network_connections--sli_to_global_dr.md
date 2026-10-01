---
page_title: "custom_network_config.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: ""
description: "custom_network_config.global_network_list.global_network_connections.sli_to_global_dr for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1934, "body_sha256": "sha256:372473c4582acae50d1ff85a53794ff756cb946749d875138b79a3eaeb3d5828", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.global_network_list.global_network_connections.sli_to_global_dr for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.global_network_list](resources--voltstack_site--properties--custom_network_config--global_network_list.md)
- [custom_network_config.global_network_list.global_network_connections](resources--voltstack_site--properties--custom_network_config--global_network_list--global_network_connections.md)
- custom_network_config.global_network_list.global_network_connections.sli_to_global_dr

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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](resources--voltstack_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--voltstack_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md)
- [custom_network_config.global_network_list.global_network_connections](resources--voltstack_site--properties--custom_network_config--global_network_list--global_network_connections.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
