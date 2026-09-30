---
page_title: "block_all_services"
subcategory: "Infrastructure"
description: "block_all_services for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1343, "body_sha256": "sha256:c35a90ac959dbc75e3d73929c314777c4bfba24c2c8026a6df5b23fa20ea5153", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:block_all_services", "child_ids": [], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:block_all_services", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "docs/guides/data-sources--azure_vnet_site--properties--block_all_services.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["block_all_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/block_all_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "block_all_services for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# block_all_services

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- block_all_services

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option. Defaults to \`map\[\]\`. Server applies default
when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [block_all_services](data-sources--azure_vnet_site--properties--block_all_services.md#section)
- [blocked_services](data-sources--azure_vnet_site--properties--blocked_services.md#section)
- [default_blocked_services](data-sources--azure_vnet_site--properties--default_blocked_services.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
