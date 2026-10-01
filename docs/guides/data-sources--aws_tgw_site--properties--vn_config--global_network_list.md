---
page_title: "vn_config.global_network_list"
subcategory: ""
description: "vn_config.global_network_list for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1213, "body_sha256": "sha256:291b4824788eea1cfed7199070b35d05f0cfe0037ddd682568567f50301ae392", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:global_network_list", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "path": "docs/guides/data-sources--aws_tgw_site--properties--vn_config--global_network_list.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "global_network_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/global_network_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.global_network_list for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.global_network_list

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- vn_config.global_network_list

<a id="section"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

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

## Direct properties

- [global_network_connections](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md): complete subsection reference.

## Next pages

- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
