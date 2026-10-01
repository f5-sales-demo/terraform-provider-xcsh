---
page_title: "vnet.new_vnet.autogenerate"
subcategory: "Infrastructure"
description: "vnet.new_vnet.autogenerate for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1087, "body_sha256": "sha256:827e0451c0956b424dc3424d4d7e21095699ae41576a737db1a988e336c0b4d0", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet:autogenerate", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet:autogenerate", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:new_vnet", "path": "docs/guides/resources--azure_vnet_site--properties--vnet--new_vnet--autogenerate.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vnet", "new_vnet", "autogenerate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/vnet/new_vnet/autogenerate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vnet.new_vnet.autogenerate for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet.new_vnet.autogenerate

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [vnet](resources--azure_vnet_site--properties--vnet.md)
- [vnet.new_vnet](resources--azure_vnet_site--properties--vnet--new_vnet.md)
- vnet.new_vnet.autogenerate

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for autogenerate.

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

Terraform syntax:

```terraform
autogenerate = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [vnet.new_vnet](resources--azure_vnet_site--properties--vnet--new_vnet.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
