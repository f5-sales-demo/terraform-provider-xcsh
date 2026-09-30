---
page_title: "vn_config.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: ""
description: "vn_config.global_network_list.global_network_connections.sli_to_global_dr for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1657, "body_sha256": "sha256:bacb27d4a2ab1cf960cce0aeae2dbe3ea0a455f1046a16be75f7173cbadd9cbf", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.global_network_list.global_network_connections.sli_to_global_dr for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vn_config.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vn_config.global_network_list](resources--aws_tgw_site--properties--vn_config--global_network_list.md)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr

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

- [global_vn](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
