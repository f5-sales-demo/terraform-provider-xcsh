---
page_title: "vn_config.global_network_list.global_network_connections.slo_to_global_dr"
subcategory: ""
description: "vn_config.global_network_list.global_network_connections.slo_to_global_dr for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1756, "body_sha256": "sha256:26dbbe093ed53ffa1d9c4ef612645b7f0505a2f52242fea5fcc8da12815f170b", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "global_network_list", "global_network_connections", "slo_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.global_network_list.global_network_connections.slo_to_global_dr for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.global_network_list.global_network_connections.slo_to_global_dr

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vn_config.global_network_list](resources--aws_tgw_site--properties--vn_config--global_network_list.md)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr

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

- [global_vn](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
