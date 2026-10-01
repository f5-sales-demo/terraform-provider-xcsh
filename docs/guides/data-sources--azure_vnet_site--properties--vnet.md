---
page_title: "vnet"
subcategory: "Infrastructure"
description: "vnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1230, "body_sha256": "sha256:cf6ecf22a3c5f17bf88b977d95b52d5808e421d8fce910dae142a6b18532ea87", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "docs/guides/data-sources--azure_vnet_site--properties--vnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- vnet

<a id="section"></a>

Type: `"single"`. Computed.

Defines choice about Azure VNet for a view.

Upstream description:

This defines choice about Azure VNet for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_vnet\",\"new_vnet\"]"
}
```

## Direct properties

- [existing_vnet](data-sources--azure_vnet_site--properties--vnet--existing_vnet.md): complete subsection reference.

- [new_vnet](data-sources--azure_vnet_site--properties--vnet--new_vnet.md): complete subsection reference.

## Next pages

- [vnet.existing_vnet](data-sources--azure_vnet_site--properties--vnet--existing_vnet.md)
- [vnet.new_vnet](data-sources--azure_vnet_site--properties--vnet--new_vnet.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
