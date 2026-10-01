---
page_title: "azure_vnet_site"
subcategory: ""
description: "azure_vnet_site for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1329, "body_sha256": "sha256:b0be89916ac9913c431f1a3167b2b446474d08dd892e4c39efac0928938c5eb0", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:site", "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site", "parent_id": "xcsh-docs:resources:cloud_connect:reference", "path": "docs/guides/resources--cloud_connect--properties--azure_vnet_site.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_vnet_site"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/azure_vnet_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_vnet_site for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_vnet_site

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- azure_vnet_site

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Azure VNet Site Type. Cloud Connect Azure VNet Site Type.

Upstream description:

Cloud Connect Azure VNet Site Type.

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
azure_vnet_site {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--cloud_connect--properties--azure_vnet_site--site.md): complete subsection reference.

- [vnet_attachments](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md): complete subsection reference.

## Next pages

- [azure_vnet_site.site](resources--cloud_connect--properties--azure_vnet_site--site.md)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- [Property reference](resources--cloud_connect--reference.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
