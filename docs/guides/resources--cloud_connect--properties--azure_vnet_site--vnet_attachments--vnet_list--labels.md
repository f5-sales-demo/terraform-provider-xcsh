---
page_title: "azure_vnet_site.vnet_attachments.vnet_list.labels"
subcategory: ""
description: "azure_vnet_site.vnet_attachments.vnet_list.labels for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1269, "body_sha256": "sha256:ab4604daab7534a9f30cfa2e2551bb578e58c8cfb0e32947e77e095c6e193d65", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:labels", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list:labels", "parent_id": "xcsh-docs:resources:cloud_connect:properties:azure_vnet_site:vnet_attachments:vnet_list", "path": "docs/guides/resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list--labels.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_vnet_site", "vnet_attachments", "vnet_list", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/azure_vnet_site/vnet_attachments/vnet_list/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_vnet_site.vnet_attachments.vnet_list.labels for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure_vnet_site.vnet_attachments.vnet_list.labels

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [azure_vnet_site](resources--cloud_connect--properties--azure_vnet_site.md)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments.md)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- azure_vnet_site.vnet_attachments.vnet_list.labels

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VNet attachments. These labels can then be used in policies such as enhanced
firewall policies.

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
labels {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--properties--azure_vnet_site--vnet_attachments--vnet_list.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
