---
page_title: "offline_survivability_mode.enable_offline_survivability_mode"
subcategory: "Infrastructure"
description: "offline_survivability_mode.enable_offline_survivability_mode for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1189, "body_sha256": "sha256:1f5ddc26fe0b063cd7c3b2679309232389bf60333473b01aab1a9d362f6c4924", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:offline_survivability_mode", "path": "docs/guides/resources--azure_vnet_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["offline_survivability_mode", "enable_offline_survivability_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/offline_survivability_mode/enable_offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "offline_survivability_mode.enable_offline_survivability_mode for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode.enable_offline_survivability_mode

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [offline_survivability_mode](resources--azure_vnet_site--properties--offline_survivability_mode.md)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [offline_survivability_mode](resources--azure_vnet_site--properties--offline_survivability_mode.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
